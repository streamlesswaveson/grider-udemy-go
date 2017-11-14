package main

import "fmt"

// go run main.go deck.go

func main()  {

	cards := newDeck()

	hand, remaining := deal(cards, 5)

	fmt.Println(hand)
	fmt.Println(remaining)
}

func newCard() string {
	return "Ace of Spades"
}