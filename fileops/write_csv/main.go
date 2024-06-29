package main

import (
	"encoding/csv"
	"os"
)

func main() {
	rows := [][]string{
		{"HOST", "IP ADDR", "LOCATION"},
		{"RTR1", "1.1.1.1", "LONDONG"},
		{"RTR2", "2.2.2.2", "SYDNEY"},
		{"RTR3", "0000000", "NEWYORK"}}
	file, _ := os.Create("hosts.csv")
	writer := csv.NewWriter(file)

	for _, row := range rows {
		_ = writer.Write(row)
	}

	writer.Flush()
	file.Close()
}
