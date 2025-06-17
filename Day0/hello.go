package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

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

func say_hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(say_hello("kishore", spanish))
}
