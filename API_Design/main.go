package main

//Basic API design

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Frau Zorokino", Year: "2013"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//log.Println("Gets all books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets one book")
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"]) //string conversion - we have to convert id param to integer
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
	log.Println(params)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adds one book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) { //response and request
	var book Book

	json.NewDecoder(r.Body).Decode(&book) //json format decode
	for i, item := range books {
		if item.ID == book.ID { //ID Check
			books[i] = book
		}
	}
	log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removes a book")
}
