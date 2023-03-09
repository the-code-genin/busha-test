package swapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/the-code-genin/busha-test/internal"
)

type Film struct {
	Title         string   `json:"title"`
	OpeningCrawl  string   `json:"opening_crawl"`
	ReleaseDate   string   `json:"release_date"`
	CharacterURLs []string `json:"characters"`
}

// Fetches movies from swapi.dev
func FetchMovies(ctx *internal.AppContext) ([]Film, error) {
	// Get the films URL
	config, err := ctx.GetConfig()
	if err != nil {
		return nil, err
	}
	filmsURL, err := config.GetSWAPIFilmsURL()
	if err != nil {
		return nil, err
	}

	// Fetch films
	resp, err := http.DefaultClient.Get(filmsURL)
	if err != nil {
		return nil, err
	}

	// Read the body
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Status code should be 200
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	// Decode the body
	var result struct {
		Results []Film `json:"results"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Results, nil
}
