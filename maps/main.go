package main

import "fmt"

type keyval map[string]string

func main() {

	// first approach
	var mapWithVar keyval

	fmt.Println(mapWithVar)

	// using make
	mapWithMake := make(keyval)

	mapWithMake["catA"] = "Fiction"

	mapWithMake["catB"] = "Comedy"

	fmt.Println(mapWithMake)

	// direct initialization
	movies := keyval{
		"catA": "fiction",
		"catB": "comedy",
		"catC": "sci-fi",
		"catD": "anime",
	}

	fmt.Println(movies)

	// can delete key-value pair in a map
	delete(movies, "catA")
	// updates map : original values will be updated
	movies.updateMap()

	movies.showMapData()

}

func (kv keyval) showMapData() {

	fmt.Println(len(kv))

	for k, v := range kv {

		fmt.Println(k, v)

	}

}

func (kv keyval) updateMap() {
	kv["catB"] = "Horror"
}
