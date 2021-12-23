package main

import "fmt"

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

	prefix := englishPrefixHello
	switch language {
	case french:
		prefix = frenchPrefixHello
	case spanish:
		prefix = spanishPrefixHello
	case russian:
		prefix = russianPrefixHello
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
