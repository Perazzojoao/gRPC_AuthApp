package handlers

import (
	"auth-service/api/dto"
	"auth-service/postgres/models"
	"auth-service/util"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserHandlers struct {
	db *gorm.DB
}

func NewUserHandlers(db *gorm.DB) *UserHandlers {
	return &UserHandlers{db}
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
		return &models.User{}, status.Error(codes.Internal, "internal error")
	}

	return &newUser, nil

}

func (u *UserHandlers) ValidateUser(user *dto.RequestUserDto) (*models.User, error) {
	newUser := models.User{Email: user.Email}
	result := u.db.First(&newUser, "email = ?", user.Email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &models.User{}, status.Error(codes.NotFound, "user not found")
	}

	err := util.CompareHash(newUser.Password, user.Password)
	if err != nil {
		return &models.User{}, status.Error(codes.Unauthenticated, "invalid password")
	}

	return &newUser, nil
}
