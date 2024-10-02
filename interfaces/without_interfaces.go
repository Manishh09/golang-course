package main

import "fmt"

type deBot struct{}

type enBot struct{}

// we cannot have same method
// if other types are added , more methods are needed for their implementation

// Its a problem ?.. How to avoid ?

// Solution: Use interfaces

// func printMessage(de deBot) {
// 	fmt.Println(de.getMessage())
// }

func printMessage(en enBot) {
	fmt.Println(en.getMessage())
}

func (de deBot) getMessage() string {
	return "HI"
}

func (en enBot) getMessage() string {
	return "Hello"
}
