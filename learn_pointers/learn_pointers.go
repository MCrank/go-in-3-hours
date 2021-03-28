package main

func main() {
	a := 10
	b := &a
	c := a
	println(a, b, *b, c)

	a = 20
	println(a, b, *b, c)

	*b = 30
	println(a, b, *b, c)

	c = 40
	println(a, b, *b, c)
}
