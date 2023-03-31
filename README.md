## Hangman
Welcome to Hangman, a classic guessing game where you try to guess a secret word before running out of chances. This implementation of Hangman is written in Go and can be easily installed using the go install command.
Installation

### Installation
To install Hangman, you'll need to have Go 1.16 or later installed on your system. Once you have Go installed, simply run the following command in your terminal:
```bash
go install github.com/ranaparth/hangman/cmd/hangman@v0.1.0
```
This will download and install the Hangman game to your $GOPATH/bin directory, which you can then run from the command line.

### How to Play
To start a game of Hangman, simply run the hangman command from your terminal. The game will choose a random word from its built-in word list, and display a series of dashes representing the letters in the word.

Your goal is to guess the secret word by guessing one letter at a time. To make a guess, simply type a letter and press enter. If the letter is in the word, it will be revealed in the word display. If the letter is not in the word, you will lose one of your remaining chances.

You have a total of 6 chances to guess the word. If you run out of chances before guessing the word, the game is over and you lose. If you guess the word before running out of chances, you win!

### Customization
If you want to customize the game to use your own word list, you can do so by editing the data/wordlist.txt file in the Hangman source code. Simply add one word per line to the file, save it, and rebuild the game using the go install command.

### TODO
- Implement built-in wordlist like Movies, AWS, etc.
- Input custom flags for total attempt & wordlist
- Improve the UI on terminal for lost and won