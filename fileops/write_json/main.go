package main

import (
	"encoding/json"
	"fmt"
)

type Inventory struct {
	Total   int      `json:"total"`
	Devices []string `json:"devices"`
}

func main() {
	data := Inventory{
		Total:   3,
		Devices: []string{"switch", "router", "server"},
	}

	inv, _ := json.MarshalIndent(data, " ", "")
	fmt.Println(string(inv))
}
