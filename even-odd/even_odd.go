package main

import "fmt"

type numbers []int

func (n numbers) printEvenOrOdd() {
	for _, v := range n {
		if v%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}
