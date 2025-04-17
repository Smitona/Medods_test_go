package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auth/token", generateTokenHandler).Name("token").Methods("POST")
	router.HandleFunc("/auth/refresh", refreshTokenHandler).Name("refresh").Methods("POST")
	http.Handle("/", router)

	return router
}
