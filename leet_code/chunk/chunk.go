package main

import (
	"fmt"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {

	if size == 0 {
		fmt.Println()
		return 
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
	

	var res [][]int
	i := 0

	for i < len(slice) {
		if i + size < len(slice) {
			b := slice[i: i+size]
			res = append(res, b)
			i += size
		} else {
			b := slice[i:]
			res = append(res, b)
			break
		}
	}
	fmt.Println(res)
}
