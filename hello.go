/*
When you write a program in Go, you will have a main package defined with a main func inside it.
Packages are ways of grouping up related Go code together.
*/

package main

import "fmt" // We import the fmt package so we have access to the Printkn function

const englishHelloPrefix = "Hello, "

// Domain -> Main Code

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

// Outside world code -> Code that actually gets executed and seen to the consumer
func main() { //func is the keyword for defining functions
	fmt.Println(Hello("world")) // How package functions are called
}
