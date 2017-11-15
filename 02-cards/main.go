package main


// go run main.go deck.go

func main()  {

	cards := newDeck()

	cards.shuffle()

	cards.print()

}
