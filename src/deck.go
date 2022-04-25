package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newCards() (cardDeck deck) {
	cardHouses := []string{"Spades ♠", "Clubs ♣", "Hearts ♥", "Diamonds ♦"}
	cardsNo := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, house := range cardHouses {
		for _, cardType := range cardsNo {
			cardDeck = append(cardDeck, fmt.Sprintf("%s Of %s", cardType, house))
		}
	}
	return
}

func (d deck) print() {
	for _, cards := range d {
		fmt.Println(cards)
	}
}

func (d deck) deal(size int) (dealDeck deck, remainingDeck deck) {
	//for _, card := range d {
	//	if len(dealDeck) == size {
	//		break
	//	}
	//	dealDeck = append(dealDeck, card)
	//}
	//
	//for i, card := range d {
	//	if i < size {
	//		continue
	//	}
	//	remainingDeck = append(remainingDeck, card)
	//}
	//return

	dealDeck = d[:size]
	remainingDeck = d[size:]
	return
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readDeckFromFIle(fileName string) (d deck) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	d = strings.Split(string(content), ",")
	return
}

func (d deck) shuffle() {
	//indexNotMap := hashmap.New()
	//for i := range d {
	//	indexNotMap.Put(i,i)
	//}

	//for i,randI := range indexNotMap.Values() {
	//	newDeck[i] = d[randI.(int)]
	//}

	rand.Seed(time.Now().UnixNano())

	var randomArray []int
	for i := range d {
		randomNumber := rand.Intn(len(d) -1)
		//d[i] =d [randomNumber]
		//d[randomNumber] = card

		d[i], d[randomNumber] = d[randomNumber],d[i]
		randomArray = append(randomArray,randomNumber)
	}

	fmt.Println(randomArray)
}
