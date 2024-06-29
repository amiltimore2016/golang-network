package main

import (
	"bufio"
	"os"
)

func main() {
	lines := []string{"RTR1 1.1.1.1",
		"RTR2 2.2.2.2",
		"RTR3 3.3.3.3",
	}
	file, _ := os.OpenFile("./file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}
	writer.WriteString("\n")
	writer.Flush()
	file.Close()

}
