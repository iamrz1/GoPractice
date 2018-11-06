package main

import (
	"testing"
)

//Avg : Takes an array of numbers and returns their average
func TestMain(t *testing.T) {
	testVal := []float64{5, 2}
	result := average(testVal)
	if result != 4.5 {
		t.Error("Expected 3.5. Got ", result)
	}

}
