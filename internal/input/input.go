package input

import (
	"bufio"
	"os"
	"strings"
)

// GetGuess prompts the player to enter a letter guess and returns the letter as a rune.
func GetGuess() rune {
	// Prompt player to enter letter
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read input
		print("Enter a letter: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// Check if input is a single lowercase letter
		input = strings.TrimSpace(input)
		if len(input) == 1 && input[0] >= 'a' && input[0] <= 'z' {
			return rune(input[0])
		}

		// Invalid input, try again
		println("Invalid input. Please enter a single lowercase letter.")
	}
}
