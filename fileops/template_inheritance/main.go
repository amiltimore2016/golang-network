package main

import (
	"fmt"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Base Base `base`
	Bgp  Bgp  `bgp`
	Ntp  Ntp  `ntp`
}

type Base struct {
	Hostname string `hostname`
}

type Neighbors struct {
	Ip string `ip`
	As string `as`
}

type Bgp struct {
	As        string      `as`
	Id        string      `id`
	Neighbors []Neighbors `neighbors`
}

type Ntp struct {
	Source    string `source`
	Primary   string `primary`
	Secondary string `secondary`
}

func main() {
	config := Config{}

	data, _ := os.ReadFile("file.yml")
	_ = yaml.Unmarshal(data, &config)
	tmpl := make(map[string]*template.Template)
	files, _ := template.ParseFiles("base.tmpl", "bgp.tmpl", "ntp.tmpl")
	tmpl["base.tmpl"] = template.Must(files, nil)
	error := tmpl["base.tmpl"].ExecuteTemplate(os.Stdout, "base", config)
	fmt.Println(error)
}
