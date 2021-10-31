package main

import (
	"fmt"
	"hydra/hydra/hlogger"
	"net/http"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra Web Service")


	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(w, "Welcome to the Hydra Software System")
	
	logger.Println(w, "Received and http request on root url")
}
