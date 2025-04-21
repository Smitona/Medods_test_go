package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/Smitona/Medods_test_go/internal/models"
)

var db = ConnectToDB()
var users models.User
var tokens models.Token

func generateTokenPairHandler(w http.ResponseWriter, r *http.Request) {
	// Generate Access Token (JWT) and Refresh Token (random bytes, base64 encoded)

	var input struct {
		GUID      string `json:"guid"`
		IPAddress string `json:"ip_address"`
	}

	if input.GUID == "" {
		http.Error(w, "GUID is required", http.StatusBadRequest)
		return
	}

	IPAddress := strings.Split(r.RemoteAddr, ":")[0]

	accessToken, err := generateAccessToken(input.GUID, IPAddress)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	db.Model(&users).
		Where("GUID = ?", input.GUID).
		Update("TokenHashed", refreshToken)

	userGUID, _ := uuid.Parse(input.GUID)

	TokenPair := models.Token{
		UserID:       userGUID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := db.Create(&TokenPair).Error; err != nil {
		http.Error(w, "Failed to save tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// refresh token logic here
	var input struct {
		RefreshToken string `json:"refresh_token"`
		IPAddress    string `json:"ip_address"`
	}

	IPAddress := strings.Split(r.RemoteAddr, ":")[0]

	if input.RefreshToken == "" {
		http.Error(w, "Refresh token is required", http.StatusBadRequest)
		return
	}

	TokenPair := db.Where("refresh_token = ?", input.RefreshToken).First(&tokens)
	if TokenPair.Error != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	if err := db.Where("GUID = ?", tokens.UserID).First(&users).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	newAccessToken, err := generateAccessToken(users.GUID.String(), IPAddress)
	if err != nil {
		http.Error(w, "Failed to generate new access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := generateRefreshToken()
	if err != nil {
		http.Error(w, "Failed to generate new refresh token", http.StatusInternalServerError)
		return
	}

	if err := db.Model(&users).Update("TokenHashed", newRefreshToken).Error; err != nil {
		http.Error(w, "Failed to update user hashed_token", http.StatusInternalServerError)
		return
	}

	tokens.AccessToken = newAccessToken
	tokens.RefreshToken = newRefreshToken
	if err := db.Save(&tokens).Error; err != nil {
		http.Error(w, "Failed to update tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
