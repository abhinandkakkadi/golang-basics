package main

import "fmt"

func main() {

	x := 40

	// switch first way
	switch {
	case x < 20:
		fmt.Println("first condition is right")
	case x < 30:
		fmt.Println("second condition is right")
	case x < 40:
		fmt.Println("third condition is right")
	case x < 50:
		fmt.Println("forth condition is right")

	default:
		fmt.Println("nothing is correct")
	}

	// switch second way
	switch x {
	case 30:
		fmt.Println("yeah first is correct")
	case 40:
		fmt.Println("yeah the second is correct")
	default:
		fmt.Println("nothing is correct")
	}

	// fallthrough way -- usually if one condition is right then rest of them will not execute. But in this case
	// even though the first condition is right second condition will be checked

	switch {
	case x < 20:
		fmt.Println("first condition is right")
	case x < 30:
		fmt.Println("second condition is right")
	case x < 40:
		fmt.Println("third condition is right")
	case x < 50:
		fmt.Println("forth condition is right")
		fallthrough
	case x < 56:
		fmt.Println("fifth is also right")
	default:
		fmt.Println("nothing is correct")
	}

}
