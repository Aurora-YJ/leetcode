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

	mapp := make(map[byte]bool)
	res := ""
	c := 0
	for i := 0; i < len(s1); i++ {
		for j := c; j < len(s2); j++ {
			//fmt.Println(c)
			if s1[i] == s2[j] && !mapp[s1[i]] {
				res += string(s1[i])
				mapp[s1[i]] = true
				c = j
				break
			}
		}
	}

	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		if len(res) != len(s1) {
			fmt.Println(0)
			return
		}
		if res[i] != s1[i] {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(1)
}