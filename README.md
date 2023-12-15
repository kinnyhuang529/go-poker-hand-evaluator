# go-poker-hand-evaluator

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
	//Create deck
	deck := evaluator.NewDeck() 
	//deck: [[d2 d3 d4 d5 d6 d7 d8 d9 dT dJ dQ dK dA] [s2 s3 s4 s5 s6 s7 s8 s9 sT sJ sQ sK sA] [h2 h3 h4 h5 h6 h7 h8 h9 hT hJ hQ hK hA] [c2 c3 c4 c5 c6 c7 c8 c9 cT cJ cQ cK cA]]

	//Shuffle deck
	evaluator.ShuffleDeck(deck)
	//deck: [[d2 hQ h3 h8 s7 c5 h7 d5 sQ h4 sJ h5 dA] [c2 cK hK c3 cT d8 s9 dT d7 h9 cJ s8 sA] [sK c6 c7 d6 cQ hT s5 dQ c8 c9 s4 s6 hA] [h6 d9 dK h2 d3 c4 s2 s3 hJ d4 sT dJ cA]]
	
	//Draw Cards
	hand := evaluator.DrawCards(deck, 5)
	//hand: [d2 hQ h3 h8 s7]

	//Evaluator
	result, rank := evaluator.Evaluator(hand)
	//result: 10, rank: 7191
}
```
### Explanation of result

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

### Explanation of rank
The range is from 1 to 7462, the smaller the number, the stronger the card.

## Links
 - [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)