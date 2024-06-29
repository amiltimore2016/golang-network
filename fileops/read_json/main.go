package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Inventory struct {
	Total   int      `json:"total"`
	Devices []string `json:"devices"`
}

func main() {
	data, _ := os.ReadFile("./inventory.json")
	inv := Inventory{}
	_ = json.Unmarshal([]byte(data), &inv)
	fmt.Printf("Total devices: %d\n Devices: %q\n", inv.Total, inv.Devices)
}
