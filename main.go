package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Konverter Suhu ---")

	// Create a scanner to read input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan suhu dalam Celsius: ")

	// Read input
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		// Try to parse the input as a number
		celsius, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("> Input Tidak Valid, hanya menerima angka")
			return
		}

		// Convert to Reamur: R = C * 4/5
		reamur := celsius * 4.0 / 5.0

		// Convert to Fahrenheit: F = (C * 9/5) + 32
		fahrenheit := (celsius * 9.0 / 5.0) + 32.0

		// Display results
		fmt.Printf("> Suhu dalam Reamur = %.0f\n", reamur)
		fmt.Printf("> Suhu dalam Fahrenheit = %.0f\n", fahrenheit)
	}
}
