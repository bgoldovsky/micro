package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	answer := Solve(a, b, c)
	fmt.Println(answer)
}

func Solve(a, b, c int) string {
	if a == 0 {
		return Linear(b, c)
	}
	return Quadratic(a, b, c)
}

func Linear(b, c int) string {
	if b == 0 {
		return "has no roots"
	}

	return "has 1 root"
}

func Quadratic(a, b, c int) string {
	d := b*b - 4*a*c
	if d == 0 {
		return "has 1 root"
	}
	if d > 0 {
		return "has 2 roots"
	}
	return "has no roots"
}
