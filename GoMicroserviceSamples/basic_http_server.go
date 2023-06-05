package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//Zuweisen einen Portnummer
	port := 8000
	
	// http.HandleFunc("/sampleendpoint", sampleEndpoint)
	// log.Printf("Server is starting on port %v\n", port) 
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)) {
		log.Prinf("Anfrage wurde angenommen")
		w.Write([]byte("Hallo Welt"))
	}

	serverBeispiel := &http.Server {
		Addr: fmt.Sprintf(":%d", port),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1 MB
	}



}

func sampleEndpoint(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Hello Message from the app \n") 
}	
