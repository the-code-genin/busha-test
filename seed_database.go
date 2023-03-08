package main

import (
	"context"

	"github.com/the-code-genin/busha-test/database/repositories"
)

// Fetch data from swapi.dev and seed the redis system on first launch
func SeedDatabase(
	ctx context.Context,
) error {
	systemRepo, err := repositories.NewSystemRepository(ctx)
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