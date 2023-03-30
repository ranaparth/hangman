package game

import (
	"errors"
	"strings"

	"github.com/ranaparth/hangman/internal/dictionary"
)

// Game represents the state of a game of Hangman.
type Game struct {
	// word is the word to be guessed
	word *dictionary.Word
	// guessed is a map of letters that have been guessed
	guessed map[rune]bool
	// remaining is the number of guesses remaining
	remaining int
	// lettersLeft is a map of letters that have not been guessed yet
	lettersLeft map[rune]bool
}

// list of all lower case characters
const allLowercaseLetters = "abcdefghijklmnopqrstuvwxyz"

// New creates a new game of Hangman.
func New(word *dictionary.Word) *Game {
	// Initialize guessed map
	guessed := make(map[rune]bool)

	// Initialize lettersLeft map
	lettersLeft := make(map[rune]bool)
	for _, letter := range allLowercaseLetters {
		lettersLeft[letter] = true
	}

	// Return new game
	return &Game{
		word:        word,
		guessed:     guessed,
		remaining:   6,
		lettersLeft: lettersLeft,
	}
}

// State returns the current state of the game.
func (g *Game) State() string {
	// Create array of characters representing the word, with '_' for unknown letters
	state := make([]string, len(g.word.String()))
	for i, letter := range g.word.String() {
		if g.guessed[letter] {
			state[i] = string(letter)
		} else {
			state[i] = "_"
		}
	}

	// Return current state as string
	return strings.Join(state, " ")
}

// Guess makes a guess in the game, updating the game state accordingly.
func (g *Game) Guess(letter rune) error {
	// Convert letter to lowercase
	letter = toLower(letter)

	// Check if letter has already been guessed
	if g.guessed[letter] {
		return errors.New("letter already guessed")
	}

	// Mark letter as guessed
	g.guessed[letter] = true

	// Check if letter is in word
	if !strings.ContainsRune(g.word.String(), letter) {
		g.remaining--
	}

	// Remove letter from lettersLeft
	delete(g.lettersLeft, letter)

	// Check if game is over
	if g.GameOver() {
		return errors.New("game over")
	}

	return nil
}

// Won returns true if the game has been won.
func (g *Game) Won() bool {
	for _, letter := range g.word.String() {
		if !g.guessed[letter] {
			return false
		}
	}
	return true
}

// Word returns the word being guessed.
func (g *Game) Word() *dictionary.Word {
	return g.word
}

// GameOver returns true if the game is over.
func (g *Game) GameOver() bool {
	return g.remaining == 0 || g.Won()
}

// RemainingGuesses returns the number of remaining guesses.
func (g *Game) RemainingGuesses() int {
	return g.remaining
}

// AvailableLetters returns the letters that have not been guessed yet.
func (g *Game) AvailableLetters() []rune {
	letters := make([]rune, 0, len(g.lettersLeft))
	for letter := range g.lettersLeft {
		letters = append(letters, letter)
	}
	return letters
}

// toLower returns the lowercase equivalent of the input rune.
func toLower(letter rune) rune {
	if letter >= 'A' && letter <= 'Z' {
		return letter + 'a' - 'A'
	}
	return letter
}
