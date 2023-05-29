package main

import (
	"fmt"

	"github.com/abhinandkakkadi/golangbasics/handler"
)

func main() {

	fmt.Println("going to handle a panic using defer")
	handler.DeferExample()
	fmt.Println("this code is printed since the error is handled")
}	