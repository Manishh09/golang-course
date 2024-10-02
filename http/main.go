package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/users"

	readAPIData(url)
}

func readAPIData(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Errorf("error while fetching response", err)
		os.Exit(1)
	}
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	// var users []User
	// err = json.Unmarshal(body, &users)
	// if err != nil {
	// 	fmt.Println("Error unmarshalling JSON:", err)
	// 	return
	// }
	// fmt.Println(users)

	// bs := make([]byte, 0, 99999) // To avoid re-allocation of capacity

	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	n, err := io.Copy(os.Stdout, resp.Body) // to Read data and Write the read data12

	if err != nil {
		return
	}

	fmt.Println(n)
}

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

func (s square) area() int {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}
