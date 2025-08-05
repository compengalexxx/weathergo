package main

import (
	// The "fmt" package provides functions for formatted I/O (like printing to the console).
	"fmt"
	// The "os" package provides a platform-independent interface to operating system functionality.
	// We need it here to access the command-line arguments.
	"os"
)

// The main function is the entry point of the application. When you run your program,
// the code inside this function is what gets executed first.
func main() {
	// os.Args is a slice of strings that holds all the command-line arguments
	// passed to the program. We pass it to our getCity function to be processed.
	city, err := getCity(os.Args)
	if err != nil {
		// If getCity returns an error, we print it to the console and exit.
		// os.Exit(1) signals that the program terminated with an error.
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// If everything goes well, we print the name of the city.
	// Later, this is where we will call the weather API.
	fmt.Printf("Fetching weather for city: %s\n", city)
}

// getCity processes the command-line arguments to extract the city name.
//
// It takes a slice of strings (args) as input, which represents the command-line arguments.
// It returns two values:
// 1. A string containing the city name.
// 2. An error object. If no error occurs, this will be `nil`.
//
// This is a very common pattern in Go: functions often return a value and an error.
func getCity(args []string) (string, error) {
	numberOfArgs := len(args)
	if numberOfArgs < 2 {
		return "", fmt.Errorf("no arguments provided")
	} else if numberOfArgs == 2 {
		return args[1], nil
	} else { // numberOfArgs > 2
		return "", fmt.Errorf("too many arguments provided")
	}
}
