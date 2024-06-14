// Golang program to check the size and
// modified time of a file
package main

// Importing the required packages
import (
	"fmt"
	"log"
	"os"
)

func main() {

	// Get the fileinfo
	fileInfo, err := os.Stat("main.go")

	// Checks for the error
	if err != nil {
		log.Fatal(err)
	}

	// Gives the modification time
	modificationTime := fileInfo.ModTime()
	fmt.Println("Name of the file:", fileInfo.Name(),
		" Last modified time of the file:",
		modificationTime)

	// Gives the size of the file in bytes
	fileSize := fileInfo.Size()
	fmt.Println("Size of the file:", fileSize)
}
