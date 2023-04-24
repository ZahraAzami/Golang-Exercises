package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	newData := IntSet{
		data: make(map[int]bool),
	}
	return newData
}

func (in *IntSet) Add(value int) {
	in.data[value] = true
}
func (in *IntSet) Remove(value int) bool {
	if in.data[value] == true {
		delete(in.data, value)
		return true
	}
	return false
}
func (in *IntSet) Contains(value int) bool {
	return in.data[value]
	// one state other
	return in.data[value] == true
	//two state other
	_, Ok := in.data[value]
	return Ok

}
func Union(d1, d2 IntSet) IntSet {
	newUnion := NewIntSet()
	for k, v := range d1.data {
		newUnion.data[k] = v
	}
	for k, v := range d2.data {
		newUnion.data[k] = v
	}
	return newUnion
}
func Intersect(d1, d2 IntSet) IntSet {
	newUnion := NewIntSet()
	for k, _ := range d1.data {
		if d2.Contains(k) {
			newUnion.Add(k)
		}
	}
	return newUnion
}
func (in *IntSet) Len() int {
	return len(in.data)
}
func (in IntSet) String() string {
	var buff bytes.Buffer
	buff.WriteString("{")
	i := 0
	for k := range in.data {
		buff.WriteString(strconv.Itoa(k))
		if in.Len()-1 == i {
			break
		}
		buff.WriteString(", ")
		i++
	}
	buff.WriteString("}")
	return buff.String()
}
func (in *IntSet) Add2(value ...int) {
	for _, v := range value {
		in.data[v] = true
	}
}
func main() {

	data1 := IntSet{
		data: map[int]bool{
			1: true,
			2: true,
			6: true,
			8: true,
			9: true,
			7: true,
		},
	}
	data2 := IntSet{
		data: map[int]bool{
			1: true,
			3: true,
			5: true,
			7: true,
		},
	}
	fmt.Println("union:", Union(data1, data2))
	fmt.Println("d1", data1)
	fmt.Println("InterSet:", Intersect(data1, data2))
	fmt.Println("St:", data1.String())
	data1.Add(2)
	fmt.Println("add:", data1)
}
