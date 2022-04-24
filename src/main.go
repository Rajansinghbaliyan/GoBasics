package main

import "fmt"

func main() {

	pokemonUrl := "https://pokeapi.co/api/v2/pokemon/1"
	pokemonUrlCustom := "https://pokeapi.co/api/v2/pokemon/"

	fmt.Println("Hello World")
	basicVariable()

	// calling functions
	greetHarry := greetings("Harry")
	harry, harryAge := nameAndAge("Harry",31)

	pokemon, err := fetchPokemon(pokemonUrl)

	fetchPokemonInRange(pokemonUrlCustom,1,10)

	fmt.Printf(greetHarry)
	fmt.Printf(harry)
	fmt.Printf(harryAge)
	fmt.Printf("%v",pokemon)

	if err != nil {
		fmt.Printf("%s",err)
	}
}
