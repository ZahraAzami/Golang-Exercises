package main

import "fmt"

const SecInMin = 60 // Global Constant
func main() {
	const MinInHour = 60 // Local Constant
	var hour, minutes, Second int
	var totalSec int
	fmt.Println("Please Enter hours")
	fmt.Scan(&hour)
	fmt.Println("Please Enter Minutes")
	fmt.Scan(&minutes)
	fmt.Println("Please Enter second")
	fmt.Scan(&Second)
	totalSec = ((hour*MinInHour + minutes) * SecInMin) + Second
	fmt.Printf("%v second since midnight", totalSec)
}
