package handler

import "fmt"



func DeferExample() {

	defer func(){
		if e := recover(); e != nil {
				fmt.Println("the panicked message: ",e)
		}
	}()
	panic("this is panicked message")
	fmt.Println("since the program panics. It will never reach here")
}