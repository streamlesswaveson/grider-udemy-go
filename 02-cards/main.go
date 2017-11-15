package main


// go run main.go deck.go

func main()  {

	cards := newDeck()

	//fmt.Println(cards.toString())

	cards.saveToFile("testcards.txt")

	othercards := newDeckFromFile("testcards.txt")

	othercards.print()

}
