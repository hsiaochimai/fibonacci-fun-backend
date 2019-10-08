package main

import (
	"errors"
	"strconv"
)

func fibonacci(s string) ([]int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return nil, errors.New("your input is not valid, please write a number greater than 0")

	}
	if n <= 0 {
		return nil, errors.New("The number you inputed is not valid")

	}
	a := 0
	b := 1
	// Iterate until desired position in sequence.

	f := []int{0}
	for i := 0; i < n-1; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
		f = append(f, a)
	}
	return f ,nil
}
