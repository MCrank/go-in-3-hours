package main

import "fmt"

// Type conversion
func main() {
	var i int8 = 20
	var f float32 = 5.6
	var j int32 = 40
	fmt.Println(float32(i) + f)
	fmt.Println(i + int8(f))
	fmt.Println(i + int8(f+1.9))
	fmt.Println(i + int8(j))
}
