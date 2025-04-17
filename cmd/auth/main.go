package main

import (
    "log"
	"fmt"
    "net/http"
)

func main() {
	// init routers
	router := AuthRouters()
	// init server
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
