package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"host"`
	Ports    []int  `yaml:"ports"`
	DnsEntry string `yaml:"dns_entry"`
}

func main() {
	data, _ := os.ReadFile("file.yml")
	conf := Config{}

	_ = yaml.Unmarshal(data, &conf)

	fmt.Printf("%s:%d:%s\n", conf.Host, conf.Ports, conf.DnsEntry)

}
