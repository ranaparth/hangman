package dictionary

import (
	"bufio"
	"os"
)

func getStaticList() []string {
	// Create slice to store words
	var wordList []string

	// Open wordlist file
	file, err := os.Open("data/wordlist.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read words from file and add to wordList slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return wordList
}
