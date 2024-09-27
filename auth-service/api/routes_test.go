package api

import (
	"auth-service/api/handlers"
	"auth-service/client"
	"auth-service/postgres"
	"auth-service/proto"
	"context"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var authClient proto.AuthServiceClient

func TestAuthService(t *testing.T) {
	// Load environment variables
	os.Setenv("GRPC_AUTH_PORT", "8000")

	// --- Server Setup ---
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})

	// --- Database Setup ---
	ctx := context.Background()
	db := postgres.NewTestDBConn(ctx, t)
	// -----------------------

	asc := AuthService{
		UserHandlers: handlers.NewUserHandlers(db),
		JwtHandler:   handlers.NewJwtHandler(db),
	}
	proto.RegisterAuthServiceServer(srv, &asc)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	// --- Client Setup ---
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn := client.NewTestingGrpcConnection(dialer)
	t.Cleanup(func() {
		conn.Close()
	})

	authClient = proto.NewAuthServiceClient(conn)

	// --- Tests ---
	handlers.CreateUserTest(authClient, t)
	userLoggedIn := handlers.LoginUserTest(authClient, t)
	handlers.JwtParseTest(authClient, userLoggedIn, t)
}
