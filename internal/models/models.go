package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GUID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Email       string    `gorm:"size:255;uniqueIndex;not null"`
	IPAddress   string    `gorm:"size:255;not null"`
	TokenHashed string    `gorm:"size:255;uniqueIndex;not null"`
}

type Token struct {
	gorm.Model
	UserID       uuid.UUID `gorm:"type:uuid;foreignKey:UserID;references:GUID"`
	AccessToken  string    `gorm:"size:255;uniqueIndex;not null"`
	RefreshToken string    `gorm:"size:255;uniqueIndex;not null"`
}
