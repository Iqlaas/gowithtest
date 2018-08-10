package main

import "fmt"

const englishPrefix = "Hello "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola "
const french = "French"
const frenchHelloPrefix = "Bonjour "

// Hello function
func Hello(name, language string) string {
	if name == "" {
		name = "World!"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Go!", ""))
}
