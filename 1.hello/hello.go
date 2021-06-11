package main

import (
	"fmt"
	"strings"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHollaPrefix = "Hola, "
const frenchBonjourPrefix = "Bonjour, "

func Hello(name, language string) string {
	name = strings.Title(strings.ToLower(name))
	var prefix string
	var defaultName string

	switch language {
	case spanish:
		prefix = spanishHollaPrefix
		defaultName = "Mundo"
	case french:
		prefix = frenchBonjourPrefix
		defaultName = "Monde"
	default:
		prefix = englishHelloPrefix
		defaultName = "World"
	}

	if name == "" {
		name = defaultName
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
