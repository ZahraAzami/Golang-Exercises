// ex.17 porPake book 4 function max .min . Absolute value . Even and odd with Conditional operator.

/*
What is the ternary operator?
The ternary operator is the ?: operator used for the conditional value.
Here is an example of how one would use it in code.
v = f > 0 ? 1 : 0     // if f > 0 then v is 1 else v is 0
This operator is absent in Golang.
Why is it absent in Go?
The reason is a simple design choice.
The operator although once understood, is an easy one is in fact a complicated construct
for someone who is new to code. The Go programming language chose the simple approach of if-else.
This is a longer version of the operator, but it is more readable.
*/
package main

import (
	"fmt"
	"math"
)

func Max(in1, in2 int) int {
	if in1 >= in2 {
		return in1
	}
	return in2
}
func Min(in1, in2 int) int {
	if in1 <= in2 {
		return in1
	}
	return in2
}

func abs(in float64) float64 {
	return math.Abs(in)
}

func isOdd(in1 int) string {
	if in1%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {

	out := Max(15, -36)
	fmt.Println("max value:", out)
	out = Min(120, -658)
	fmt.Println("Min value :", out)
	out1 := abs(-87.5)
	fmt.Println("absolute value :", out1)
	out2 := isOdd(45)
	fmt.Println("even or odd :", out2)

}
