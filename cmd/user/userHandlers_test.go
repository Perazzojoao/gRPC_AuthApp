package user

import (
	"authApp/proto"
	"context"
	"sync"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	tt := []struct {
		name    string
		payload *proto.UserRequest
		wantErr bool
	}{
		{
			name:    "Test 1",
			payload: &proto.UserRequest{Email: "testgmail.com", Password: "password"},
			wantErr: true,
		},
		{
			name:    "Test 2",
			payload: &proto.UserRequest{Email: "test@gmail.com", Password: "passw"},
			wantErr: true,
		},
		{
			name:    "Test 3",
			payload: &proto.UserRequest{Email: "@gmail.com", Password: "password"},
			wantErr: true,
		},
		{
			name:    "Test 4",
			payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password"},
			wantErr: false,
		},
		{
			name:    "Test 5",
			payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password"},
			wantErr: true,
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
}

func TestLoginUser(t *testing.T) {
	tt := []struct {
		name    string
		payload *proto.UserRequest
		wantErr bool
	}{
		{
			name:    "Test 1",
			payload: &proto.UserRequest{Email: "test1@gmail.com", Password: "password"},
			wantErr: true,
		},
		{
			name:    "Test 2",
			payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password1"},
			wantErr: true,
		},
		{
			name:    "Test 3",
			payload: &proto.UserRequest{Email: "test@gmail", Password: "password"},
			wantErr: true,
		},
		{
			name:    "Test 4",
			payload: &proto.UserRequest{Email: "test@gmail.com", Password: "password"},
			wantErr: false,
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
				if tc.wantErr && err == nil {
					t.Errorf("expected error, got nil")
				}
				if !tc.wantErr && err != nil {
					t.Errorf("expected no error, got %v", err)
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
}

func TestJwtParse(t *testing.T) {
	token := userLoggedIn.Token

	tt := []struct {
		name    string
		payload *proto.Jwt
		wantErr bool
	}{
		{
			name:    "Test 1",
			payload: &proto.Jwt{},
			wantErr: true,
		},
		{
			name:    "Test 2",
			payload: &proto.Jwt{Token: token + "1"},
			wantErr: true,
		},
		{
			name:    "Test 3",
			payload: &proto.Jwt{Token: token},
			wantErr: false,
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
				if tc.wantErr && err == nil {
					t.Errorf("expected error, got nil")
				}
				if !tc.wantErr && err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			})
		}()
	}

	wg.Wait()
}
