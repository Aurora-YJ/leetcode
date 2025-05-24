package main

import "fmt"

func main() {
	words :=  []string{"leet","code"}
	x := 'e'
	res := findWordsContaining(words, byte(x))
	fmt.Println(res)
}
func findWordsContaining(words []string, x byte) []int {
    var res []int 
    for i := 0; i < len(words) ; i++ {
        for j := 0; j < len(words[i]) ; j++ {
			if words[i][j] == x {
				res = append(res, i)
				break
			}
		}
    }
	return res
}

