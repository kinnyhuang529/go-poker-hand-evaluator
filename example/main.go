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
	result, rank := evaluator.Evaluator(hand)
	fmt.Println(result) //牌型
	fmt.Println(rank)   //牌型等級 數字越小表示牌型越大
}
