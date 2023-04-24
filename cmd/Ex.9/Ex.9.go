package main

import "fmt"

func All(data []int, f func(a int) bool) bool {
	for _, v := range data {
		if !f(v) {
			return false
		}
	}
	return true
}
func Any(data []int, f func(a int) bool) bool {
	for _, v := range data {
		if f(v) {
			return true
		}
	}
	return true
}
func Folding(data []int, f func(out1 int, out int) int, base int) int {
	out := base
	for _, v := range data {
		out = f(v, out)
	}
	return out
}
func main() {
	data := []int{20, 10, 16, 15, 26, 45, 2, 1, 19, 81, 5, 12}
	allData := All(data, func(a int) bool {
		if a%5 == 0 {
			return true
		}
		return false
	})
	fmt.Println("All elements are divisible by 5", allData)
	anyData := Any(data, func(a int) bool {
		if a%3 == 0 {
			return true
		}
		return false
	})
	fmt.Println("Any elements are divisible by 3", anyData)

	outData := Folding(data, func(out1 int, out int) int {
		return out1 + out
	}, 5)
	fmt.Println("Folding value =", outData)
}
