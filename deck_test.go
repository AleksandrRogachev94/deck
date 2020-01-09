package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element to equal Ace of Spades, but got %s", d[0])
	}

	if d[15] != "Four of Clubs" {
		t.Errorf("Expected last element to equal Four of Clubs, but got %s", d[15])
	}
}

func TestSave(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	ld := newDeckFromFile(filename)

	if len(ld) != len(d) {
		t.Errorf("Expected %d cards in deck, got %d", len(ld), len(d))
	}

	os.Remove(filename)
}
