package greetings

import "fmt"

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	massage := fmt.Sprintf("Hi, %v. Welcome!", name)
	return massage
}
