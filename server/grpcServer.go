package server

import (
	"authApp/cmd/jwt"
	"authApp/cmd/user"
	"authApp/db"
	"authApp/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

var GRpcPort = "8000"

func NewServer() *Server {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &Server{db}
}

func (app *Server) GrpcListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", GRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRpc: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &user.AuthService{
		UserHandlers: user.NewUserHandlers(app.DB),
		JwtHandler:   jwt.NewJwtHandler(app.DB),
	})
	log.Println("gRPC server started on port ", GRpcPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRpc: %v", err)
	}
}
