package main

import (
	"fmt"
	evaluator "github.com/kinnyhuang529/go-poker-hand-evaluator/table"
)

func main() {
	//建立牌組
	deck := evaluator.NewDeck()
	fmt.Println(deck)
	//洗牌
	evaluator.ShuffleDeck(deck)
	fmt.Println(deck)
	//抽牌
	hand := evaluator.DrawCards(deck, 5)
	fmt.Println(hand)
	//開始算牌
	result := evaluator.Evaluator(hand)
	fmt.Println(result)
}
