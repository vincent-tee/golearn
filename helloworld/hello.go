package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
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
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("world", "Spanish"))
}
