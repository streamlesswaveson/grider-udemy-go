package main

// go run main.go deck.go

func main()  {

	cards := newDeck()

	hand, remaining := deal(cards, 5)

	hand.print()
	remaining.print()
}
