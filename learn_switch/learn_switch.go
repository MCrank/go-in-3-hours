package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	if word == "hello" {
		fmt.Println("Hi yourself")
	} else if word == "goodbye" {
		fmt.Println("So long!")
	} else if word == "greetings!" {
		fmt.Println("Salutaions!")
	} else {
		fmt.Println("I don't know what you said")
	}

	// Only code that matches the case in run.  No fall through by default
	switch word {
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye":
		fmt.Println("So long!")
	case "greetings!":
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said")
	}

	// If you want fall through do it like this
	switch word {
	case "hi":
		fmt.Println("Yo Yo girlfriend")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye":
		fmt.Println("So long!")
	case "greetings!":
		fmt.Println("Salutatins!")
	default:
		fmt.Println("I don't know what you said")
	}

	// Can also have multiple checks per case and an empty case as well.  Treated as NoOp
	switch word {
	case "hi":
		fmt.Println("Yo Yo girlfriend")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	// Either of these will match
	case "goodbye", "bye":
		fmt.Println("So long!")
	// This will match but since no body then no code is run
	case "farewell":
	case "greetings!":
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said")
	}

	// Can also initialize variables in the switch.  Here we are getting the length of the word
	// This can be used as the case values as well if you so desire
	switch length := len(word); word {
	case "hi":
		fmt.Println("Yo Yo girlfriend")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	// Either of these will match
	case "goodbye", "bye":
		fmt.Println("So long!")
	// This will match but since no body then no code is run
	case "farewell":
	case "greetings!":
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said but it was", length, "characters long")
	}

	c := "crackerjack"
	switch length := len(word); {
	case word == "hi":
		fmt.Println("Yo Yo girlfriend")
		fallthrough
	case word == "hello":
		fmt.Println("Hi yourself")
	// Either of these will match
	case length == 1:
		fmt.Println("I don't know any 1 letter words")
	// This will match but since no body then no code is run
	case 1 < length && length < 10, word == c:
		fmt.Println("This word is either", c, "or it is 2-9 characters long")
	default:
		fmt.Println("I don't know what you said but it was", length, "characters long")
	}
}
