package api

import (
	"auth-service/api/dto"
	"auth-service/api/handlers"
	"auth-service/proto"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
	UserHandlers *handlers.UserHandlers
	JwtHandler   *handlers.JwtHandler
}

func (u *AuthService) CreateUser(ctx context.Context, req *proto.UserRequest) (*proto.User, error) {
	payload, err := dto.NewCreateUserDto(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		return &proto.User{}, status.Error(codes.InvalidArgument, err.Error())
	}

	newUser, err := u.UserHandlers.CreateUser(payload)
	if err != nil {
		return &proto.User{}, err
	}

	return &proto.User{
		Id:        newUser.Id.String(),
		Name:      newUser.Name,
		Email:     newUser.Email,
		IsActive:  newUser.IsActive,
		Role:      proto.Role(proto.Role_value[newUser.Role]),
		CreatedAt: newUser.CreatedAt.String(),
		UpdatedAt: newUser.UpdatedAt.String(),
	}, nil
}

func (u *AuthService) ValidateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserValidated, error) {
	payload, err := dto.NewRequestUserDto(req.Email, req.Password)
	if err != nil {
		return &proto.UserValidated{}, status.Error(codes.InvalidArgument, err.Error())
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

func (u *AuthService) ActivateUser(ctx context.Context, req *proto.VerificationCodeRequest) (*proto.UserResponse, error) {
	payload, err := dto.NewRequestVerificationCodeDto(req.Email, req.Code)
	if err != nil {
		return &proto.UserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := u.UserHandlers.ActivateUser(payload.Code, payload.Email)
	if err != nil {
		return &proto.UserResponse{}, err
	}

	token, err := u.JwtHandler.GenerateToken(user)
	if err != nil {
		return &proto.UserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.UserResponse{
		Token: token,
		User: &proto.User{
			Id:        user.Id.String(),
			Email:     user.Email,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}, nil
}

func (u *AuthService) ResendVerificationCode(ctx context.Context, req *proto.ResendVerificationCodeRequest) (*proto.ResendVerificationCodeResponse, error) {
	err := u.UserHandlers.ResendVerificationCode(req.Email)
	if err != nil {
		return &proto.ResendVerificationCodeResponse{}, err
	}

	return &proto.ResendVerificationCodeResponse{
		Message: "Verification code sent!",
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
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (u *AuthService) ResetPassword(ctx context.Context, req *proto.ResetPasswordRequest) (*proto.ResetPasswordResponse, error) {
	_, err := u.JwtHandler.ParseToken(req.Token, "reset_password")
	if err != nil {
		return &proto.ResetPasswordResponse{}, err
	}

	err = u.UserHandlers.ResetPassword(req.Email, req.Password)
	if err != nil {
		return &proto.ResetPasswordResponse{}, err
	}

	return &proto.ResetPasswordResponse{
		Message: "Password reset successfully!",
	}, nil
}

func (u *AuthService) SendResetPassword(ctx context.Context, req *proto.SendResetPasswordRequest) (*proto.SendResetPasswordResponse, error) {
	err := u.UserHandlers.SendResetPasswordEmail(req.FrontBaseUrl, req.Email)
	if err != nil {
		return &proto.SendResetPasswordResponse{}, err
	}

	return &proto.SendResetPasswordResponse{
		Message: "Reset password email sent!",
	}, nil
}
