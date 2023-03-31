package output

import (
	"fmt"
	"strings"

	"github.com/ranaparth/hangman/internal/game"
)

// DisplayState displays the current state of the game.
func DisplayState(state string) {
	fmt.Println(state)
}

// DisplayError displays an error message.
func DisplayError(err error) {
	fmt.Printf("Oops! %s.\n", err)
}

// DisplayGameIntro displays the game introduction.
func DisplayGameIntro() {
	fmt.Println("Welcome to Hangman!")
}

func DisplayGameRules() {
	fmt.Println("Hangman is a word guessing game. The computer will pick a random word and you will have to guess it one letter at a time. You have 6 guesses. Good luck!")
}

// DisplayRemainingGuesses displays the number of remaining guesses.
func DisplayRemainingGuesses(remaining int) {
	fmt.Printf("You have %d guesses remaining.\n", remaining)
}

// DisplayAvailableLetters displays the available letters that have not been guessed yet.
func DisplayAvailableLetters(letters []rune) {
	fmt.Printf("Available letters: %s\n", strings.Join(strings.Split(string(letters), ""), ", "))
}

// DisplayGameOutro displays the game result.
func DisplayGameOutro(g *game.Game) {
	if g.Won() {
		fmt.Println("Congratulations, you won!")
	} else {
		fmt.Println("Sorry, you lost. The word was", g.Word().String())
		fmt.Println("It means — ", g.Word().Defn())
	}
}
