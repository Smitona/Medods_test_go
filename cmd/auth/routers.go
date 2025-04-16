package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routers() {
	router := mux.NewRouter()
	router.HandleFunc("/auth/token", generateTokenHandler).Name("token")
	router.HandleFunc("/auth/refresh", refreshTokenHandler).Name("refresh")
	http.Handle("/", router)
}
