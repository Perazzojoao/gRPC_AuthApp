package jwt

import (
	"authApp/models"
	"authApp/util"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type JwtHandler struct {
	db *gorm.DB
}

func NewJwtHandler(db *gorm.DB) *JwtHandler {
	return &JwtHandler{db: db}
}

func (j *JwtHandler) GenerateToken(user *models.User) string {
	token, err := util.GenerateToken(user)
	if err != nil {
		log.Println("Could not generate jwt token: ", err)
		return ""
	}
	return token
}

func (j *JwtHandler) ParseToken(tokenString string) (*models.User, error) {
	token, err := util.ParseToken(tokenString)
	if err != nil {
		log.Println("Could not parse jwt token: ", err)
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("could not parse jwt token: %v", err))
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, status.Error(codes.Unauthenticated, "token expired")
		}
	}

	var user models.User
	j.db.First(&user, "id = ?", token.Claims.(jwt.MapClaims)["sub"])
	if user.Id == uuid.Nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &user, nil
}
