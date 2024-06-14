package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Stat("main.go")

	fl := fmt.Println

	fl("File Name: ", file.Name())
	fl("Size in Bytes: ", file.Size())

	fl("Last modified: ", file.ModTime())
	fl("Is a directory: ", file.IsDir())

	ff := fmt.Printf

	ff("Permissions 9 bit: %s \n", file.Mode())
	ff("Permissions 3-digit: %o \n", file.Mode())
	ff("Permissions 4-digit: %04o \n", file.Mode())
}
