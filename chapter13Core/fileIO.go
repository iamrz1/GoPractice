package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
		return
	}
	str := string(bs)
	fmt.Println(str)
	barr := []byte("TEST LINE. ")
	in := append(bs, barr...)
	errmsg := ioutil.WriteFile("test.txt", in, 0644)
	if errmsg != nil {
		panic(err)
	}

}
