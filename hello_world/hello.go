package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const russian = "Russian"
const englishPrefixHello = "Hello, "
const spanishPrefixHello = "Hola, "
const frenchPrefixHello = "Bonjour, "
const russianPrefixHello = "Privyet, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefixHello
	case spanish:
		prefix = spanishPrefixHello
	case russian:
		prefix = russianPrefixHello
	default:
		prefix = englishPrefixHello
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
