package main

import (
	"github.com/the-code-genin/busha-test/database"
	"github.com/the-code-genin/busha-test/internal"
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

	return nil
}
