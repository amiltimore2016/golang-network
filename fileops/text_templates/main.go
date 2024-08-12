package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"text/template"
)

type Vlans struct {
	Id, Name string
}

func main() {
	data, _ := os.Open("vlans.csv")

	reader := csv.NewReader(data)

	vlans := []Vlans{}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Printf("Row is %v", row)
		vlans = append(vlans, Vlans{
			Id:   row[0],
			Name: row[1],
		})
	}

	tmpl, _ := template.ParseFiles("file.tmpl")
	_ = tmpl.Execute(os.Stdout, vlans)
	file, _ := os.Create("vlans.txt")
	defer file.Close()
	_ = tmpl.Execute(file, vlans)
}
