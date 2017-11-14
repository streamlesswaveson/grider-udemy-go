package main

// go run main.go deck.go

func main()  {

	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}