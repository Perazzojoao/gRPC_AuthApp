package user

import (
	"authApp/cmd/jwt"
	"authApp/cmd/user/dto"
	"authApp/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
	UserHandlers *UserHandlers
	JwtHandler   *jwt.JwtHandler
}

func (u *AuthService) CreateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	payload, err := dto.NewRequestUserDto(req.Email, req.Password)
	if err != nil {
		return &proto.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	newUser, err := u.UserHandlers.CreateUser(payload)
	if err != nil {
		return &proto.UserResponse{}, err
	}

	token, err := u.JwtHandler.GenerateToken(newUser)
	if err != nil {
		return &proto.UserResponse{}, status.Error(codes.Internal, err.Error())
	}

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
		return &proto.UserValidated{}, err
	}

	newUser, err := u.UserHandlers.ValidateUser(payload)
	if err != nil {
		return &proto.UserValidated{}, err
	}

	token, err := u.JwtHandler.GenerateToken(newUser)
	if err != nil {
		return &proto.UserValidated{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.UserValidated{
		Token: token,
		Id:    newUser.Id.String(),
		Email: newUser.Email,
	}, nil
}

func (u *AuthService) JwtParse(ctx context.Context, req *proto.Jwt) (*proto.User, error) {
	user, err := u.JwtHandler.ParseToken(req.Token)
	if err != nil {
		return &proto.User{}, err
	}

	return &proto.User{
		Id:        user.Id.String(),
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
