package main

import (
	"fmt"
	m "maths"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Avg(xs)
	average := average(xs)
	fmt.Println(avg)
	fmt.Println(average)
	fmt.Println("Nice")
	fmt.Println("dfgdf")
	//jhhkjsd
}

//Average counter
func average(arr []float64) float64 {
	var total float64
	for _, x := range arr {
		total += x
	}
	t := 5
	t = t + 1
	return total / float64(len(arr))
}
