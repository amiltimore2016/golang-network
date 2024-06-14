package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	scan := func(
		path string, i fs.DirEntry, _ error) error {
		fmt.Println(i.IsDir(), path)
		return nil
	}
	_ = filepath.WalkDir(".", scan)
}
