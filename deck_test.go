package main

import (
	"os"
	"testing"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spaced, but got %v", d[0])
	}

	if d[len(d) -1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	os.Remove("_decktesting")
	newDeck().saveToFile("_decktesting")
	d := newDeckFromFile("_decktesting")
	if (len(d) != 16) {
		t.Errorf("Expected 16 cards in deck, got %v", len(d))
	}
	os.Remove("_decktesting")
}