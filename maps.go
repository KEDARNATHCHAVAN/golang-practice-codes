package main

import "fmt"

func main() {

	//maps
	city := make(map[string]string)
	city["S"] = "sangli"
	city["P"] = "palus"
	city["K"] = "kolhapur"
	fmt.Println(city)

	//to search for an value using the key
	c := city["S"]
	fmt.Println(c)

	//to delete an element
	delete(city, "K")
	fmt.Println(city)

	//to append
	city["PU"] = "pune"
	fmt.Println(city)

	//to iterate through the map
	for k, v := range city {
		fmt.Printf("%v: %v\n", k, v)
	}
}
