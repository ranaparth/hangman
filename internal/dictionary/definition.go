package dictionary

import (
	"encoding/json"
	"io"
	"net/http"
)

func getWordDefinition(word string) (string, error) {
	// Make a call to the API
	resp, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse JSON response
	var words []struct {
		Meanings []struct {
			Definitions []struct {
				Definition string `json:"definition"`
			}
		}
	}
	err = json.Unmarshal(body, &words)
	if err != nil {
		return "", err
	}

	return words[0].Meanings[0].Definitions[0].Definition, nil
}
