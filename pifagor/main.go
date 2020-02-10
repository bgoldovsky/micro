package main

import (
	"fmt"
)

func main() {
	// 3! = 6 combines
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	isPifagor := isPifagor(a, b, c)
	fmt.Println(isPifagor)
}

func isPifagor(a, b, c int) string {
	if a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b {
		return "yes"
	}
	return "no"
}
