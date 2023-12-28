# go-poker-hand-evaluator

*[English](README.md) ∙ [繁體中文](README_zh-tw.md) ∙ [简体中文](README_zh-cn.md)*

## Introduction
`go-poker-hand-evaluator` is a poker card counter that uses Cactus Kev's Poker Hand Evaluator algorithm. It can evaluate the poker card combination in the player's hand and give the corresponding card strength or card type.

## Prerequisites
Make sure you have installed the Go environment.

## Installation
```bash
go get -u github.com/kinnyhuang529/go-poker-hand-evaluator
```

## Configuration
Import this package in Go code:
```bash
import "github.com/kinnyhuang529/go-poker-hand-evaluator"
```

## Example
```go
package main

import (
	"fmt"
	evaluator "github.com/kinnyhuang529/go-poker-hand-evaluator"
)

func main() {
	//Build deck
	deck := evaluator.NewDeck() 
	//deck: [[d2 d3 d4 d5 d6 d7 d8 d9 dT dJ dQ dK dA] [s2 s3 s4 s5 s6 s7 s8 s9 sT sJ sQ sK sA] [h2 h3 h4 h5 h6 h7 h8 h9 hT hJ hQ hK hA] [c2 c3 c4 c5 c6 c7 c8 c9 cT cJ cQ cK cA]]

	//Shuffle deck
	evaluator.ShuffleDeck(deck)
	//deck: [[cQ d5 c2 c6 s5 hQ s3 s2 h5 s8 cT d2 dA] [h2 c5 c8 h3 d3 dJ hT s4 sQ cK sK h4 sA] [cJ s9 c3 c7 dT sT h9 hK c4 d7 d6 d4 hA] [h8 hJ dK dQ s6 h7 h6 d8 d9 sJ c9 s7 cA]]

	//Draw 2 Cards (hand card)
	hand := evaluator.DrawCards(deck, 2)
	//hand: [cQ d5]

	//Draw 5 Cards (public card)
	public := evaluator.DrawCards(deck, 2)
	//public:  [c2 c6 s5 hQ s3]

	//hand + public
	cards := append(hand, public...)
	//cards:  [cQ d5 c2 c6 s5 hQ s3]
	
	var cardType, cardStrength int
	
	//Evaluator
	cardType, cardStrength = evaluator.Evaluator(public)
	//cardType: 10, cardStrength: 7214

	//HoldemEvaluator
	cardType, cardStrength, highest = evaluator.HoldemEvaluator(cards)
	//cardType: 8, cardStrength: 2794, highest: [cQ d5 c6 s5 hQ]
}
```
### Explanation of cardType

- 0: Exceptional Situation
- 1: Royal Flush
- 2: Straight Flush
- 3: Four of a Kind
- 4: Full House
- 5: Flush
- 6: Straight
- 7: Three of a Kind
- 8: Two Pairs
- 9: One Pair
- 10: High Card

### Explanation of cardStrength
The range is from 1 to 7462, the smaller the number, the stronger the card.

## Links
 - [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)