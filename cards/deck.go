package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings / string slice

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println("Index:", i, "|| Card:", card)
	}
}