package basic_http

import (
	"fmt"
	"log"
	"net/http"
)

func server_initialize() {
	port := 3000

	http.HandleFunc("/sampleendpoint", sampleEndpoint)
	http.HandleFunc("/sendmessage", sendMessage)
	log.Printf("Server is starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func sampleEndpoint(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Hello Message from the app \n")
}

func sendMessage(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Message from URL Endpoint \n")
}
