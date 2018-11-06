package main

import (
	"fmt"
)

func avg(arr []float64) {
	var total float64
	for _, values := range arr {
		total = total + values
	}
	fmt.Println(total / float64(len(arr)))
}
func main() {
	fmt.Println("what?")
	arr := make([]float64, 5)

	for i := 0; i < 5; i++ {
		var value float64
		fmt.Scanf("%f", &value)
		arr[i] = value
	}

	avg(arr)
}
