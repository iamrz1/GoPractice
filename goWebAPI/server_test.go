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
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

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
	testBook := Book{ID: "0", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)

	// Create a request to pass to our handler. We have a POST to make, so we'll
	// pass body as the third parameter.
	req := httptest.NewRequest("POST", "http://127.0.0.1:8000/books/3", body)
	getResponse(t, req)
	req = httptest.NewRequest("POST", "http://127.0.0.1:8000/books/2", body)
	getResponse(t, req)
}

//DeleteTest : Deletes a specified entry
func TestDelete(t *testing.T) {
	fmt.Println("Testing DELETE")
	// Create a request to serveHTTP. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req := httptest.NewRequest("DELETE", "http://127.0.0.1:8000/books/1", nil)
	getResponse(t, req)
}

func getResponse(t *testing.T, req *http.Request) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	res := httptest.NewRecorder()
	//Now the response request pair is served via http
	router.ServeHTTP(res, req)

	// Check the status code is what we expect.
	statusAssesment(res.Code, t, res)

}

func statusAssesment(status int, t *testing.T, res *httptest.ResponseRecorder) {

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	} else {
		fmt.Println("Response code is", http.StatusOK)
		printStruct(res.Body.String())
	}
}

func printStruct(responseBody string) {
	var books []Book
	json.Unmarshal([]byte(responseBody), &books)
	fmt.Println("Response body returned = ")
	for _, bk := range books {
		fmt.Println("Book ID = ", bk.ID)
		fmt.Println("Name = ", bk.Name)
		fmt.Println("Author = ", bk.Author)
		fmt.Println("Remaining", bk.Count)
		fmt.Println(" ")
	}

}
