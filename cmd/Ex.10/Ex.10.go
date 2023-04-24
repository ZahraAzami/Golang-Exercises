package main

import (
	"bytes"
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
