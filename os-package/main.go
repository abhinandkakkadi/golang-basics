package main

import (
	"fmt"
	"os"
)



func main() {

	file,err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := make([]byte,100)
	// Read will take a byte slice and sent back the number of bytes written
	writeCount,err := file.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data[:writeCount]))
}