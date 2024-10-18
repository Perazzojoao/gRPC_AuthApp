package handlers

import (
	"auth-service/proto"
	"context"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateUserTest(authClient proto.AuthServiceClient, t *testing.T) {
	t.Run("Create User", func(t *testing.T) {
		tt := []struct {
			name    string
			payload *proto.UserRequest
			wantErr bool
			expect  codes.Code
		}{
			{
				name:    "Test 1",
				payload: &proto.UserRequest{Email: "testgmail.com", Password: "password"},
				wantErr: true,
				expect:  codes.InvalidArgument,
			},
			{
				name:    "Test 2",
				payload: &proto.UserRequest{Email: "test@gmail.com", Password: "passw"},
				wantErr: true,
				expect:  codes.InvalidArgument,
			},
			{
				name:    "Test 3",
				payload: &proto.UserRequest{Email: "@gmail.com", Password: "password"},
				wantErr: true,
				expect:  codes.InvalidArgument,
			},
			{
				name:    "Test 4",
				payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password"},
				wantErr: false,
				expect:  codes.OK,
			},
			{
				name:    "Test 5",
				payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password"},
				wantErr: true,
				expect:  codes.AlreadyExists,
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		t.Cleanup(cancel)

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				_, err := authClient.CreateUser(ctx, tc.payload)
				if tc.wantErr && err == nil {
					t.Errorf("expected error, got nil")
				}
				if !tc.wantErr && err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			})
		}
	})
}

func LoginUserTest(authClient proto.AuthServiceClient, t *testing.T) *proto.UserValidated {
	var userLoggedIn *proto.UserValidated
	t.Run("Login User", func(t *testing.T) {
		tt := []struct {
			name    string
			payload *proto.UserRequest
			wantErr bool
			expect  codes.Code
		}{
			{
				name:    "Test 1",
				payload: &proto.UserRequest{Email: "test1@gmail.com", Password: "password"},
				wantErr: true,
				expect:  codes.NotFound,
			},
			{
				name:    "Test 2",
				payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password1"},
				wantErr: true,
				expect:  codes.Unauthenticated,
			},
			{
				name:    "Test 3",
				payload: &proto.UserRequest{Email: "test@gmail", Password: "password"},
				wantErr: true,
				expect:  codes.InvalidArgument,
			},
			{
				name:    "Test 4",
				payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password"},
				wantErr: false,
				expect:  codes.PermissionDenied,
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		t.Cleanup(cancel)

		var user = make(chan *proto.UserValidated, 1)
		var wg sync.WaitGroup

		for _, tc := range tt {
			wg.Add(1)
			go func() {
				defer wg.Done()
				t.Run(tc.name, func(t *testing.T) {
					response, err := authClient.ValidateUser(ctx, tc.payload)
					errStatus, _ := status.FromError(err)
					statusCode := errStatus.Code()

					if tc.expect != statusCode {
						t.Errorf("expected status code %v, got %v\n error message: %v", tc.expect, statusCode, err)
					}

					if response != nil {
						select {
						case user <- response:
						default:
						}
					}
				})
			}()
		}

		wg.Wait()
		userLoggedIn = <-user
		close(user)
	})
	return userLoggedIn
}

func JwtParseTest(authClient proto.AuthServiceClient, userLoggedIn *proto.UserValidated, t *testing.T) {
	t.Run("Parse JWT", func(t *testing.T) {
		token := userLoggedIn.Token

		tt := []struct {
			name    string
			payload *proto.Jwt
			wantErr bool
			expect  codes.Code
		}{
			{
				name:    "Test 1",
				payload: &proto.Jwt{},
				wantErr: true,
				expect:  codes.Unauthenticated,
			},
			{
				name:    "Test 2",
				payload: &proto.Jwt{Token: token + "1"},
				wantErr: true,
				expect:  codes.Unauthenticated,
			},
			{
				name:    "Test 3",
				payload: &proto.Jwt{Token: token},
				wantErr: false,
				expect:  codes.OK,
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		t.Cleanup(cancel)

		var wg sync.WaitGroup
		for _, tc := range tt {
			wg.Add(1)
			go func() {
				defer wg.Done()
				t.Run(tc.name, func(t *testing.T) {
					_, err := authClient.JwtParse(ctx, tc.payload)
					errStatus, _ := status.FromError(err)
					statusCode := errStatus.Code()

					if tc.expect != statusCode {
						t.Errorf("expected status code %v, got %v\n error message: %v", tc.expect, statusCode, err)
					}
				})
			}()
		}

		wg.Wait()
	})
}
