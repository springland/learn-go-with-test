package main

import "fmt"

const(
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix = "Bonjour "
	englishWorld = "World"
	ENGLISH = "english"
	FRENCH = "french"
	SPANISH = "spanish"
)
func HelloWorld() string {
	return "Hello World"
}

func Hello(name  string , language string) string{
	if (name == ""){
		name = englishWorld
	}

	var prefix string
	switch language {
	case ENGLISH:
		prefix = englishHelloPrefix
		break
	case FRENCH:
		prefix = frenchHelloPrefix
		break
	case SPANISH:
		prefix = spanishHelloPrefix
		break
	default: 
		prefix = englishHelloPrefix
		break			
	}
	return prefix + name
}
func main() {
	fmt.Print(HelloWorld())
}