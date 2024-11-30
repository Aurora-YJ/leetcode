package main

import "fmt"

func main() {
	arr := []int{5, 3, 8, 6, 2}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) []int {

	for i:=1; i < len(arr); i++ {
		key := arr[i]
		j := i-1
		for j >=0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
