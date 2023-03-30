package dictionary

import (
	"math/rand"
)

type Word struct {
	word       string
	definition string
}

func newWord(word, definition string) *Word {
	return &Word{
		word:       word,
		definition: definition,
	}
}

func (w *Word) String() string {
	return w.word
}

func (w *Word) Defn() string {
	return w.definition
}

func GetWord() *Word {
	return getRandomFromStaticList()
}

func getRandomFromStaticList() *Word {
	wordList := getStaticList()
	word := wordList[rand.Intn(len(wordList))]

	definition, _ := getWordDefinition(word)

	return newWord(word, definition)
}
