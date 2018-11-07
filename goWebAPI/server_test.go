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
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(testBook)
	res, err := http.Post("http://127.0.0.1:8000/book/0", "application/json; charset=utf-8", b)
	if err != nil {
		t.Error("No response")
		return
	}

	sc := res.StatusCode
	if sc != 200 {
		log.Println("Couldn't post. Status Code = " + strconv.Itoa(sc))
	}

	var bk []Book
	//log.Println()

	//fmt.Println()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	} else {
		//fmt.Println(string(data))
	}
	err = json.Unmarshal(data, &bk)

	fmt.Println(bk)

}
