package main

import "fmt"

func main() {
	s := "abcdefghit"
	k := 3
	fill := byte('x')
	res := divideString(s, k, fill)
	fmt.Println("result==>", res)
}

func divideString(s string, k int, fill byte) []string {
	c := getThediv(len(s), k)

	var result []string
	res := ""
	help := 0
	for i := 0; i < c; i++ {
		for j := 0; j < k; j++ {
			if help > len(s)-1 {
				res += string(fill)
			} else {
				res += string(s[help])
				help++
			}
		}
		result = append(result, res)
		res = ""
	}

	return result
}

func getThediv(lenString, lenK int) int {
	if lenString%lenK == 0 {
		return lenString / lenK
	} else {
		return (lenString / lenK) + 1
	}
}
