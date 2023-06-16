package main

import "fmt"

type Teacher interface {
	marks()
}

type Student interface {
	Teacher
	fun()
}

// if you want to be a teacher type. just implement the marks function
// but if you want to be Student type. You have to implement marks() and fun().
type teacher struct {
	name string
}

// teacher just needs marks
func (t1 *teacher) marks() {
	fmt.Println("teacher needs marks")
}

// this is a student who should satisfy marks and fun
type student struct {
	name string
}

func (s1 *student) marks() {
	fmt.Println("this is for student")
}

func (s1 *student) fun() {
	fmt.Println("this is defently student")
}

// here student have implemented Teacher and Student interface
func main() {

	s1 := &student{}

	// here the variable of Student interface is able to assign to student struct because student struct have implemented
	// the interface Student and now they can be interchangebly used
	// if the struct did not implemented those function . the assign that we do below will not be possible
	var s Student
	s = s1
	s.marks()

}
