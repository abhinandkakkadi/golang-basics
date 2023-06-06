package structembed

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Sound() {
	fmt.Println(c.Name)
}

type Tiger struct {
	Cat
}

// so what actually happend is we embed a struct into another struct and a this another struct have access to all the fields in the former struct and also it can access all the function that is implemented by that particualr
// type using dot operator
func StructEmbeding() {
	c := Cat{"Malu"}
	c.Sound()

	t := Tiger{Cat: Cat{Name: "Mal"}}
	// tiger instance  can call sound since tiger hve embedded cat by the property of composition.
	t.Sound()
}
