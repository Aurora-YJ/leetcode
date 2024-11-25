package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1 , str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}
	if len(str1) == 0 {
		return len(str2)
	}
	if len(str2) == 0 {
		return len(str1)
	}

	c := 0
	mapp := make(map[rune]bool)
	mapp1 := make(map[rune]bool) 

	for _, char := range str1 {
		mapp[char] = true	
	}
	for _, st := range str2{
		mapp1[st] = true
	}

	for char := range mapp{
		if !mapp1[char]{
			c++
		}
	}
	for char := range mapp1{
		if !mapp[char]{
			c++
		}
	}

	return c
}