package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func generateTokenPairHandler(w http.ResponseWriter, r *http.Request) {
	// Generate Access Token (JWT) and Refresh Token (random bytes, base64 encoded)

	var input struct {
		GUID      string `json:"guid"`
		IPAddress string `json:"pt_address"`
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	// refresh token logic here

	refreshToken := "str"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"refresh_token": refreshToken,
	})
}
