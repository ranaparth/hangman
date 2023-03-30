package main

import (
	"math/rand"
	"time"

	"github.com/ranaparth/hangman/internal/dictionary"
	"github.com/ranaparth/hangman/internal/game"
	"github.com/ranaparth/hangman/internal/input"
	"github.com/ranaparth/hangman/internal/output"
)

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Display game introduction
	output.DisplayGameIntro()
	output.DisplayGameRules()

	// Generate a random word
	word := dictionary.GetWord()

	// Create new game
	g := game.New(word)

	// Game loop
	for !g.GameOver() {
		// Display current game state
		output.DisplayState(g.State())

		// Accept player input
		guess := input.GetGuess()

		// Make guess
		err := g.Guess(guess)
		if err != nil {
			output.DisplayError(err)
		}

		// Display remaining guesses
		output.DisplayRemainingGuesses(g.RemainingGuesses())

		// Display available letters
		output.DisplayAvailableLetters(g.AvailableLetters())
	}

	// Display final game state
	output.DisplayState(g.State())

	// Display game result
	output.DisplayGameOutro(g)
}
