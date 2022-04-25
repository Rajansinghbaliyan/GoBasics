package main

import (
	"encoding/json"
	"fmt"
	sll "github.com/emirpasic/gods/lists/singlylinkedlist"
	"net/http"
	"reflect"
)

func greetings(name string) string {
	return fmt.Sprintf("Hello %s how are you?", name)
}

func nameAndAge(name string, age int) (string, string) {
	return fmt.Sprintf("Hello %s", name), fmt.Sprintf("Your age is: %d", age)
}

func fetchPokemonConcurrently(url string, pokemonCh chan Pokemon, errorCh chan error) {
	resp, err := http.Get(url)

	if err != nil {
		errorCh <- fmt.Errorf("there is an error : %s",err)
	}

	decoder := json.NewDecoder(resp.Body)

	var pokemon Pokemon
	err = decoder.Decode(&pokemon)

	if err != nil {
		errorCh <- fmt.Errorf("there is an error : %s",err)
	}

	pokemonCh <- pokemon
}

func fetchPokemon(url string) (pokemon Pokemon, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&pokemon)

	if err != nil {
		err = fmt.Errorf("error will fetching %s :%s", url, err)
		return
	}

	return
}

func fetchPokemonInRange(url string, start int, end int) {
	pokemonCh := make(chan Pokemon)
	errorCh := make(chan error)

	for c := start; c < end; c++ {
		tempUrl := fmt.Sprintf("%s%d",url,c)
		go fetchPokemonConcurrently(tempUrl,pokemonCh,errorCh)
	}

	list := sll.New()
	list.Add(getField(pokemonCh,"Name"))

	fmt.Printf("%v",list)

	for c := start; c < end; c++ {
		fmt.Printf("%v\n", getField(pokemonCh, "Name"))
		fmt.Println()
	}

	for c := start; c < end; c++ {
		fmt.Printf("%v",<-errorCh)
		fmt.Println()
	}
}

func getField( pokemon chan Pokemon, field string) string {
	r := reflect.ValueOf(<- pokemon)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}
