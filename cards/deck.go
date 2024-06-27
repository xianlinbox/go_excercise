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