package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

type Output struct {
	A      int64  `json:"a"`
	B      int64  `json:"b"`
	C      int64  `json:"c"`
	Result string `json:"result"`
}

func main() {
	input := ImputData{A: 20, B: 20, C: 30}
	data, err := proto.Marshal(&input)
	if err != nil {
		log.Fatalln(input)
	}
	fmt.Println(data)

	output := ImputData{}
	err = proto.Unmarshal(data, &output)
	if err != nil {
		log.Fatalln(data)
	}
	fmt.Println(output)

	return

	r := mux.NewRouter()
	r.HandleFunc("/{a}/{b}/{c}", getPifagor).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func getPifagor(w http.ResponseWriter, r *http.Request) {
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

	result := isPifagor(a, b, c)

	data := Output{
		A:      a,
		B:      b,
		C:      c,
		Result: result,
	}

	json.NewEncoder(w).Encode(data)
}

func isPifagor(a, b, c int64) string {
	if a*a == b*b+c*c || b*b == a*a+c*c || c*c == a*a+b*b {
		return "yes"
	}
	return "no"
}
