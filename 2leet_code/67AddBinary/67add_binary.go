package main

import (
	"fmt"
	"math/big"
)

func main() {
	res := addBinary(
		"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011")

	fmt.Println("res:", res)
}

func addBinary(a string, b string) string {
	n1 := replacing(a)
	n2 := replacing(b)

	sum := new(big.Int).Add(n1, n2)

	return intobinar(sum)
}

func replacing(s string) *big.Int {
	res := big.NewInt(0)
	two := big.NewInt(2)

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			exp := int64(len(s) - 1 - i)
			pow := power(two, big.NewInt(exp))
			res.Add(res, pow)
		}
	}
	return res
}

func power(base, exponent *big.Int) *big.Int {
	result := new(big.Int)
	result.Exp(base, exponent, nil)
	return result
}

func intobinar(n *big.Int) string {
	zero := big.NewInt(0)
	two := big.NewInt(2)
	result := ""

	val := new(big.Int).Set(n) 

	if val.Cmp(zero) == 0 {
		return "0"
	}

	for val.Cmp(zero) > 0 {
		remainder := new(big.Int)
		val.QuoRem(val, two, remainder) 
		result = remainder.String() + result
	}
	return result
}
