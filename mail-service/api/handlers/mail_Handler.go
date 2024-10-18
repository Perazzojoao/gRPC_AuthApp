package handlers

import (
	"mail-service/proto"
	"mail-service/util"
)

type MailHandler struct{}

func NewMailHandler() *MailHandler {
	return &MailHandler{}
}

func (m *MailHandler) SendVerificationCodeMail(msg *proto.MailRequest) error {
	newMsg := util.Message{
		To:      msg.To,
		Subject: msg.Subject,
		Data:    msg.Body,
	}

	return newMsg.SendGomailWithTemplate(util.VerificationCode)
}

func (m *MailHandler) SendResetPasswordMail(msg *proto.MailRequest) error {
	newMsg := util.Message{
		To:      msg.To,
		Subject: msg.Subject,
		Data:    msg.Body,
	}

	return newMsg.SendGomailWithTemplate(util.ResetPassword)
}

func (m *MailHandler) SendPlainTextMail(msg *proto.MailRequest) error {
	newMsg := util.Message{
		To:      msg.To,
		Subject: msg.Subject,
		Data:    msg.Body,
	}

	return newMsg.SendGomailPlainText()
}
