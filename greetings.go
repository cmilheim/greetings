package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hell(name string) string {
	// Return a greeting that embeds the name in the message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
