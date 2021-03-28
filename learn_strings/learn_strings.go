package main

import "fmt"

func main() {
	s := "👋 🌎"
	fmt.Println(s)

	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6

	fmt.Println(vals, vals[0], vals[1], vals[2])

	// This throws an error because array of different sizes
	var vals2 [4]int = vals
	fmt.Println(vals, vals2)

}
