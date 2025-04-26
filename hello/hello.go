package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const italian = "Italian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(HelloWorld("Chris", ""))
}
