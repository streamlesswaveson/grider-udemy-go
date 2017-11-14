package main

import "fmt"

// go run main.go deck.go

func main()  {

	cards := deck{newCard(), "Ace of Base"}
	fmt.Println(cards)
}

func newCard() string {
	return "Ace of Spades"
}