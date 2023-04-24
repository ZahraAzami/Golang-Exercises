package main

import "fmt"

func main() {
	sum := 0.0  // static variable
	sumX := 0.0 // static variable
	mean := func(x float64) float64 {
		sum = sum + 1
		sumX = sumX + x
		return sumX / sum
	}
	fmt.Printf("mean = %v \n", mean(1))
	fmt.Printf("mean = %v \n", mean(1))
	fmt.Printf("maen = %v \n", mean(4))
	fmt.Printf("maen = %v \n", mean(10))
	fmt.Printf("mean = %v \n", mean(11))
}
