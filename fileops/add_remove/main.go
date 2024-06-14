package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var err error
	dpath := "./dir"
	fname := "file.txt"
	err = os.MkdirAll(dpath, 0755)
	if err != nil {
		panic(err)
	}
	fpath := filepath.Join(dpath, fname)
	fmt.Println(fpath)
	// create file
	file, _ := os.Create(fpath)
	file.Close()

	_ = os.Remove(fpath)
	_ = os.RemoveAll(dpath)
}
