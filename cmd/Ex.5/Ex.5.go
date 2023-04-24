package main

import "fmt"

func main() {
	var data1 int
	fmt.Print("Please Enter data1:")
	fmt.Scan(&data1)
	var data2 int
	fmt.Print("Please Enter data2:")
	fmt.Scan(&data2)
	out := data1 + data2
	fmt.Println("this is operation + : ", out)
	out = data1 - data2
	fmt.Println("this is operation - : ", out)
	out = data1 * data2
	fmt.Println("this is operation * : ", out)
	out = data1 / data2
	fmt.Println("this is operation / : ", out)
	out = data1 % data2
	fmt.Println("this is operation % : ", out)
	var data3 int
	fmt.Print("Please Enter data3:")
	fmt.Scan(&data3)
	out += data3
	fmt.Println("this is operation += : ", out)
	out -= data3
	fmt.Println("this is operation -= : ", out)
	out *= data3
	fmt.Println("this is operation *= : ", out)
	out /= data3
	fmt.Println("this is operation /= : ", out)
	out %= data3
	fmt.Println("this is operation %= : ", out)

}
