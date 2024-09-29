package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationCode struct {
	ID        uuid.UUID `bson:"id" gorm:"type:uuid;primaryKey"`
	Code      string    `bson:"code" gorm:"not null"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	UserID    uuid.UUID `bson:"user_id" gorm:"not null"`
	User      User      `bson:"user;omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (code *VerificationCode) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	code.ID = uuid.New()
	code.Code = strconv.Itoa(int(GenerateVerifyCode()))
	return
}

func GenerateVerifyCode() uint {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return uint(r.Intn(900000) + 100000)
}
