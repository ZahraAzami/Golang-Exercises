package main

import "fmt"

func All(data []int, f func(a int) bool) bool {
	for _, v := range data {
		if !f(v) {
			return false
		}
	}
	return true
}
func Any(data []int, f func(a int) bool) bool {
	for _, v := range data {
		if f(v) {
			return true
		}
	}
	return true
}
func Folding(data []int, f func(out1 int, out int) int, base int) int {
	out := base
	for _, v := range data {
		out = f(v, out)
	}
	return out
}
func Partition(data []int, f func(com int) bool) ([]int, int) {
	j := 0
	for i, _ := range data {
		if f(data[i]) {
			data[i], data[j] = data[j], data[i]
			j++
		}
	}
	return data, j
}
func MergeSlice(data1, data2 []int, f func(com1, com2 int) bool) []int {
	lenTotal := len(data1) + len(data2)
	data := make([]int, 0, lenTotal)
	for i := 0; i < len(data); i++ {
		if f(data1[i], data2[i]) {
			data = append(data, data[i])
		}
	}

	return data
}
func QuickSort(data []int) {
	firstGrater := 0
	pivot := data[len(data)-1]
	subData := data[:len(data)-1]
	subData, firstGrater = Partition(subData, func(pre int) bool {
		return pivot > pre
	})
	data[firstGrater], data[len(data)-1] = data[len(data)-1], data[firstGrater]
}

func main() {
	data := []int{20, 10, 16, 15, 26, 45, 2, 1, 19, 81, 5, 12, 35}
	allData := All(data, func(a int) bool {
		if a%5 == 0 {
			return true
		}
		return false
	})
	fmt.Println("All elements are divisible by 5", allData)
	anyData := Any(data, func(a int) bool {
		if a%3 == 0 {
			return true
		}
		return false
	})
	fmt.Println("Any elements are divisible by 3", anyData)
	outData := Folding(data, func(out1 int, out int) int {
		return out1 + out
	}, 5)
	fmt.Println("Folding value =", outData)
	partData, Num := Partition(data, func(com int) bool {
		if com%2 == 0 {
			return true
		}
		return false
	})
	fmt.Printf("Separation data = %v ,Number data = %v\n", partData, Num)
	QuickSort(data)
	fmt.Println("quick Sort data = ", data)
	Data2 := []int{12, 45, 78, 65, 9, 53, 44, 54, 94}
	mergeData := MergeSlice(data, Data2, func(com1, com2 int) bool {
		if com1 != com2 {
			return true
		}
		return false
	})
	fmt.Printf("Merge Slice data = %v", mergeData)
}
