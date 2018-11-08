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
)

func TestPost(t *testing.T) {
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

	fmt.Println("POST success. Returned:", bks)

}
func TestGet(t *testing.T) {
	//time.Sleep(time.Second * 3)
	testBook := Book{ID: "0", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)

	res, err := http.Get("http://127.0.0.1:8000/books")

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

//DeleteTest : Deletes a specified entry
func TestDelete(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	//req, err := http.NewRequest("GET", "/health-check", nil)

	//b := new([]byte)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(DeleteBook)
	rqst := httptest.NewRequest("DELETE", "http://127.0.0.1:8000/books/1", nil)
	DeleteBook(rr, rqst)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, rqst)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"alive": true}`
	// if req, err := http.NewRequest("DELETE", "/books/1", nil) != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
	fmt.Println(rr.Body.String())

}
