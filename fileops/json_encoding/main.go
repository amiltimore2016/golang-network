package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func main() {
	var P = Person{
		Name:   "John",
		Age:    30,
		Active: true,
	}

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	err = encoder.Encode(P)
	if err != nil {
		fmt.Println("ERROR encoding", err)
	}
}
