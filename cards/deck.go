package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings / string slice

type deck []string

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	// Enumeration Algorithm
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {			// (deck, deck) means we want to return 2 values
	return d[:handSize], d[handSize:]					// d[:handSize] returns the dealt hand
														// d[handSize:] returns the cards remaining in the deck that havent been dealt
}