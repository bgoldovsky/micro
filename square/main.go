package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Output struct {
	A        int64  `json:"a"`
	B        int64  `json:b"`
	C        int64  `json:"c"`
	Equation string `json:"equation"`
	Result   string `json:"result"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{a}/{b}/{c}", getSolve).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getSolve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars((r))

	a, err := strconv.ParseInt(params["a"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseInt(params["b"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := strconv.ParseInt(params["c"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := Solve(a, b, c)

	data := Output{
		A:        a,
		B:        b,
		C:        c,
		Equation: "ax^2 + bx + c = 0",
		Result:   result,
	}

	json.NewEncoder(w).Encode(data)
}

func Solve(a, b, c int64) string {
	if a == 0 {
		return Linear(b, c)
	}
	return Quadratic(a, b, c)
}

func Linear(b, c int64) string {
	if b == 0 {
		return "has no roots"
	}

	return "has 1 root"
}

func Quadratic(a, b, c int64) string {
	d := b*b - 4*a*c
	if d == 0 {
		return "has 1 root"
	}
	if d > 0 {
		return "has 2 roots"
	}
	return "has no roots"
}
