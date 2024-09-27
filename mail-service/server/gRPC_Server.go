package server

import (
	"fmt"
	"log"
	"mail-service/api"
	"mail-service/api/handlers"
	"mail-service/proto"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) GrpcListen() {
	err := godotenv.Load()
	gRPCPort := os.Getenv("GRPC_MAIL_PORT")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRpc: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	proto.RegisterMailServiceServer(s, &api.MailService{
		MailHandler: handlers.NewMailHandler(),
	})
	log.Println("gRPC server started on port ", gRPCPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRpc: %v", err)
	}
}
