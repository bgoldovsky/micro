package main

import (
	"fmt"
	"testing"
)

type LinearData struct {
	B        int64
	C        int64
	Expected string
}

type QuadraticData struct {
	A        int64
	B        int64
	C        int64
	Expected string
}

type SolveData struct {
	A        int64
	B        int64
	C        int64
	Expected string
}

var linearData = []LinearData{
	{2, 2, "has 1 root"},
	{2, 2, "has 1 root"},
	{0, 2, "has no roots"},
}

var quadraticData = []QuadraticData{
	{2, 1, 2, "Has No roots"},
	{1, -3, 2, "Has 2 roots"},
	{1, -200, 2, "Has 2 roots"},
}

var solveData = []SolveData{
	{0, 2, 2, "has 1 root"},
	{1, -3, 2, "has 2 roots"},
	{0, 0, 2, "has no roots"},
}

func TestLinear(t *testing.T) {
	for _, val := range linearData {
		if act := Linear(val.B, val.C); act != val.Expected {
			fmt.Println("Error")
		}
	}
}

func TestQuadratic(t *testing.T) {
	for _, val := range quadraticData {
		if act := Quadratic(val.A, val.B, val.C); act != val.Expected {
			fmt.Println("Error")
		}
	}
}

func TestSolve(t *testing.T) {
	for _, val := range solveData {
		if act := Solve(val.A, val.B, val.C); act != val.Expected {
			fmt.Println("Error")
		}
	}
}
