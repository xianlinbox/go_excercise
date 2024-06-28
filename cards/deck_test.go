package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Call the newDeck function and assign the returned value to a variable
	d := newDeck()

	// Add your assertions here to test the behavior of the newDeck function
	// For example, you can check if the length of the deck is correct or if certain cards are present

	// Example assertion: check if the deck has the correct number of cards
	expectedLength := 16
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
}