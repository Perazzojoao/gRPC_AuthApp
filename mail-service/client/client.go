package client

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection() *grpc.ClientConn {
	serverUrl := fmt.Sprintf("localhost:%s", os.Getenv("GRPC_MAIL_PORT"))

	conn, err := grpc.NewClient(serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	return conn
}

func NewTestingGrpcConnection(dialer func(context.Context, string) (net.Conn, error)) *grpc.ClientConn {
	serverUrl := fmt.Sprintf("localhost:%s", os.Getenv("GRPC_MAIL_PORT"))

	conn, err := grpc.NewClient(serverUrl, grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	return conn
}
