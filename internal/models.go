package main

import (
	"github.com/google/uuid"
)

type User struct {
	GUID  uuid.UUID `json:"guid"`
	Email string    `json:"email"`
}

type Token struct {
	ID           int       `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	TokenHash    string    `json:"token_hash"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
