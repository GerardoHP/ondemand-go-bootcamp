package config

import (
	"testing"
)

// TestInitialValues calls NewConfig to the initial values,
// checking rather if they are the correct ones or not.
func TestInitialValues(t *testing.T) {
	cfg := NewConfig()

	if cfg == nil {
		t.Fatal("Config should't be null")
	}

	if cfg.Port != "8080" {
		t.Fatal("Port should be 8080")
	}

	if cfg.StorageFileName != "pokemons.csv" {
		t.Fatal("Storage file name should be pokemons.csv")
	}
}
