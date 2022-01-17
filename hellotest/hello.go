package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Fred"))
}
