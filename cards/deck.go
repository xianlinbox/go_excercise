package main

import "fmt"

// Define a new type 'deck' which is a slice of strings
type deck []string

// Function 'print' that prints out all the index and value in the deck
func (cards deck) print() {
    for i, card := range cards {
        fmt.Println(i, card)
    }
}

// Function 'newDeck' that creates a new deck of cards
func newDeck() deck {
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}