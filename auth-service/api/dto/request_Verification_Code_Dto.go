package dto

import "github.com/go-playground/validator/v10"

type RequestVerificationCodeDto struct {
	Code  string `validate:"required"`
	Email string `validate:"required,email"`
}

var validateCode = validator.New(validator.WithRequiredStructEnabled())

func NewRequestVerificationCodeDto(email, code string) (*RequestVerificationCodeDto, error) {
	request := RequestVerificationCodeDto{
		Email: email,
		Code:  code,
	}
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func (u *RequestVerificationCodeDto) Validate() error {
	return validateCode.Struct(u)
}
