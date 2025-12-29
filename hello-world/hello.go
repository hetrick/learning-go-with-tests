package main

import "fmt"

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"

	helloEnglish = "Hello"
	helloSpanish = "Hola"
	helloFrench  = "Bonjour"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = helloEnglish
	case spanish:
		prefix = helloSpanish
	case french:
		prefix = helloFrench
	default:
		prefix = helloEnglish
	}

	return
}
