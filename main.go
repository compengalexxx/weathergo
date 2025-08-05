package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Weather struct {
	// Weather holds the overall structure of the wttr.in JSON response.
	CurrentCondition []CurrentCondition `json:"current_condition"`
}

type CurrentCondition struct {
	// CurrentCondition holds the detailed weather information.
	TempC       string `json:"temp_C"`
	FeelsLikeC  string `json:"FeelsLikeC"`
	WeatherDesc []struct {
		Value string `json:"value"`
	} `json:"weatherDesc"`
}

func main() {
	// 1. Get the city from command-line arguments.
	city, err := getCity(os.Args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// 2. Construct the API URL.
	// Use `url.QueryEscape` to handle spaces or special characters in city names.
	apiURL := fmt.Sprintf("https://wttr.in/%s?format=j1", url.QueryEscape(city))

	// 3. Fetch the weather data.
	weather, err := getWeather(apiURL)
	if err != nil {
		fmt.Printf("Error fetching weather: %s\n", err)
		os.Exit(1)
	}

	// 4. Print the results.
	// We add a check to make sure we actually got data back.
	if len(weather.CurrentCondition) > 0 {
		current := weather.CurrentCondition[0]
		fmt.Printf("Weather for %s:\n", city)
		fmt.Printf("Temperature: %s°C\n", current.TempC)
		fmt.Printf("Feels Like: %s°C\n", current.FeelsLikeC)
		if len(current.WeatherDesc) > 0 {
			fmt.Printf("Description: %s\n", current.WeatherDesc[0].Value)
		}
	} else {
		fmt.Println("Could not find weather information.")
	}
}

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
func getWeather(apiURL string) (Weather, error) {
	// Step 1: Make the HTTP GET request.
	// http.Get returns a response and an error. We handle the error immediately.
	resp, err := http.Get(apiURL)
	if err != nil {
		return Weather{}, fmt.Errorf("could not get weather data: %w", err)
	}
	// `defer` is a special Go keyword. It schedules the `resp.Body.Close()` call
	// to be run just before the function returns. This is the standard way to
	// ensure resources like network connections are always closed.
	defer resp.Body.Close()

	// Step 2: Check for a successful status code.
	// A successful response is usually 200 OK. If not, the city might not exist.
	if resp.StatusCode != http.StatusOK {
		return Weather{}, fmt.Errorf("weather API returned a non-success status: %s", resp.Status)
	}

	// Step 3: Read the response body.
	// `io.ReadAll` reads everything from the response body into a byte slice.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Weather{}, fmt.Errorf("could not read response body: %w", err)
	}

	// Step 4: Parse (unmarshal) the JSON data.
	var weather Weather
	// `json.Unmarshal` takes the JSON data (as a byte slice) and a pointer to
	// the struct where the data should be stored. It parses the JSON and fills
	// in the struct fields based on the `json:"..."` tags we defined.
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return Weather{}, fmt.Errorf("could not parse weather JSON: %w", err)
	}

	// Step 5: Return the populated struct and a nil error.
	return weather, nil
}
