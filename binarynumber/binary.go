package main

import (
	"fmt"
	"math"
)

func main() {
	v := "Hello world"
	fmt.Println(v)
	for _, e := range v {
		fmt.Printf("%4v -> %[1]b\n", e)
	}

	fmt.Println(binToDec(1001000))
}

func binToDec(n int) int {
	var decNum int
	for i := 0; n != 0; i++ {
		rMost := n % 10
		n = n / 10
		decNum = decNum + rMost * int(math.Pow(2, float64(i)))
	}

	return decNum
}
