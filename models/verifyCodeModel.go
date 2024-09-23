package models

import (
	"time"

	"github.com/google/uuid"
)

type VerificationCode struct {
	ID        uuid.UUID `bson:"id" gorm:"primaryKey"`
	Code      string    `bson:"code" gorm:"not null"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	UserID    uuid.UUID `bson:"user_id" gorm:"not null"`
	User      User      `bson:"user" gorm:"foreignKey:UserID"`
}
