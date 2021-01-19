package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const prefixenglish = "Hello, "
const prefixespanhol = "Hola, "
const prefixfrances = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixsalutation(language) + name
}

func prefixsalutation(language string) (prefix string) {
	switch language {
	case frances:
		prefix = prefixfrances
	case espanhol:
		prefix = prefixespanhol
	default:
		prefix = prefixenglish
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
