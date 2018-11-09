package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestGetAll Fetches all entries
func TestGetAll(t *testing.T) {
	fmt.Println("Testing GetAll")
	// Create a request to serveHTTP. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := httptest.NewRequest("GET", "http://127.0.0.1:8000/books", nil)
	getResponse(t, req)

}

//GetTest : Fetches a specified entry
func TestGet(t *testing.T) {
	fmt.Println("Testing GET")
	// Create a request to serveHTTP. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := httptest.NewRequest("GET", "http://127.0.0.1:8000/books/1", nil)
	getResponse(t, req)
}

//TestPost
func TestPost(t *testing.T) {
	fmt.Println("Testing POST")
	//First create a test POST body
	testBook := Book{ID: "3", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)

	// Create a request to pass to our handler. We have a POST to make, so we'll
	// pass body as the third parameter.
	req := httptest.NewRequest("POST", "http://127.0.0.1:8000/books", body)
	getResponse(t, req)

}

//TestUpdate
func TestUpdate(t *testing.T) {
	fmt.Println("Testing Update")
	//First create a test POST body
	testBook := Book{ID: "10", Name: "Pride and Prejudice", Author: "Jane Austen", Count: 5}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)

	// Create a request to pass to our handler. We have a POST to make, so we'll
	// pass body as the third parameter.
	req := httptest.NewRequest("UPDATE", "http://127.0.0.1:8000/books/1", body)
	getResponse(t, req)

}

//DeleteTest : Deletes a specified entry
func TestDelete(t *testing.T) {
	fmt.Println("Testing DELETE")
	// Create a request to serveHTTP. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := httptest.NewRequest("DELETE", "http://127.0.0.1:8000/books/2", nil)
	getResponse(t, req)
}

func getResponse(t *testing.T, req *http.Request) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	res := httptest.NewRecorder()
	//Now the response request pair is served via http
	router.ServeHTTP(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	} else {
		fmt.Println("Response code is", status)
		printStruct(res.Body.String())
	}
}

func printStruct(responseBody string) {
	var books []Book
	json.Unmarshal([]byte(responseBody), &books)
	fmt.Println("Response body returned:")
	for _, bk := range books {
		fmt.Println("Book ID = ", bk.ID)
		fmt.Println("Name = ", bk.Name)
		fmt.Println("Author = ", bk.Author)
		fmt.Println("Remaining", bk.Count)
		fmt.Println(" ")
	}

}

/*


 */
