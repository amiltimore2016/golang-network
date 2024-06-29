package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	Ports    []int  `yaml:"ports"`
	DnsEntry string `yaml:"dns_entry"`
}

func main() {
	conf := Config{
		Host:     "localhost",
		Ports:    []int{8080, 8081},
		DnsEntry: "0000000",
	}
	data, _ := yaml.Marshal(&conf)
	_ = os.WriteFile("config.yaml", data, 0644)
}
