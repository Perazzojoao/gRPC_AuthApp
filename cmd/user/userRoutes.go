package user

import (
	"authApp/cmd/jwt"
	"authApp/cmd/user/dto"
	"authApp/proto"
	"context"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
	UserHandlers *UserHandlers
	JwtHandler   *jwt.JwtHandler
}

func (u *AuthService) CreateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	payload, err := dto.NewRequestUserDto(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	newUser, err := u.UserHandlers.CreateUser(payload)
	if err != nil {
		return nil, err
	}

	token := u.JwtHandler.GenerateToken(newUser)

	return &proto.UserResponse{
		Token: token,
		User: &proto.User{
			Id:        newUser.Id.String(),
			Email:     newUser.Email,
			Active:    newUser.Active,
			CreatedAt: newUser.CreatedAt.String(),
			UpdatedAt: newUser.UpdatedAt.String(),
		},
	}, nil
}

func (u *AuthService) ValidateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserValidated, error) {
	payload, err := dto.NewRequestUserDto(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	newUser, err := u.UserHandlers.ValidateUser(payload)
	if err != nil {
		return nil, err
	}

	token := u.JwtHandler.GenerateToken(newUser)

	return &proto.UserValidated{
		Token: token,
		Id:    newUser.Id.String(),
		Email: newUser.Email,
	}, nil
}

func (u *AuthService) JwtParse(ctx context.Context, req *proto.Jwt) (*proto.User, error) {
	user, err := u.JwtHandler.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &proto.User{
		Id:        user.Id.String(),
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
