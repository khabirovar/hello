package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with message.
    if name == "" {
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in function.
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns on of a set of greeting messages.
// The return message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
	"Great to see you, %v!",
	"Hello, %v! Well met!",
    }

    // Return a randomly selected message format
    return formats[rand.Intn(len(formats))]
}
