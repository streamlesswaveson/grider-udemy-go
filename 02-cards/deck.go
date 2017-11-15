package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()),0644)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if nil != err {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}