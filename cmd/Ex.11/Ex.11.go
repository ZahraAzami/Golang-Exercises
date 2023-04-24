package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type BitVector struct {
	data []byte
	lim  uint
}

func NewBitvector(l uint) BitVector {
	return BitVector{make([]byte, l/8+1), l}
}

func (b *BitVector) Add(data1 uint) bool {
	if int(data1) > int(b.lim) {
		return false
	}
	index := data1 / 8
	shifted := byte(1) << byte(data1%8)
	b.data[index] |= shifted
	return true
}

func (b *BitVector) Contains(data uint) bool {
	if int(data) > int(b.lim) {
		return false
	}
	byteIndex := data / 8
	bitIndex := byte(data % 8)
	mask := byte(1) << bitIndex
	ret := b.data[byteIndex] & mask
	if ret == 0 {
		return false
	}
	return true
}

func (b *BitVector) Remove(data uint) bool {
	if int(data) > int(b.lim) {
		return false
	}
	byteIndex := data / 8
	bitIndex := byte(data % 8)
	mask := byte(1) << bitIndex
	b.data[byteIndex] &= ^(mask)
	return true
}

func (b BitVector) String() string {
	var buff bytes.Buffer
	hasElment := false
	buff.WriteString("{")
	for i := 0; i < int(b.lim); i++ {
		if b.Contains(uint(i)) == true {
			if !hasElment {
				buff.WriteString(strconv.Itoa(i))
				hasElment = true
			} else {
				buff.WriteString(", ")
				buff.WriteString(strconv.Itoa(i))
			}

		}

	}
	buff.WriteString("}")
	return buff.String()
}
func main() {
	dr := NewBitvector(18)
	dr.Add(12)
	dr.Add(10)
	dr.Add(14)
	dr.Add(9)
	fmt.Println("add;", dr.data)
	fmt.Println("Contains;", dr.Contains(11))
	fmt.Println("Remove ;", dr.Remove(10))
	fmt.Println("add;", dr.data)
	dr.Add(8)
	fmt.Println("add;", dr.data)
	fmt.Println("Remove ;", dr.Remove(8))
	fmt.Println("add;", dr.data)
	fmt.Println("St:", dr.String())

}
