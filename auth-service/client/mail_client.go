package client

import (
	proto "auth-service/proto/mail_proto"

	"google.golang.org/grpc"
)

func NewMailGrpcClient(conn *grpc.ClientConn) proto.MailServiceClient {
	return proto.NewMailServiceClient(conn)
}
