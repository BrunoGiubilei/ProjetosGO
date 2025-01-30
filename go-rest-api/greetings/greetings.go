package greetings

import "fmt"

func GetGreeting(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to our API.", name)
}
