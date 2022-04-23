package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
