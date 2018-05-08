package main

import "fmt"

const helloPrefix = "Hello, "

// Hello returns a greeting
func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	
	if language == "Spanish" {
	    return "Hola, " + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
