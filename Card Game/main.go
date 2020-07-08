package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	//To convert string into byte slice
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting)) //Output:  [72 105 32 116 104 101 114 101 33]

	// cards := newDeck()
	// fmt.Println(cards.toString()) //Output: Ace of Spades,Two of Spades,Three of Spades,Four of Spades,Ace of Diamonds,Two of Diamonds,Three of Diamonds,Four of Diamonds,Ace of Hearts,Two of Hearts,Three of Hearts,Four of Hearts,Ace of Clubs,Two of Clubs,Three of Clubs,Four of Clubs

	//To save deck to a file on harddisk
	/*	cards := newDeck()
		cards.saveTofile("my_cards") //Output: Generates a new file which displays all the cards
	*/

	//To get a deck from file
	/*	cards := newDeckFromFile("my_cards")
		cards.print() //Output: 0 Ace of Spades	1 Two of Spades	2 Three of Spades	3 Four of Spades....

		cards = newDeckFromFile("file")
		cards.print() //Output:   Error: open file: The system cannot find the file specified.
		//exit status 1
	*/

	cards := newDeck()
	cards.pseduoshuffle() //Ouput: The last four values are always same no matter how many times you run the code
	cards.shuffle()       //Output: Randomised deck of cards
	cards.print()

}
