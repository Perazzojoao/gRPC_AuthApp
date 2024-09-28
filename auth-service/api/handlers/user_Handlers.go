package handlers

import (
	"auth-service/api/dto"
	"auth-service/postgres/models"
	"auth-service/util"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserHandlers struct {
	db          *gorm.DB
	MailHandler *MailHandler
}

func NewUserHandlers(db *gorm.DB, mailHandler *MailHandler) *UserHandlers {
	return &UserHandlers{db, mailHandler}
}

func (u *UserHandlers) CreateUser(user *dto.RequestUserDto) (*models.User, error) {
	newUser := models.User{Email: user.Email}
	result := u.db.First(&newUser, "email = ?", user.Email)

	if result.Error == nil {
		return &models.User{}, status.Error(codes.AlreadyExists, "user already exists")
	}

	var err error
	newUser.Password, err = util.GenerateHash(user.Password)
	if err != nil {
		log.Println(err)
	}

	result = u.db.Create(&newUser)
	if result.Error != nil {
		log.Println("error creating user: ", result.Error)
		return &models.User{}, status.Error(codes.Internal, "internal error")
	}

	verifyCode := models.VerificationCode{UserID: newUser.Id}
	result = u.db.Create(&verifyCode)
	if result.Error != nil {
		log.Println("error creating verification code: ", result.Error)
		return &models.User{}, status.Error(codes.Internal, "internal error")
	}

	msg := MailMessage{
		To:      newUser.Email,
		Subject: "Verify your email",
		Body:    "Your verification code is " + verifyCode.Code,
	}
	ctx := context.Background()
	go u.MailHandler.SendPlainTextMail(ctx, msg)

	return &newUser, nil
}

func (u *UserHandlers) ValidateUser(user *dto.RequestUserDto) (*models.User, error) {
	newUser := models.User{Email: user.Email}
	result := u.db.First(&newUser, "email = ?", user.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &models.User{}, status.Error(codes.NotFound, "user not found")
	}
	if !newUser.Active {
		return &models.User{}, status.Error(codes.PermissionDenied, "user not active")
	}

	err := util.CompareHash(newUser.Password, user.Password)
	if err != nil {
		return &models.User{}, status.Error(codes.Unauthenticated, "invalid password")
	}

	return &newUser, nil
}

func (u *UserHandlers) ActivateUser(verifyCode string, userId uuid.UUID) error {
	code := models.VerificationCode{}
	result := u.db.First(&code, "code = ? AND user_id = ?", verifyCode, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("error finding verification code: ", result.Error)
		return status.Error(codes.NotFound, "verification code not found")
	}

	result = u.db.Model(&models.User{}).Where("id = ?", userId).Update("is_active", true)
	if result.Error != nil {
		log.Println("error updating user: ", result.Error)
		return status.Error(codes.Internal, "internal error")
	}

	return nil
}
