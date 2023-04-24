// exercise 7 porPak book
package main

import (
	"fmt"
	"math"
)

func sqr(data float64) float64 {
	return data * data
}
func Distance(x1, y1, x2, y2 float64) (out float64) {
	out = math.Sqrt(sqr(x2-x1) + sqr(y2-y1))
	return out
}
func main() {
	var x1, x2, y1, y2 float64
	fmt.Print("Please Enter X coordinate for point 1:")
	fmt.Scan(&x1)
	fmt.Print("Please Enter Y coordinate for point 1:")
	fmt.Scan(&y1)
	fmt.Print("Please Enter X coordinate for point 2:")
	fmt.Scan(&x2)
	fmt.Print("Please Enter Y coordinate for point 2:")
	fmt.Scan(&y2)
	dec := Distance(x1, y1, x2, y2)
	fmt.Println("distance between points:", dec)
	dec = Distance(x1, y1, 0, 0)
	fmt.Println("distance between point1 and (0,0):", dec)
	dec = Distance(0, 0, x2, y2)
	fmt.Println("distance between point2 and (0,0):", dec)
}
