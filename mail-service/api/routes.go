package api

import (
	"context"
	"mail-service/api/handlers"
	"mail-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MailService struct {
	proto.UnimplementedMailServiceServer
	MailHandler *handlers.MailHandler
}

const (
	successMsg = "Mail sent successfully"
	errorMsg   = "Failed to send mail"
)

func (s *MailService) SendVerificationCodeMail(ctx context.Context, req *proto.MailRequest) (*proto.MailResponse, error) {
	err := s.MailHandler.SendVerificationCodeMail(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errorMsg)
	}
	return &proto.MailResponse{
		Message: successMsg,
	}, nil
}

func (s *MailService) SendResetPasswordMail(ctx context.Context, req *proto.MailRequest) (*proto.MailResponse, error) {
	err := s.MailHandler.SendResetPasswordMail(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errorMsg)
	}
	return &proto.MailResponse{
		Message: successMsg,
	}, nil
}
