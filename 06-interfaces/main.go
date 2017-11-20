package main

import (
	"fmt"
)

// do not need to explicitly call out that the bots implement this
type bot interface {
	getGreeting() string
	//getVersion() float64 // this will break the program b/c the eng/span bots do not implement the method
}

type englishBot struct {

}

type spanishBot struct {

}

func main()  {

	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)

}

func (eb englishBot) getGreeting() string {
	return "hey there"
}

func (sp spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

