package handlers

import (
	"auth-service/api/dto"
	"auth-service/postgres/models"
	"auth-service/util"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (u *UserHandlers) ActivateUser(verifyCode string, userId string) (*models.User, error) {
	code := models.VerificationCode{}
	result := u.db.First(&code, "code = ? AND user_id = ?", verifyCode, userId)
	if result.Error != nil {
		log.Println("error finding verification code: ", result.Error)
		return nil, status.Error(codes.NotFound, "verification code not found")
	}

	user := models.User{}
	result = u.db.Model(&user).Where("id = ?", userId).Update("active", true).First(&user)
	if result.Error != nil {
		log.Println("error updating user: ", result.Error)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &user, nil
}

func (u *UserHandlers) ResendVerificationCode(email string) error {
	user := models.User{}
	result := u.db.First(&user, "email = ?", email)
	if result.Error != nil {
		log.Println("error finding user: ", result.Error)
		return status.Error(codes.NotFound, "user not found")
	}

	code := models.VerificationCode{Code: strconv.Itoa(int(models.GenerateVerifyCode()))}
	result = u.db.Model(&code).Where("user_id = ?", user.Id).Update("code", code.Code)
	if result.Error != nil {
		log.Println("error finding verification code: ", result.Error)
		return status.Error(codes.NotFound, "verification code not found")
	}

	msg := MailMessage{
		To:      user.Email,
		Subject: "Verify your email",
		Body:    "Your verification code is " + code.Code,
	}
	ctx := context.Background()
	go u.MailHandler.SendPlainTextMail(ctx, msg)

	return nil
}

func (u *UserHandlers) ResetPassword(email, password string) error {
	hashedPassword, err := util.GenerateHash(password)
	if err != nil {
		log.Println("error hashing password: ", err)
		return status.Error(codes.Internal, "internal error")
	}

	result := u.db.Model(&models.User{}).Where("email = ?", email).Update("password", hashedPassword)
	if result.Error != nil {
		log.Println("error updating password: ", result.Error)
		return status.Error(codes.Internal, "internal error")
	}

	return nil
}

func (u *UserHandlers) SendResetPasswordEmail(frontUrl, email string) error {
	user := models.User{}
	result := u.db.First(&user, "email = ?", email)
	if result.Error != nil {
		log.Println("error finding user: ", result.Error)
		return status.Error(codes.NotFound, "user not found")
	}

	token, err := util.GenerateToken(&user, "reset_password")
	if err != nil {
		log.Println("error generating token: ", err)
		return status.Error(codes.Internal, "internal error")
	}

	tokenChanel := make(chan *jwt.Token)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go func() {
		defer close(tokenChanel)
		parsedT, err := util.ParseToken(token, "reset_password")
		if err != nil {
			log.Println("error parsing token: ", err)
		}
		tokenChanel <- parsedT
	}()

	var parsedToken *jwt.Token

	select {
	case <-ctx.Done():
		log.Println("error parsing token: ", ctx.Err())
	case parsedToken = <-tokenChanel:
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	exp := claims["exp"].(float64)

	msg := MailMessage{
		To:      user.Email,
		Subject: "Reset your password",
		Body:    "Click this link to reset your password!\n" + frontUrl + fmt.Sprintf("?exp=%ftoken=%s", exp, token),
	}
	context := context.Background()
	go u.MailHandler.SendPlainTextMail(context, msg)

	return nil
}
