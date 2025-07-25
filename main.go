package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// TemperatureData represents the temperature measurement data
type TemperatureData struct {
	Location    string
	Celsius     float64
	Reamur      float64
	Fahrenheit  float64
	Category    string
}

// NewTemperatureData creates a new TemperatureData instance
func NewTemperatureData(location string, celsius float64) *TemperatureData {
	return &TemperatureData{
		Location: location,
		Celsius:  celsius,
	}
}

// convertTemperatures converts Celsius to Reamur and Fahrenheit using pointer receiver
func (td *TemperatureData) convertTemperatures() {
	// Convert to Reamur: R = C * 4/5
	td.Reamur = td.Celsius * 4.0 / 5.0
	
	// Convert to Fahrenheit: F = (C * 9/5) + 32
	td.Fahrenheit = (td.Celsius * 9.0 / 5.0) + 32.0
}

// classifyTemperature classifies temperature based on Celsius value using pointer receiver
func (td *TemperatureData) classifyTemperature() {
	switch {
	case td.Celsius < 18:
		td.Category = "dingin"
	case td.Celsius >= 18 && td.Celsius <= 25:
		td.Category = "hangat"
	default:
		td.Category = "panas"
	}
}

// displayResults displays the temperature conversion results using pointer receiver
func (td *TemperatureData) displayResults() {
	fmt.Printf("> Suhu di %s adalah %s\n", td.Location, td.Category)
	fmt.Printf("> Suhu di %s dalam Reamur = %.0f\n", td.Location, td.Reamur)
	fmt.Printf("> Suhu di %s dalam Fahrenheit = %.0f\n", td.Location, td.Fahrenheit)
}

// validateLocation validates that location input contains only letters and spaces
func validateLocation(location string) bool {
	// Check if location is empty
	if strings.TrimSpace(location) == "" {
		return false
	}
	
	// Check if location contains only letters, spaces, and common location characters
	matched, _ := regexp.MatchString(`^[a-zA-Z\s\-'\.]+$`, location)
	return matched
}

// getLocationInput prompts user for location input with validation
func getLocationInput(scanner *bufio.Scanner) (string, error) {
	fmt.Print("Masukkan lokasi pengukuran suhu: ")
	
	if scanner.Scan() {
		location := strings.TrimSpace(scanner.Text())
		
		if !validateLocation(location) {
			return "", fmt.Errorf("input lokasi tidak valid, hanya menerima huruf dan spasi")
		}
		
		return location, nil
	}
	
	return "", fmt.Errorf("gagal membaca input lokasi")
}

// getTemperatureInput prompts user for temperature input with validation
func getTemperatureInput(scanner *bufio.Scanner) (float64, error) {
	fmt.Print("Masukkan suhu dalam Celsius: ")
	
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		
		// Try to parse the input as a number
		celsius, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return 0, fmt.Errorf("input tidak valid, hanya menerima angka")
		}
		
		return celsius, nil
	}
	
	return 0, fmt.Errorf("gagal membaca input suhu")
}

// processTemperatureConversion handles the complete temperature conversion process
func processTemperatureConversion(td *TemperatureData) {
	// Convert temperatures using pointer
	td.convertTemperatures()
	
	// Classify temperature using pointer
	td.classifyTemperature()
	
	// Display results using pointer
	td.displayResults()
}

func main() {
	fmt.Println("--- Konverter Suhu ---")

	// Create a scanner to read input
	scanner := bufio.NewScanner(os.Stdin)

	// Get location input with validation
	location, err := getLocationInput(scanner)
	if err != nil {
		fmt.Printf("> %s\n", err.Error())
		return
	}

	// Get temperature input with validation
	celsius, err := getTemperatureInput(scanner)
	if err != nil {
		fmt.Printf("> %s\n", err.Error())
		return
	}

	// Create TemperatureData struct using constructor
	tempData := NewTemperatureData(location, celsius)

	// Process temperature conversion using pointer
	processTemperatureConversion(tempData)
}
