package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Tow", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func(d deck) print() {
	for i, card := range  d {
		fmt.Println(i, card);
	}
}