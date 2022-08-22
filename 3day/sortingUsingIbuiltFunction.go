package main

import "fmt"

func main() {
	var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	funcResult := sliceArray(arr)
	fmt.Println(funcResult)
}

func sliceArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
