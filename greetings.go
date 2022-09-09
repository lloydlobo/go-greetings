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
    // message := fmt.Sprint(randomFormat()) // Failing test

    // nil means no error as a second value in the successful return
    return message, nil
}

// Function Hellos returns a map that associates each of the named
// people with a greeting message
func Hellos(names []string) (map[string]string, error)  {
    // A map to associate names with messages.
    // Syntax of make =>  messages := make(type, 0)
    messages := make(map[string]string)

    // Loop throught the received slice of names, calling
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with the name.
        messages[name] = message
    }

    return messages, nil
    
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
        "Great to see you, %v.",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by
    // specifying a random index for the slice of formats
    return formats[rand.Intn(len(formats))]
}
