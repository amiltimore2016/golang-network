package main

import (
	"fmt"
	"os"
)

func main() {
	files, _ := os.ReadDir(".")

	for _, file := range files {
		fileInfo, _ := file.Info()
		fmt.Println(file, fileInfo.ModTime())
	}
}
