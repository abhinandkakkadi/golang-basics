package main

import "fmt"


func main() {

	const a,b = 22,"abhinand"
	fmt.Printf("Name is %s and \t age is %d", b,a)
}

// here %s and %d and are called formatting verb and above we can see that two different type of values are assigned 
// and const have a weird way of type inference - check that out
// to know more about formatting verbs check go specification - official


// packages user
// fmt