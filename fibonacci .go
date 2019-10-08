package main

func fibonacci(n int) []int {
    a := 0
    b := 1
	// Iterate until desired position in sequence.
	
	f:= []int{0}
    for i := 0; i < n-1; i++ {
        // Use temporary variable to swap values.
        temp := a
        a = b
		b = temp + a
		f= append(f, a)
    }
    return f
}
