package main

import "fmt"

func main() {

	//pokemonUrl := "https://pokeapi.co/api/v2/pokemon/1"
	//pokemonUrlCustom := "https://pokeapi.co/api/v2/pokemon/"
	//
	//fmt.Println("Hello World")
	//basicVariable()

	// calling functions
	//greetHarry := greetings("Harry")
	//harry, harryAge := nameAndAge("Harry",31)
	//
	//pokemon, err := fetchPokemon(pokemonUrl)
	//
	//fetchPokemonInRange(pokemonUrlCustom,1,10)
	//
	//fmt.Printf(greetHarry)
	//fmt.Printf(harry)
	//fmt.Printf(harryAge)
	//fmt.Printf("%v",pokemon)
	//
	//if err != nil {
	//	fmt.Printf("%s",err)
	//}

	cardDeck := newCards()

	cardDeck.print()
	dealDeck,remainingDeck := cardDeck.deal(5)


	fmt.Println(dealDeck)
	fmt.Println(len(dealDeck), len(remainingDeck), len(cardDeck))

	err := cardDeck.saveToFile("card_deck.txt")
	err = dealDeck.saveToFile("deal_deck.txt")
	err = remainingDeck.saveToFile("remaining_deck.txt")

	fmt.Println(err)

	fileFromDisk := readDeckFromFIle("deal_deck.txt")

	fileFromDisk.print()

	cardDeck.print()

	fmt.Println("..................................Shuffled deck............................................")
	cardDeck.shuffle()

	cardDeck.print()

}
