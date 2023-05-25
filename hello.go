/*
When you write a program in Go, you will have a main package defined with a main func inside it.
Packages are ways of grouping up related Go code together.
*/

package main

import "fmt" // We import the fmt package so we have access to the Printkn function

const (
	//Language modes
	spanish = "Spanish"
	french  = "French"

	//Prefix Suite
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Domain -> Main Code

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return prefixSetter(language) + name

}
func prefixSetter(language string) (prefix string) { //the "(prefix string)" part is a named return value that can also be part of a function signature
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

// Outside world code -> Code that actually gets executed and seen to the consumer
func main() { //func is the keyword for defining functions
	fmt.Println(Hello("world", "")) // How package functions are called
}
