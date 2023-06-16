package helper

import (
	"errors"
	"fmt"
)

// Here the type is available outside the class but the fields are not, in order to access those fields we can make use of public methods
type Person struct {
	name string
	age  int
}

// This is the idiomatic way of writing Setter function using a Set in front of the field name
func (p *Person) SetName(name string) {
	p.name = name
}

// In cases where we want some constraints to check before adding value directly to the field using . operator this can be used
// if this function was not provided even if the user type in age 500. there is no one to check the authenticity for that.
// But using SetAge functionality we can give conditions for error handling
func (p *Person) SetAge(age int) error {

	if age > 100 {
		return errors.New("age should be below 100")
	}

	p.age = age
}

// This is the idiomatic way of writing Getter function. i.e capitalized form of the field name
// We can distinguish between the field and function name as function is in capitalized form - and setter and getter are only required when fields are private/unexported
// In this case non capitalized
func (p *Person) Name() string {
	p.sayHi(p.name)
	return p.name
}

func (p *Person) Age() int {
	return p.age
}

// This right here is a private method. this is not exported, so the instance of the type which is implemented in other package cannot directly access this function using . operator
func (p *Person) sayHi(name string) {
	fmt.Println("hi abhinand")
}
