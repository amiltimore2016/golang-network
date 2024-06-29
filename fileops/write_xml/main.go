package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Campus struct {
	XMLName  xml.Name `xml:"campus"`
	Name     string   `xml:"name,attr"`
	Comment  string   `xml:"comment"`
	Building Building
}

type Building struct {
	XMLName xml.Name `xml:"building"`
	Name    string   `xml:"name,attr"`
	Comment string   `xml:"comment"`
	Device  Device
}

type Device struct {
	XMLName xml.Name `xml:"device"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

func main() {
	data, _ := os.ReadFile("campus.xml")
	camp := &Campus{}

	_ = xml.Unmarshal([]byte(data), &camp)

	fl := fmt.Println

	fl("Campus Name: ", camp.Name)
}
