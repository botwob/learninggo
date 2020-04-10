package main

import "fmt"

func main() {
	//asignment and variables

	a := 5
	b := 3.141592653
	c := 7

	fmt.Println(float64(a) + b)
	fmt.Println(c % a)

	// simple function
	newline()

	//loops
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("Hello world: ", i)
	}
	newline()

	// while loop
	i := 0
	for i < 10 {
		fmt.Println("Hello world: ", i)
		i++
	}
	newline()

	//endless loop
	// for {
	// 	fmt.Println("Hello world!")
	// }

	//strings
	text := "Hello world!"
	fmt.Println(text)
	newline()

	//concatenating text
	text += " abcd"
	fmt.Println(text)

	newline()

	//conditionals
	if a > 5 {
		println("a is greater than five!")
	} else if a < 5 {
		println("a is less than five!")
	} else {
		println("a is equal to five!")
	}

	newline()

	//this also works, thank god
	x := true

	if x {
		println("x is true")
	} else {
		println("x is false")
	}

	newline()
	//also this

	x = false

	if !x {
		println("x is false")
	} else {
		println("x is true")
	}

}

func newline() {
	fmt.Println()
}
