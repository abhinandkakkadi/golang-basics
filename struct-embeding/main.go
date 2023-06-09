package main

import "fmt"


type Pascal struct {
	s string
}

func (p *Pascal) function1() {
	fmt.Println("this is actually a function of Pascal by try composition")
}

func (p *Pascal) function2() {
	fmt.Println("this is also a function of Pascal try me too")
}

// this function have overridden the other function2().
// actually composition is like inheritence only
// all the methods and fields above above embedded struct is made available for embedding struct
func (g *Go) function2() {
	fmt.Println("this is a function of same name but we  we look which will execute")
}

// here we are embedding Pascal into Go
type Go struct {
	Pascal
}

func main() {

	// here we are creating a new empty instance address of Go and it is calling functions
	// which have Pascal as receiver. This is possible because Go embeds Pascal and due the property called composition 
	// we are able to do the above thing

	g := &Go{}
	g.function1()
	g.function2()

	// if we look at
}