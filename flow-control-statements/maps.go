package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var menu map[string]float64

func main() {
	m = make(map[string]Vertex)
	
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m)

	menu = map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
	}

	menu["pizza"] = 12.99

	fmt.Println(menu, menu["pie"])
	delete(menu, "pie")

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	soup := menu["soup"]
	soup, soup_ok := menu["soup"]
	pizza, ok := menu["pizza"]
	fmt.Println("Soup's price:", soup, "Present?:", soup_ok)
	fmt.Println("Pizza's price:", pizza, "Present?:", ok)

}