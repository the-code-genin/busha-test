package swapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Character struct {
	Name      string `json:"name"`
	Height    int    `json:"height"`
	Mass      int    `json:"mass"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}

// Fetch a swapi.dev character by the character URL
func FetchCharacter(url string) (Character, error) {
	var character Character

	// Fetch films
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return character, err
	}

	// Read the body
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return character, err
	}

	// Status code should be 200
	if resp.StatusCode != http.StatusOK {
		return character, errors.New(string(body))
	}

	// Decode the body
	if err := json.Unmarshal(body, &character); err != nil {
		return character, err
	}

	return character, nil
}
