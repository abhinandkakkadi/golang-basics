package main

import (
	"fmt"

	"github.com/abhinandkakkadi/golangbasics/errorhandler"
	"github.com/abhinandkakkadi/golangbasics/structembed"
)

func main() {

	fmt.Println("going to handle a panic using defer")
	errorhandler.DeferExample()
	fmt.Println("this code is printed since the error is handled")

	structembed.StructEmbeding()

}
