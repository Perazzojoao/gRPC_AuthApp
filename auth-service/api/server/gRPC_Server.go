package server

import (
	"auth-service/api/jwt"
	"auth-service/api/user"
	"auth-service/postgres"
	"auth-service/proto"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

var gRPCPort string

func NewServer() *Server {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &Server{db}
}

func (app *Server) GrpcListen() {
	gRPCPort = os.Getenv("GRPC_PORT")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRpc: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &user.AuthService{
		UserHandlers: user.NewUserHandlers(app.DB),
		JwtHandler:   jwt.NewJwtHandler(app.DB),
	})
	log.Println("gRPC server started on port ", gRPCPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRpc: %v", err)
	}
}