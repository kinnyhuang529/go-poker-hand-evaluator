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
	//deck: [[d2 hQ h3 h8 s7 c5 h7 d5 sQ h4 sJ h5 dA] [c2 cK hK c3 cT d8 s9 dT d7 h9 cJ s8 sA] [sK c6 c7 d6 cQ hT s5 dQ c8 c9 s4 s6 hA] [h6 d9 dK h2 d3 c4 s2 s3 hJ d4 sT dJ cA]]
	
	//抽牌
	hand := evaluator.DrawCards(deck, 5)
	//hand: [d2 hQ h3 h8 s7]

	//算牌器
	result, rank := evaluator.Evaluator(hand)
	//result: 10, rank: 7191
}
```
### result 說明

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

### rank 說明
範圍是1到7462，數字越小牌力越強。

## Links
- [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)