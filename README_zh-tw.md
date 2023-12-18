# go-poker-hand-evaluator

*[English](README.md) ∙ [繁體中文](README_zh-tw.md) ∙ [简体中文](README_zh-cn.md)*

## 介紹
`go-poker-hand-evaluator` 是一個撲克牌算牌器，使用 Cactus Kev's Poker Hand Evaluator 演算法。 它可以計算玩家手中的撲克牌組合併給出相應的牌力或牌型。

## 先決條件
確保您已經安裝了Go環境。

## 安裝
```bash
go get -u github.com/kinnyhuang529/go-poker-hand-evaluator
```

## 配置
在 Go 程式碼中匯入此套件：
```bash
import "github.com/kinnyhuang529/go-poker-hand-evaluator"
```

## 範例
```go
package main

import (
	"fmt"
	evaluator "github.com/kinnyhuang529/go-poker-hand-evaluator"
)

func main() {
	//建立牌組
	deck := evaluator.NewDeck()
	//deck: [[d2 d3 d4 d5 d6 d7 d8 d9 dT dJ dQ dK dA] [s2 s3 s4 s5 s6 s7 s8 s9 sT sJ sQ sK sA] [h2 h3 h4 h5 h6 h7 h8 h9 hT hJ hQ hK hA] [c2 c3 c4 c5 c6 c7 c8 c9 cT cJ cQ cK cA]]

	//洗牌
	evaluator.ShuffleDeck(deck)
	//deck: [[cQ d5 c2 c6 s5 hQ s3 s2 h5 s8 cT d2 dA] [h2 c5 c8 h3 d3 dJ hT s4 sQ cK sK h4 sA] [cJ s9 c3 c7 dT sT h9 hK c4 d7 d6 d4 hA] [h8 hJ dK dQ s6 h7 h6 d8 d9 sJ c9 s7 cA]]

	//抽兩張牌 (手牌)
	hand := evaluator.DrawCards(deck, 2)
	//hand: [cQ d5]

	//抽五張牌 (公牌)
	public := evaluator.DrawCards(deck, 2)
	//public:  [c2 c6 s5 hQ s3]

	//手牌 + 公牌
	cards := append(hand, public...)
	//cards:  [cQ d5 c2 c6 s5 hQ s3]

	var cardType, cardStrength int

	//算牌器
	cardType, cardStrength = evaluator.Evaluator(public)
	//cardType: 10, cardStrength: 7214

	//德州撲克算牌器
	cardType, cardStrength = evaluator.HoldemEvaluator(cards)
	//cardType: 8, cardStrength: 2794
}
```
### cardType 說明

- 0: 例外情況
- 1: 皇家同花順
- 2: 同花順
- 3: 四條/鐵支
- 4: 葫蘆
- 5: 同花
- 6: 順子
- 7: 三條
- 8: 兩對
- 9: 一對
- 10: 高牌

### cardStrength 說明
範圍是1到7462，數字越小牌力越強。

## Links
- [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)