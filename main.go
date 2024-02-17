package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/books/{book_id}/page/{page_id}", func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		fmt.Fprintf(w, "Hello, you've requested: %s\n", params["book_id"])
	})

	// r.HandleFunc("/books", ReadAllBooks).Methods("POST")
	// r.HandleFunc("/books", CreateBook).Methods("POST")
	// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	// SUBDOMAINS
	// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	fmt.Println("Server started on port 8080:")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
