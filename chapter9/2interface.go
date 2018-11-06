package main

import "fmt"

type animal interface {
	Speak() string
}

type dog struct {
	color string
}
type cat struct {
	color string
}

func (d dog) Speak() string {
	return "Wooff!"
}

func (c *cat) Speak() string {
	return "Meeww!"
}

func hain() {
	dg := dog{"black"}
	ct := cat{"white"}
	animals := []animal{dg, &ct}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal)
	}

	fmt.Println("Program Ended")
}
