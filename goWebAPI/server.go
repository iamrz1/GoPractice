package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book (struct) : inventory of books in the shop
type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Count  int    `json:"count,omitempty"`
	//Author *Author `json:"address,omitempty"`
}

// Author (struct) : details of Author
// type Author struct {
// 	City  string `json:"city,omitempty"`
// 	State string `json:"state,omitempty"`
// }

var books []Book
var router = mux.NewRouter()
var ids map[int]int
var count = 2

func init() {
	books = append(books, Book{ID: "1", Name: "Pride and Prejudice", Author: "Jane Austen", Count: 5})
	books = append(books, Book{ID: "2", Name: "Things fall apart,", Author: "Chinua Achebe", Count: 9})
	ids = make(map[int]int)
	ids[1] = 1
	ids[2] = 1

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
}

// GetBooks : Display all books from the books variable
func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
	//fmt.Println(w.Header)
}

// GetBook : get a single book by id
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("mux id = ", vars["id"])
	for _, item := range books {
		if item.ID == vars["id"] {
			fmt.Println("Match found")
			var b []Book
			b = append(b, item)
			json.NewEncoder(w).Encode(b)
			return
		}
	}
	//if no match is found,
	// return an empty book struct
	fmt.Println("This shoulnt be executed")

	//json.NewEncoder(w).Encode(&Book{})
}

// CreateBook : create a new book entry
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// extract id from URL
	id, _ := strconv.Atoi(r.URL.Path[len("/books/"):])

	if ids[id] == 1 {
		fmt.Println("Duplicate Entry. ID = ", id, " exists")
		json.NewEncoder(w).Encode(books)
		return
	}
	count++
	ids[count] = 1

	vars := mux.Vars(r)
	var book Book
	// get the struct equivalent of the json
	//and save that to our book variable
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	book.ID = vars["id"]
	//fmt.Println("Added ID = ", vars["id"])
	//add the new entry to our existing book entries
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

// DeleteBook : Delete a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(len(books))

	vars := mux.Vars(r)
	fmt.Println("========", vars["id"])
	fmt.Println("========", r.URL.Path)
	id, idErr := strconv.Atoi(r.URL.Path[len("/books/"):])
	fmt.Println("id = ", id)
	if idErr != nil {
		//w.Write("URL Error.")
	}
	for index, item := range books {
		if item.ID == strconv.Itoa(id) {
			fmt.Println("IDs are equal")
			books = append(books[:index], books[index+1:]...)
			break
		}
		// json.NewEncoder(w).Encode(books)
	}
	fmt.Println(len(books))
	json.NewEncoder(w).Encode(books)
}

// main function to boot up everything
func main() {

	log.Fatal(http.ListenAndServe(":8000", router))
}
