package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestPost(t *testing.T) {
	testBook := Book{ID: "0", Name: "Catch 22", Author: "Joseph Heller", Count: 43}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(testBook)
	res, err := http.Post("http://127.0.0.1:8000/books/0", "application/json; charset=utf-8", body)
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
	} else {
		fmt.Println("")
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
	} else {
		fmt.Println("")
	}
	err = json.Unmarshal(data, &bks)

	fmt.Println("GET success. Returned:", bks)

}
