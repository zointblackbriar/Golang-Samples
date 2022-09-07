package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8000
	
	http.HandleFunc("/sampleendpoint", sampleEndpoint)
	log.Printf("Server is starting on port %v\n", port) 
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func sampleEndpoint(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Hello Message from the app \n") 
}	
