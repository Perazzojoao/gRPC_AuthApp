package dto

import (
	"github.com/go-playground/validator/v10"
)

const (
	ADMIN  string = "ADMIN"
	CLIENT string = "CLIENT"
)

type CreateUserDto struct {
	Name     string `validate:"min=3,max=100"`
	Role     string `validate:"role"`
	Email    string `validate:"email"`
	Password string `validate:"min=6"`
}

func (u *CreateUserDto) Validate() error {
	validateUser.RegisterValidation("role", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "" || fl.Field().String() == ADMIN || fl.Field().String() == CLIENT
	})
	return validateUser.Struct(u)
}

type RequestUserDto struct {
	Email    string `validate:"email"`
	Password string `validate:"min=6"`
}

func (u *RequestUserDto) Validate() error {
	return validateUser.Struct(u)
}

var validateUser = validator.New(validator.WithRequiredStructEnabled())

func NewCreateUserDto(name, email, password, role string) (*CreateUserDto, error) {
	user := CreateUserDto{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

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
