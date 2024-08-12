package main

import (
	"encoding/xml"
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
	camp := &Campus{
		Name:    "campus1",
		Comment: "Building Comment",
		Building: Building{
			Name:    "building2",
			Comment: "Device Comment",
			Device: Device{
				Name: "RTR-1",
				Type: "router",
			},
		},
	}

	indt, _ := xml.MarshalIndent(camp, "", "  ")
	_ = os.WriteFile("file.xml", indt, 0644)
}
