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
	//deck: [[d2 hQ h3 h8 s7 c5 h7 d5 sQ h4 sJ h5 dA] [c2 cK hK c3 cT d8 s9 dT d7 h9 cJ s8 sA] [sK c6 c7 d6 cQ hT s5 dQ c8 c9 s4 s6 hA] [h6 d9 dK h2 d3 c4 s2 s3 hJ d4 sT dJ cA]]
	
	//抽牌
	hand := evaluator.DrawCards(deck, 5)
	//hand: [d2 hQ h3 h8 s7]

	//算牌器
	result, rank := evaluator.Evaluator(hand)
	//result: 10, rank: 7191
}
```
### result 说明

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

### rank 说明
范围是1到7462，数字越小牌力越强。

## Links
- [Cactus Kev's Poker Hand Evaluator](https://suffe.cool/poker/evaluator.html)