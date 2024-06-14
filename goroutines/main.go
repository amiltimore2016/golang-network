package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from a gorouting")
}

func main() {
	fmt.Println("Hello from main")
	// go sayHello()
	go sayHello()
	fmt.Println("I'm running concurrently")
	time.Sleep(2 * time.Second)

	fmt.Println("Back in the main thred")
}
