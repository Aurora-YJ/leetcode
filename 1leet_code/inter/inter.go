package main

import (
	"fmt"
	"os"
)


func main(){
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	mapp := make(map[rune]bool)
	res := ""
	for _, char := range s1 {
		for _, st := range s2 {
			if char == st && !mapp[char] {
				res += string(char)
				mapp[char] = true
			}
		}
	}
	fmt.Println(res)
}