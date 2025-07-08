package main

import (
	"bufio" // Needed for reading input line by line
	"fmt"
	"os"      // Needed for standard input (os.Stdin)
	"strings" // Helpful for trimming whitespace from user input
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create a new reader for standard input

	fmt.Print("Enter domain name (e.g., boot.dev): ") // Prompt the user
	input, err := reader.ReadString('\n')             // Read until a newline character
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// Clean up the input: remove leading/trailing whitespace and the newline character
	domainToLookup := strings.TrimSpace(input)

	if domainToLookup == "" {
		fmt.Println("No domain entered. Exiting.")
		return
	}

	ip, err := getIPAddress(domainToLookup) // Call your function with the user's input
	if err != nil {
		fmt.Printf("Error getting IP address for %s: %v\n", domainToLookup, err)
		return
	}
	fmt.Printf("The IP address for %s is: %s\n", domainToLookup, ip)
}
