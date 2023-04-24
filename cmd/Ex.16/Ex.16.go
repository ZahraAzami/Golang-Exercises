package main

import (
	"fmt"
	"time"
)

func TimerCounter(n int) {
	for i := n; i >= 0; i-- {
		if i == n {
			fmt.Println(i)
			continue
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
func Timelimit() {
	var T time.Duration
	for T <= 4*time.Second {
		time.Sleep(33 * time.Millisecond)
		fmt.Println("*")
		T += 33 * time.Millisecond
	}
}

func callLater(sec int, f func()) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println("", f)
}
