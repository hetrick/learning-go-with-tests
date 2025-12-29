package main

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return helloPrefix + name
}
