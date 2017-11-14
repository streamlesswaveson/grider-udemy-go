package main

import "fmt"

// new type 'deck' extending from []string
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades","Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace","Two", "Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}

	for _, s := range suits{
		for _, v := range values {
			cards = append(cards, fmt.Sprint(v, " of ", s))
		}
	}
	return cards
}

// any variable of type 'deck' now gets the print() function (receiver)
func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}