package main

import "os"

func main() {
	file, _ := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file.Write([]byte("Append Data \n"))
	file.Close()
}
