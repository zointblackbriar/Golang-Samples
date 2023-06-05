package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func headerregex() {
	router := mux.newRouter()
	route := router.NewRoute().HeadersRegexp("Accept", "html")

	firstRequest, _ := http.newRequest("GET", "sample.com", nil)
	
}