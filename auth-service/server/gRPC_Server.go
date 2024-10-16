package server

import (
	"auth-service/api"
	"auth-service/api/handlers"
	"auth-service/postgres"
	"auth-service/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func NewServer() *Server {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &Server{db}
}

func (app *Server) GrpcListen() {
	gRPCPort := "8000"

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRpc: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &api.AuthService{
		UserHandlers: handlers.NewUserHandlers(app.DB, handlers.NewMailHandler()),
		JwtHandler:   handlers.NewJwtHandler(app.DB),
	})
	log.Println("gRPC server started on port ", gRPCPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRpc: %v", err)
	}
}
