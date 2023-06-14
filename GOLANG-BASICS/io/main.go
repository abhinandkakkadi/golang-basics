package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


func main() {

	r := strings.NewReader("hey my name is abhinand")
	dst := os.Stdout
	 _,err := io.Copy(dst,r)
	 if err != nil {
		fmt.Println(err)
	 }

	_,err = io.WriteString(dst,"Hey my name is again abhinand")
	if err != nil {
		fmt.Println(err)
	}
}