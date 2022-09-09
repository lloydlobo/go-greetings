package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error)  {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // Returns a greeting that embeds the name in a message.
    // message := fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf(randomFormat(), name)

    // nil means no error as a second value in the successful return
    return message, nil
}

// NOTE: lowerCase function names can't be exported. private func
// Function init sets initial value for variables used here
func init()  {
    rand.Seed(time.Now().UnixNano())
}

// Function randomFormat returns one of a set of greeting messages.
// The returned message is selected at random
func randomFormat() string  {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by
    // specifying a random index for the slice of formats
    return formats[rand.Intn(len(formats))]
}
