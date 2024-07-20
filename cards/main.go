package main

import (
	"fmt"
	"reflect"
)

func main() {
  cards := deck {"Ace of Diamonds", newCard()}
  cards = append(cards, "Six of Spades")
  cards.print()
  
  fmt.Println(reflect.TypeOf(cards))
}

func newCard() string {
  return "Five of Diamonds"
}