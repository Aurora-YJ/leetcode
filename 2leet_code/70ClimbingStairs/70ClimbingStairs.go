package main

import "fmt"


func main() {
	var n int
	fmt.Print("Enter the number of steps: ")
	fmt.Scan(&n)

	result := climbStairs(n)
	fmt.Printf("The number of distinct ways to climb %d steps is: %d\n", n, result)
}


func climbStairs(n int) int {
	if n == 2 {
		return n
	}
	if n == 1 {
		return 1
	}
	nmb1 := 1
	nmb2 := 2
	res := 0
	for i := 3 ; i <= n ; i++ {
		res = nmb1 + nmb2
		nmb1 = nmb2 
		nmb2 = res 
	}
	return res
}