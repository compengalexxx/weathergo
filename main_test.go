package main

import "testing"

func TestGetCity(t *testing.T) {
	t.Run("returns the city name when one argument is provided", func(t *testing.T) {
		// 1. SETUP: Define the input for our function.
		// Command-line arguments are passed as a "slice of strings" ([]string).
		// The first element is always the program name, and the rest are the actual arguments.
		// This simulates running `go run main.go "São Paulo"` in the terminal.
		args := []string{"weather", "São Paulo"}

		// 2. CALL: Execute the function we are testing
		// We pass our simulated arguments to getCity. It will return two values:
		// the city string and a potential error.
		city, err := getCity(args)

		// Check that the returned city is "São Paulo" and the error is nil.

		// 3. ASSERTION 1: Check if an unexpected error occurred.
		// In this successful case, we expect the error to be `nil`.
		// If `err` is NOT nil, something went wrong. We fail the test immediately
		// using t.Fatalf because there's no point continuing if we got an error.
		if err != nil {
			t.Fatalf("getCity() returned an unexpected error: %v", err)
		}

		// 4. ASSERTION 2: Check if the returned city is correct.
		// We define what we expect the result to be.
		expected := "São Paulo"
		// If the city we got doesn't match what we expected, we fail the test
		// using t.Errorf. It marks the test as failed but would allow other checks
		// in the same test to run (if there were any).
		if city != expected {
			t.Errorf("getCity() returned %q, expected %q", city, expected)
		}
	})

	t.Run("returns an error when no arguments are provided", func(t *testing.T) {
		// 1. SETUP: Define the input for our function.
		args := []string{"weathergo"}

		// 2. CALL: Execute the function we are testing
		city, err := getCity(args)

		// Check that the returned city is an empty string "" and the error is NOT nil.

		// 3. ASSERTION 1: Check if an unexpected error occurred.
		if err == nil {
			t.Fatalf("getCity() incorrectly did not return an error")
		}

		// 4. ASSERTION 2: Check if the returned city is correct (an empty string)
		if city != "" {
			t.Errorf("getCity() returned %q, expected empty string", city)
		}

	})

	t.Run("returns an error when more than one argument is provided", func(t *testing.T) {
		// 1. SETUP: Define the input for our function.
		args := []string{"weathergo", "New York", "Extra"}

		// 2. CALL: Execute the function we are testing
		city, err := getCity(args)

		// Check that the returned city is an empty string "" and the error is NOT nil.
		// 3. ASSERTION 1: Check if an unexpected error occurred.
		if err == nil {
			t.Fatalf("getCity() incorrectly did not return an error")
		}

		// 4. ASSERTION 2: Check if the returned city is correct (an empty string)
		if city != "" {
			t.Errorf("getCity() returned %q, expected empty string", city)
		}
	})
}
