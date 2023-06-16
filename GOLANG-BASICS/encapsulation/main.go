package main

import (
	"encap/helper"
	"fmt"
	"log"
)


func main() {
	// Here we can't directly access the fields inside the Person struct as it is unexported, (encapsulatoin)
	//  

	var p helper.Person
	p.SetName("abhinand")
	err := p.SetAge(23)
	if err != nil {
		// log fatal works similar to println with the difference begin log.Fatal call the os.Exit() function so that the program will be terminated. 
		// But in case of println function the program will still be running even if error ocurred
		log.Fatal(err)
	}

	fmt.Println(p.Name())
	fmt.Println(p.Age())

}