package jwt

import (
	"authApp/models"
	"authApp/util"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("token expired")
		}
	}

	var user models.User
	j.db.First(&user, "id = ?", token.Claims.(jwt.MapClaims)["sub"])
	if user.Id == uuid.Nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
