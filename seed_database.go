package main

import (
	"fmt"
	"sync"

	"github.com/inconshreveable/log15"
	"github.com/the-code-genin/busha-test/database"
	"github.com/the-code-genin/busha-test/internal"
	"github.com/the-code-genin/busha-test/swapi"
)

// Fetch data from swapi.dev and seed the redis system on first launch
func SeedDatabase(
	ctx *internal.AppContext,
) error {
	systemRepo, err := database.NewSystemRepository(ctx)
	if err != nil {
		return err
	}

	// Check if the database seeding has been done
	seeded, err := systemRepo.GetSeeded()
	if err != nil {
		return err
	} else if seeded {
		return nil
	}

	// Fetch the films
	films, err := swapi.FetchMovies(ctx)
	if err != nil {
		return err
	}

	// Create a set of unique character URLs extracts from all films
	characterURLs := make([]string, 0)
	for _, film := range films {
		for _, newURL := range film.CharacterURLs {
			exists := false
			for _, existingURL := range characterURLs {
				if existingURL == newURL {
					exists = true
					break
				}
			}
			if !exists {
				characterURLs = append(characterURLs, newURL)
			}
		}
	}

	// Fetch all unique characters from swapi.dev in parallel
	characters := make([]swapi.Character, 0)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, url := range characterURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			character, err := swapi.FetchCharacter(url)
			if err != nil {
				log15.Error(err.Error())
				return
			}

			mutex.Lock()
			defer mutex.Unlock()

			characters = append(characters, character)
		}(url)
	}
	wg.Wait()

	// Unable to fetch all characters from swapi.dev
	if len(characters) != len(characterURLs) {
		return fmt.Errorf(
			"unable to fetch all characters from swapi.dev. Fetched %d expected %d",
			len(characters),
			len(characterURLs),
		)
	}

	return nil
}
