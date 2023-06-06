package main

import (
	"flag"
	"fmt"
)


func main() {

	word := flag.Int("val",2,"Enter a number")
	fmt.Println(*word)
}