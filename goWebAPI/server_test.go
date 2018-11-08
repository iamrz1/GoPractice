package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	return router
}

func TestLivePost(t *testing.T) {
	testBook := Book{ID: "0", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)
	res, err := http.Post("http://127.0.0.1:8000/books/3", "application/json; charset=utf-8", body)
	if err != nil {
		t.Error("POST failed. No response")
		return
	}

	sc := res.StatusCode
	if sc != 200 {
		log.Println("Couldn't post. Status Code = " + strconv.Itoa(sc))
	}

	var bks []Book
	//log.Println()

	//fmt.Println()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	err = json.Unmarshal(data, &bks)
	for _, bk := range bks {
		fmt.Println("Book ID = ", bk.ID)
		fmt.Println("Name = ", bk.Name)
		fmt.Println("Author = ", bk.Author)
		fmt.Println("Remaining", bk.Count)
		fmt.Println(" ")
	}

	fmt.Println("POST success. Returned:", bks)

}
func TestLiveGet(t *testing.T) {
	//time.Sleep(time.Second * 3)
	testBook := Book{ID: "0", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)

	res, err := http.Get("http://127.0.0.1:8000/books/")

	if err != nil {
		t.Error("GET failed. No response.")
		return
	}

	sc := res.StatusCode
	if sc != 200 {
		log.Println("Couldn't post. Status Code = " + strconv.Itoa(sc))
	}

	var bks []Book
	//log.Println()

	//fmt.Println()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(data, &bks)

	fmt.Println("GET success. Returned:", bks)

}

//TestGetAll Fetches all entries
func TestGetAll(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	res := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "http://127.0.0.1:8000/books", nil)
	//DeleteBook(rr, rqst)
	router.ServeHTTP(res, req)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("GET All= ", res.Body.String())

}

//GetTest : Fetches a specified entry

func TestGet(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	res := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "http://127.0.0.1:8000/books/1", nil)
	//DeleteBook(rr, rqst)
	router.ServeHTTP(res, req)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("GET = ", res.Body.String())

}

//TestPost
func TestPost(t *testing.T) {

	//First create a test POST body
	testBook := Book{ID: "0", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	res := httptest.NewRecorder()

	req := httptest.NewRequest("POST", "http://127.0.0.1:8000/books/3", body)
	//DeleteBook(rr, rqst)
	router.ServeHTTP(res, req)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println(res.Body.String())

}

//DeleteTest : Deletes a specified entry
func TestDelete(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	res := httptest.NewRecorder()

	req := httptest.NewRequest("DELETE", "http://127.0.0.1:8000/books/1", nil)
	//DeleteBook(rr, rqst)
	router.ServeHTTP(res, req)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println(res.Body.String())

}
