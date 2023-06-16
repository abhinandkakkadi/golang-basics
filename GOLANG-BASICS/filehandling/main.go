package main

import (
	"log"
	"os"
)

// creating a file using os package
func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	// write some string to file
	_, err = file.WriteString("Hey my name is abhinand")
	if err != nil {
		log.Fatal(err.Error())
	}
}
