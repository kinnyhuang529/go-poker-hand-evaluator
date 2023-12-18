package main

import (
	"fmt"
	evaluator "github.com/kinnyhuang529/go-poker-hand-evaluator"
)

func main() {
	//建立牌組
	deck := evaluator.NewDeck()
	fmt.Println("deck: ", deck)
	//洗牌
	evaluator.ShuffleDeck(deck)
	fmt.Println("shuffle deck: ", deck)
	//手牌
	hand := evaluator.DrawCards(deck, 2)
	fmt.Println("hand: ", hand)
	//公牌
	public := evaluator.DrawCards(deck, 5)
	fmt.Println("public: ", public)
	//把手牌+公牌
	cards := append(hand, public...)
	fmt.Println("cards: ", cards)

	var cardType, cardStrength int
	//開始算牌 5張
	cardType, cardStrength = evaluator.Evaluator(public)
	fmt.Println("5 cards card type: ", cardType)         //牌型
	fmt.Println("5 cards card strength: ", cardStrength) //牌力 牌力數值越小表示牌型越大
	//開始算牌 7張
	cardType, cardStrength = evaluator.HoldemEvaluator(cards)
	fmt.Println("7 cards card type: ", cardType)         //牌型
	fmt.Println("7 cards card strength: ", cardStrength) //牌力 牌力數值越小表示牌型越大
}
