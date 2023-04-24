package main

import "fmt"

func Factorial(data int) float64 {
	if data == 1 {
		return 1
	} else if data == 2 {
		return 2
	} else {
		return float64(data) * Factorial(data-1)
	}
}
func Fibonacci(data int) float64 {
	if data == 0 {
		return 0
	} else if data == 1 || data == 2 {
		return 1
	} else {
		return Fibonacci(data-1) + Fibonacci(data-2)
	}

}
func main() {
	Factorial := Factorial(4)
	fmt.Println("Factorial:", Factorial)
	out3 := Fibonacci(8)
	fmt.Println("Fibonacci:", out3)
}
