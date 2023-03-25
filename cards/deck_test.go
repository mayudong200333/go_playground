package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "AceofSpades" {
		t.Errorf("Expected deck length of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	laodedDeck := newDeckFromFile("_decktesting")

	if len(laodedDeck) != 16 {
		t.Errorf("Expected deck length of 20, but got %v", len(laodedDeck))
	}

	os.Remove("_decktesting")
}
