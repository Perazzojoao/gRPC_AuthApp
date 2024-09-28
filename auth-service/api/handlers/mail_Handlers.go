package handlers

import (
	"auth-service/client"
	proto "auth-service/proto/mail_proto"
	"context"
	"log"
	"time"
)

type MailMessage struct {
	To      string
	Subject string
	Body    string
}

type MailHandler struct {
	serviceUrl string
}

func NewMailHandler() *MailHandler {
	return &MailHandler{
		serviceUrl: "mail-service:8000",
	}
}

func (m *MailHandler) SendPlainTextMail(ctx context.Context, msg MailMessage) {
	conn := client.NewGrpcConnection(m.serviceUrl)
	defer conn.Close()
	mailClient := proto.NewMailServiceClient(conn)

	ctx, cancel := context.WithTimeout(ctx, 12*time.Second)
	defer cancel()
	resp := make(chan struct {
		*proto.MailResponse
		error
	})

	go func() {
		defer close(resp)
		response, err := mailClient.SendPlainTextMail(context.Background(), &proto.MailRequest{
			To:      msg.To,
			Subject: msg.Subject,
			Body:    msg.Body,
		})

		resp <- struct {
			*proto.MailResponse
			error
		}{response, err}
	}()

	select {
	case <-ctx.Done():
		log.Println("mail-service taking too long to respond")
	case r := <-resp:
		if r.error != nil {
			log.Println(r.error)
		} else {
			log.Println(r.MailResponse)
		}
	}
}
