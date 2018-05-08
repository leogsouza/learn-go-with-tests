package main

import "fmt"


const (
    spanish = "Spanish"
    french = "French"
    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)


// Hello returns a greeting
func Hello(name, language string) string {
    
	if name == "" {
		name = "World"
	}
	
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
        default:
            prefix = englishHelloPrefix
	    case french:
	        prefix = frenchHelloPrefix
	    case spanish:
	        prefix = spanishHelloPrefix
	}
	
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
