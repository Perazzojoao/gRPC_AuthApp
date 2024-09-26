package api

import (
	"context"
	"mail-service/proto"
)

type MailService struct {
	proto.UnimplementedMailServiceServer
}

func (s *MailService) SendMail(ctx context.Context, req *proto.MailRequest) (*proto.MailResponse, error) {
	return &proto.MailResponse{}, nil
}
