package user

import (
	"authApp/cmd/user/dto"
	"authApp/models"
	"authApp/util"
	"errors"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandlers struct {
	db *gorm.DB
}

func NewUserHandlers(db *gorm.DB) *UserHandlers {
	return &UserHandlers{db}
}

func (u *UserHandlers) CreateUser(user *dto.RequestUserDto) (*models.User, error) {
	newUser := &models.User{Email: user.Email}
	u.db.First(newUser, "email = ?", user.Email)

	if newUser.Id != uuid.Nil {
		return &models.User{}, errors.New("user already exists")
	}

	var err error
	newUser.Password, err = util.GenerateHash(user.Password)
	if err != nil {
		log.Println(err)
	}

	u.db.Create(newUser)

	return newUser, nil

}

func (u *UserHandlers) ValidateUser(user *dto.RequestUserDto) (*models.User, error) {
	newUser := &models.User{Email: user.Email}
	u.db.First(newUser, "email = ?", user.Email)

	if newUser.Id == uuid.Nil {
		return &models.User{}, errors.New("user not found")
	}

	err := util.CompareHash(newUser.Password, user.Password)
	if err != nil {
		log.Println(err)
		return &models.User{}, errors.New("invalid password")
	}

	return newUser, nil
}
