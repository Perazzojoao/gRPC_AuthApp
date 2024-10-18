package dto

import "github.com/go-playground/validator/v10"

type RequestVerificationCodeDto struct {
	Code string `validate:"required"`
	Id   string `validate:"required,uuid4"`
}

var validateCode = validator.New(validator.WithRequiredStructEnabled())

func NewRequestVerificationCodeDto(userId, code string) (*RequestVerificationCodeDto, error) {
	request := RequestVerificationCodeDto{
		Id:   userId,
		Code: code,
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
