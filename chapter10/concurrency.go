package main

import (
	"fmt"
	"time"
)

func numerals() {
	for i := 10; i < 20; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}

}
func alphabet() {
	for i := 0; i < 10; i++ {
		fmt.Println("a")
		time.Sleep(time.Millisecond * 500)
	}

}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {

	go f(0)
	numerals()
	alphabet()

}
