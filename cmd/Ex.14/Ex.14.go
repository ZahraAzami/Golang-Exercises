package main

import "math"

type InValue int

func (i InValue) Bigger(value int) bool {
	return i > InValue(value)
}
func SquareOfSum(n int) int {
	data := 0
	for i := 0; i <= n; i++ {
		data += i
	}
	data2 := math.Pow(float64(data), 2)
	data = int(data2)
	return data
}
func SumOfSquares(n int) int {
	data := 0
	var data2 float64
	for i := 0; i <= n; i++ {
		data2 = math.Pow(float64(i), 2)
		data += int(data2)
	}
	return data
}
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
