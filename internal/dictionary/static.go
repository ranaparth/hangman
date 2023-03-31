package dictionary

import (
	"bytes"

	_ "embed"
)

//go:embed wordlist.txt
var wordListContent string

func getStaticList() []string {
	// Create slice to store words
	var wordList []string

	// Split file contents into lines
	lines := bytes.Split([]byte(wordListContent), []byte{'\n'})

	// Loop through lines and print each one
	for _, line := range lines {
		wordList = append(wordList, string(line))
	}

	return wordList
}
