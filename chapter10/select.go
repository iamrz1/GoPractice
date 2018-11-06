package main

import (
	"fmt"
	"time"
)

func func1(c1 chan string) {
	for {
		c1 <- "from 1"
		time.Sleep(time.Second * 1)
	}
}
func func2(c2 chan string) {
	for {
		c2 <- "from 2"
		time.Sleep(time.Second * 2)
	}
}

func func3(c1, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func1(c1)

	go func2(c2)

	go func3(c1, c2)

	var input string
	fmt.Scanln(&input)
}
