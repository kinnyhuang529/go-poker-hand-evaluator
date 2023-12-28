package main

import (
	"fmt"
	evaluator "github.com/kinnyhuang529/go-poker-hand-evaluator"
)

func main() {
	//Build deck
	deck := evaluator.NewDeck()
	fmt.Println("deck: ", deck)

	//Shuffle deck
	evaluator.ShuffleDeck(deck)
	fmt.Println("shuffle deck: ", deck)

	//Draw 2 Cards (hand card)
	hand := evaluator.DrawCards(deck, 2)
	fmt.Println("hand: ", hand)

	//Draw 5 Cards (public card)
	public := evaluator.DrawCards(deck, 5)
	fmt.Println("public: ", public)

	//hand + public
	cards := append(hand, public...)
	fmt.Println("cards: ", cards)

	//The remaining number of cards in the deck.
	fmt.Println("The remaining number of cards in the deck: ", evaluator.NumberOfCardsInTheDeck(deck))

	var cardType, cardStrength int
	//Start evaluating cards (5 cards)
	cardType, cardStrength = evaluator.Evaluator(public)
	fmt.Println("5 cards card type: ", cardType)         //card type
	fmt.Println("5 cards card strength: ", cardStrength) //card strength. The smaller the card strength value, the larger the card type.

	//Start evaluating cards (7 cards)
	ttt := []string{"cQ", "d5", "c2", "c6", "s5", "hQ", "s3"}
	cardType, cardStrength, highest := evaluator.HoldemEvaluator(ttt)
	fmt.Println("7 cards card type: ", cardType)         //card type
	fmt.Println("7 cards card strength: ", cardStrength) //card strength. The smaller the card strength value, the larger the card type.
	fmt.Println("7 cards best 5 card: ", highest)        //highest. Best Five-Card.
}
