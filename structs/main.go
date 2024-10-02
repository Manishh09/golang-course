package main

import "fmt"

type address struct {
	email    string
	location string
}

type employee struct {
	name        string
	designation string
	address     // address address
}

type data []string

func main() {
	// order of the values matters
	// not recommended
	emp1 := employee{
		"Emp1",
		"Angular Dev",
		address{},
	}
	fmt.Println(emp1)
	// using properties inside the struct
	// order is ignored
	emp2 := employee{
		name:        "Emp2",
		designation: "Go Dev",
		address: address{
			email:    "hello@go.dev",
			location: "India",
		},
	}
	fmt.Println(emp2)

	// using var keyword
	var emp3 employee
	// zero-values of the properties inside the struct
	fmt.Println(emp3)

	emp3.name = "Emp3"
	emp3.designation = ".Net Dev"
	emp3.address.email = "abc@go.dev"
	emp3.address.location = "IN"

	// print struct with all the property names inside it
	fmt.Printf("%+v", emp3)

	// will not update original data
	emp3.update("Emp4")

	// store address of the struct employee
	ePtr := &emp3
	fmt.Println("memory address  before updating is:", &ePtr)

	// updates original value
	ePtr.updateStructValue("Emp4")

	fmt.Println("memory address after updating is:", &ePtr)

	emp3.showDetails()

}

func (e employee) showDetails() {
	fmt.Println(e)
}

func (e employee) update(name string) {
	e.name = name
}

func (ePtr *employee) updateStructValue(name string) {

	(*ePtr).name = name
}
