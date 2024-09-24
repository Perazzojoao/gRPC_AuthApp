package client

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection() *grpc.ClientConn {
	serverUrl := fmt.Sprintf("localhost:%s", os.Getenv("GRPC_PORT"))

	conn, err := grpc.NewClient(serverUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	return conn
}
