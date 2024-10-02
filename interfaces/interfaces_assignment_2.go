package main

import (
	"fmt"
	"io"
	"os"
)

func readFileContent(name string) {

	bs, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Errorf("Error while reading the file content")
	}

	data, err := io.Copy(os.Stdout, bs)

	fmt.Println(data)
}
