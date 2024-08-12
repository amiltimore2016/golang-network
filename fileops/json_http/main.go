package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "https://jsonplaceholder.typicode.com/todos/1/"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal("Error", err)
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {

		todoItem := Todo{}

		json.NewDecoder(response.Body).Decode(&todoItem)

		fmt.Printf("Todo item: %v\n", todoItem.UserID)

	} else {
		fmt.Println("Error: ", response.StatusCode)
	}
}
