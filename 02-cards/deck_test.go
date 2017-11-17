package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("expected length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected first card Ace of Spaces, got %v", d[0])
	}

	if d[51] != "King of Clubs" {
		t.Errorf("expected first card King of Clubs, got %v", d[51])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards from loaded deck, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}