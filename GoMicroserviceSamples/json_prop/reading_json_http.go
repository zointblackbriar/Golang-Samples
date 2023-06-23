package json_prop

// package main

import (
	"encoding/json"
	_ "log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string
}

func main() {
	port := 8000

	http.HandleFunc("/helloworld", handlerFunc)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{Message: "hello message from the server"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Oops")
	}
}
