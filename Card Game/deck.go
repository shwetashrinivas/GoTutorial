package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Craete a new type of deck which is a slice of strings

type deck []string

//Create and return a list of playing cards, essentially an array of strings
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Sisplays out the contents of a deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Creates a "hand" of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Converts the slice of string to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Saves a list of cards to a file on the local machine
func (d deck) saveTofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Loads a list of cards from the local machine
//err will have a value of 'nil' if nothing goes wrong
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// 1- Log the errorand return a call to newdeck()
		//Or 2 - Log the error and entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//We have a []byte and we need a []string
	//fmt.Println(string(bs)) //Output: Ace of Spades , Two of Spades, ...
	s := strings.Split(string(bs), ",")

	//Now we have []string and we need a deck
	return deck(s)

}

//Shuffles all the cards in a deck
/*Logic:  for each index, card in cards
Generate a random number between 0 and len(cards) -1
Swap the current card and the card at cards[randomNumber] */
func (d deck) pseduoshuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1) //pseudo random generator where we cannot specify a source
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}

//type Rand is a source of random numbers
//func New returns new Rand that uses random values from src to generate other random values

//we need to pass an int64 value to newSource(). So we use UnixNano which returns t of type Time as a Unix time
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
