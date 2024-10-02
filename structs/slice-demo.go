package main

import "fmt"

func (d data) updateSlice() {
	d[0] = "NewData"
}

func (d data) showSliceData() {
	for _, v := range d {
		fmt.Println(v)
	}
}
