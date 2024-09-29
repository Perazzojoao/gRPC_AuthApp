package dto

import (
	"github.com/go-playground/validator/v10"
)

type RequestUserDto struct {
	Email    string `validate:"email"`
	Password string `validate:"min=6"`
}

var validateUser = validator.New(validator.WithRequiredStructEnabled())

func NewRequestUserDto(email, password string) (*RequestUserDto, error) {
	user := RequestUserDto{
		Email:    email,
		Password: password,
	}
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *RequestUserDto) Validate() error {
	return validateUser.Struct(u)
}
