package main

import "fmt"

func main() {
	res := addBinary("11", "1") // "100"
	fmt.Println("res:", res)
}

func addBinary(a string, b string) string {

	n1 := replacing(a)

    n2 := replacing(b)

	tt := n1 + n2

	fmt.Println("SDFSDG", n1, n2 ,tt)
	res :=  intobinar(tt)


	return res
}


func replacing(s string) int {

	res := 1
	help := 0

	for range s {
		res += binartoint(help)

		help++
	}

	return res
}


func binartoint(n int) int {

	res := 1

	for range n  {
		res = res * n 
	}

	return res

}


func intobinar(n int) string {


	if n == 0 {
		return ""	
	} 
		
	return intobinar(n/2) + itoa(n%2)

}


func itoa(i int) string {

	tt := string(i + '0')

	return tt

}