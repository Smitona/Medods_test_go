package main

import (
    "log"
    "net/http"
)

func main() {
	// init routers
	router := AuthRouters()
	// init server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
