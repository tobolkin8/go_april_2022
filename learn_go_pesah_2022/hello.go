package main

import "fmt"

var Languages = map[string]bool{"english": true, "spanish": true, "french": true}

func Hello(name string, language string) string {
	prefixString := ReturnLangText(language)
	fullReturnText := prefixString + name
	if name == "" {
		fullReturnText = prefixString + "world!"
	}
	return fullReturnText
}

func ReturnLangText(language string) string {
	if !Languages[language] {
		language = "default"
	}
	returnString := ""
	switch language {
	case "english":
		returnString = "Hello, "
	case "spanish":
		returnString = "Holla, "
	case "french":
		returnString = "Bonjour, "
	default:
		returnString = "Hello, "
	}
	return returnString
}

func main() {
	fmt.Println(Hello("", ""))
}
