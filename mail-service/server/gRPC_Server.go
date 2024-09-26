package server

import (
	"fmt"
	"log"
	"mail-service/api/mail"
	"mail-service/proto"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) GrpcListen() {
	gRPCPort := "8000"

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))
	if err != nil {
		log.Fatalf("failed to listen for gRpc: %v", err)
	}
	defer listen.Close()

	s := grpc.NewServer()
	proto.RegisterMailServiceServer(s, &mail.MailService{})
	log.Println("gRPC server started on port ", gRPCPort)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRpc: %v", err)
	}
}
