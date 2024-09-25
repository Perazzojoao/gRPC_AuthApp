package main_test

import (
	"authApp/client"
	"authApp/proto"
	"context"
	"testing"
	"time"
)

var payload = &proto.UserRequest{
	Email:    "test@gmail.com",
	Password: "password",
}

func TestMain(t *testing.T) {
	t.Setenv("GRPC_PORT", "8000")

	// _, err := db.Connect()
	// if err != nil {
	// 	t.Fatalf("failed to connect to database: %v", err)
	// }

	conn := client.NewGrpcConnection()
	defer conn.Close()

	client := proto.NewAuthServiceClient(conn)

	t.Run("CreateUser", func(t *testing.T) {
		ctx, close := context.WithTimeout(context.Background(), 2*time.Second)
		defer close()

		response, err := client.CreateUser(ctx, payload)
		if err != nil {
			t.Errorf("CreateUser failed: %v", err)
		}

		if response.Token == "" {
			t.Errorf("CreateUser failed: token is empty")
		}

		if response.User.Email != payload.Email {
			t.Errorf("CreateUser failed: email is not the same")
		}
	})
}
