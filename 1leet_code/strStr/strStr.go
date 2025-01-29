package main

import "fmt"

func main() {
	nm := strStr("leetcode", "code")
	fmt.Println(nm)
}

func strStr(haystack string, needle string) int {

	f := needle[0]
	
	for j, c := range haystack {
		if byte(c) == f {

			if valid(haystack[j:], needle) {
				return j
			}

		}
	}
	return -1
}

func valid(s1 string, s2 string) bool {

	if len(s1) < len(s2) {
		return false
	}

	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
