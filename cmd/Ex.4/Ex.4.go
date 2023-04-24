// exercise 8 PorPak book
package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type subInc1 struct {
	i int
}
type subInc2 struct {
	x float64
}
type subInc3 struct {
	s string
}

func (Si *subInc1) inc(i int) {
	Si.i = i + 1
}
func (Si *subInc2) inc(x float64) {
	Si.x = x + 1
}
func (Si *subInc3) inc(s string) {
	R, _ := utf8.DecodeRuneInString(s)
	R = R + 1
	data := strconv.QuoteRune(R)
	Si.s = data
}
func main() {
	data1 := subInc1{}
	data2 := subInc2{}
	data3 := subInc3{}
	fmt.Print("Please Enter Integer:")
	fmt.Scan(&data1.i)
	fmt.Print("Please Enter Float:")
	fmt.Scan(&data2.x)
	fmt.Print("Please Enter char:")
	fmt.Scan(&data3.s)
	data1.inc(data1.i)
	data2.inc(data2.x)
	data3.inc(data3.s)
	fmt.Println("Data1 = ", data1)
	fmt.Println("Data2 = ", data2)
	fmt.Println("Data3 = ", data3)
}
