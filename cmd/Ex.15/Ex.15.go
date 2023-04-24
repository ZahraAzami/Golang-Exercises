package main

import "fmt"

type mobile struct {
	brand string
}
type Laptop struct {
	cpu string
}
type toaster struct {
	Amount int
}
type kettle struct {
	quantity string
}
type socket struct{}

func (m mobile) Draw(power int) {
	fmt.Printf("%T-> brand :%s , power : %d \n", m, m.brand, power)
}
func (L Laptop) Draw(power int) {
	fmt.Printf("%T-> CPU :%s , power : %d \n", L, L.cpu, power)
}
func (t toaster) Draw(power int) {
	fmt.Printf("%T-> Amount :%d , power : %d \n", t, t.Amount, power)
}
func (k kettle) Draw(power int) {
	fmt.Printf("%T-> quantity :%s , power : %d \n", k, k.quantity, power)
}

type PowerDrawer interface {
	Draw(power int)
}

func (socket) plung(divace PowerDrawer, power int) {
	divace.Draw(power)
}
