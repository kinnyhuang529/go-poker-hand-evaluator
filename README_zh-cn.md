# go-poker-hand-evaluator

*[English](README.md) ∙ [繁體中文](README_zh-tw.md) ∙ [简体中文](README_zh-cn.md)*

## 介绍
`go-poker-hand-evaluator` 是一个扑克牌算牌器，使用 Cactus Kev's Poker Hand Evaluator 演算法。它可以计算玩家手中的扑克牌组合并给出相应的牌力或牌型。

## 先决条件
确保您已经安装了Go环境。

## 安装
```bash
go get -u github.com/kinnyhuang529/go-poker-hand-evaluator
```

## 配置
在 Go 代码中导入此包：
```bash
import "github.com/kinnyhuang529/go-poker-hand-evaluator"
```

## 范例
```go
package main

import (
	"fmt"
	evaluator "github.com/kinnyhuang529/go-poker-hand-evaluator"
)

func main() {
	//建立牌组
	deck := evaluator.NewDeck()
	//deck: [[d2 d3 d4 d5 d6 d7 d8 d9 dT dJ dQ dK dA] [s2 s3 s4 s5 s6 s7 s8 s9 sT sJ sQ sK sA] [h2 h3 h4 h5 h6 h7 h8 h9 hT hJ hQ hK hA] [c2 c3 c4 c5 c6 c7 c8 c9 cT cJ cQ cK cA]]

	//洗牌
	evaluator.ShuffleDeck(deck)
	//deck: [[cQ d5 c2 c6 s5 hQ s3 s2 h5 s8 cT d2 dA] [h2 c5 c8 h3 d3 dJ hT s4 sQ cK sK h4 sA] [cJ s9 c3 c7 dT sT h9 hK c4 d7 d6 d4 hA] [h8 hJ dK dQ s6 h7 h6 d8 d9 sJ c9 s7 cA]]

	//抽两张牌 (手牌)
	hand := evaluator.DrawCards(deck, 2)
	//hand: [cQ d5]

	//抽五张牌 (公牌)
	public := evaluator.DrawCards(deck, 2)
	//public:  [c2 c6 s5 hQ s3]

	//手牌 + 公牌
	cards := append(hand, public...)
	//cards:  [cQ d5 c2 c6 s5 hQ s3]

	var cardType, cardStrength int

	//算牌器
	cardType, cardStrength = evaluator.Evaluator(public)
	//cardType: 10, cardStrength: 7214

	//德州扑克算牌器
	cardType, cardStrength, highest = evaluator.HoldemEvaluator(cards)
	//cardType: 8, cardStrength: 2794, highest: [cQ d5 c6 s5 hQ]
}
```
### cardType 说明

- 0: 例外情况
- 1: 皇家同花顺
- 2: 同花顺
- 3: 四条/金刚
- 4: 葫芦
- 5: 同花
- 6: 顺子
- 7: 三条
- 8: 两对
- 9: 一对
- 10: 高牌

### cardStrength 说明
范围是1到7462，数字越小牌力越强。

## Links
- [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)