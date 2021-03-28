package main

func main() {
	a := addNumbers(2, 3)
	println(a)

	a = addNumbers(5, 5)
	println(a)

	a = addNumbers(3, 4)
	println(a)

	b, remainder := divAndRemainder(5, 7)
	println(b, remainder)

}

func addNumbers(a int, b int) int {
	return a + b
}

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}
