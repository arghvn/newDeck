package main

import "fmt"

func main() {

	cards := newDeck ()

	fmt.Println(cards.tostring())
	
}

cards.savetofile("my_cards")

//If we write in the terminal go run main.go deck.go
//output = a new file named my_cards
