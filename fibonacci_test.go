package main

import (
	
	"testing"
)

func TestFibonacci( f *testing.T){
	fb:= fibonacci("4")

	if len(fb)!=4{
		f.Errorf("Expected length of input 4 to be 4 but got %v", len(fb))
	}
	if fb[0]!=0{
		f.Errorf("Expected Fibonacci to start with 0 but got %v", fb[0])
	}
	if fb[3]!=2{
		f.Errorf("Expected Fibonacci to start with 2 but got %v", fb[3])
	}
}