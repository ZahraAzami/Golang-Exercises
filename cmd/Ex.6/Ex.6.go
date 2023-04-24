// ex.14 SizeOf in int8,int16,int32,int64,float32, float32, string , rune , byte , complex64
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var Int8 int8
	var Int16 int16
	var Int32 int32
	var Int64 int64
	var Float32 float32
	var Float64 float64
	var String string
	var Rune rune
	var Byte byte
	var Complex64 complex64
	fmt.Println("Table 1. Data sizes using sizeOf(variable)")
	fmt.Println("    Data Type               memory used (bytes)")
	fmt.Println("--------------------        ----------------------")

	fmt.Printf("    TYPE = %T                        %d\n    ", Int8, unsafe.Sizeof(Int8))
	fmt.Printf("TYPE = %T                       %d\n    ", Int16, unsafe.Sizeof(Int16))
	fmt.Printf("TYPE = %T                       %d\n    ", Int32, unsafe.Sizeof(Int32))
	fmt.Printf("TYPE = %T                       %d\n    ", Int64, unsafe.Sizeof(Int64))
	fmt.Printf("TYPE = %T                     %d\n    ", Float32, unsafe.Sizeof(Float32))
	fmt.Printf("TYPE = %T                     %d\n    ", Float64, unsafe.Sizeof(Float64))
	fmt.Printf("TYPE = %T                     %d\n    ", String, unsafe.Sizeof(String))
	fmt.Printf("TYPE = %T                       %d\n    ", Rune, unsafe.Sizeof(Rune))
	fmt.Printf("TYPE = %T                       %d\n    ", Byte, unsafe.Sizeof(Byte))
	fmt.Printf("TYPE = %T                   %d\n    ", Complex64, unsafe.Sizeof(Complex64))

}
