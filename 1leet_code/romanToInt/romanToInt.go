package main

import "fmt"

func main() {
	str :=  "MCMXCIVVI"
	//621
	nm := romanToInt(str)
	fmt.Println(nm)
}

func romanToInt(s string) int {
	helper := 0
	mmap := make(map[string]int)
	mmap["I"] = 1
	mmap["V"] = 5
	mmap["X"] = 10
	mmap["L"] = 50
	mmap["C"] = 100
	mmap["D"] = 500
	mmap["M"] = 1000
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && mmap[string(s[i])] < mmap[string(s[i+1])] {
			helper += mmap[string(s[i+1])] - mmap[string(s[i])]
			if i+1 == len(s) - 1 {
				break
			}
			i++
		} else if i != len(s)-1 && mmap[string(s[i])] > mmap[string(s[i+1])] {
			helper += mmap[string(s[i])]
		} else if i != len(s)-1 && mmap[string(s[i])] == mmap[string(s[i+1])] {
			fmt.Println("")
			helper += mmap[string(s[i])]
			
		}

		if i == len(s)-1 {
			helper += mmap[string(s[i])]
		}
	}
	return helper
}
