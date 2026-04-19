package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := "Hellow this is text"
	err := os.WriteFile("test.txt", []byte(data), 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done  ")
}
