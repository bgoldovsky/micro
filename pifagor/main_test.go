package main

import "testing"

type PifagorianData struct {
	A        int64
	B        int64
	C        int64
	Expected string
}

var data = []PifagorianData{
	{3, 4, 5, "yes"},
	{4, 3, 5, "yes"},
	{8, 10, 6, "yes"},
	{1, 2, 3, "no"},
}

func TestPrifagor(t *testing.T) {
	for _, val := range data {
		if act := isPifagor(val.A, val.B, val.C); act != val.Expected {
			t.Errorf("expected: %s, act %s. vars: %d, %d, %d\n", val.Expected, act, val.A, val.B, val.C)
		}
	}
}
