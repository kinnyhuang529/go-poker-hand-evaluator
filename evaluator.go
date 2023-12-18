package evaluator

import (
	"github.com/kinnyhuang529/go-poker-hand-evaluator/tables"
	"sort"
)

var (
	Deck = [][]string{
		{"d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "dT", "dJ", "dQ", "dK", "dA"}, //方塊 258~270
		{"s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "sT", "sJ", "sQ", "sK", "sA"}, //梅花 514~526
		{"h2", "h3", "h4", "h5", "h6", "h7", "h8", "h9", "hT", "hJ", "hQ", "hK", "hA"}, //紅心 770~782
		{"c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "cT", "cJ", "cQ", "cK", "cA"}, //黑桃 1026~1038
	}
	Primes     = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	SuitsValue = []int{4, 8, 2, 1} //方塊, 梅花, 紅心, 黑桃
	Poker      = map[string]Card{}
)

type Card struct {
	B     int
	Cdhs  int
	R     int
	P     int
	Value int
}

// 先算出所有牌的value
func init() {
	for i, cards := range Deck {
		for j, card := range cards {
			b := 1 << j << 16
			cdhs := SuitsValue[i] << 12
			r := j << 8
			p := Primes[j]
			value := b | cdhs | r | p
			key := card
			Poker[key] = Card{
				B:     b,
				Cdhs:  cdhs,
				R:     r,
				P:     p,
				Value: value,
			}
		}
	}
}

// Evaluator Cactus Kev's Poker Hand Evaluator 演算法
func Evaluator(input []string) (int, int) {
	//牌數一定要是5張
	if len(input) != 5 {
		return 0, 0
	}

	var key int
	//照演算法 做或閘後取前16位
	r := (Poker[input[0]].Value | Poker[input[1]].Value | Poker[input[2]].Value | Poker[input[3]].Value | Poker[input[4]].Value) >> 16
	//判斷是不是同花
	if Poker[input[0]].Value&Poker[input[1]].Value&Poker[input[2]].Value&Poker[input[3]].Value&Poker[input[4]].Value&0xF000 != 0 {
		//同花的話找同花表
		key = tables.Flushes[r]
	} else {
		//反之找unique5表
		key = tables.Unique5[r]
		//如果不是順子或高牌
		if key == 0 {
			//質數相成
			prime := Poker[input[0]].P * Poker[input[1]].P * Poker[input[2]].P * Poker[input[3]].P * Poker[input[4]].P
			//先從products表找出index
			productsIndex := sort.SearchInts(tables.Products, prime)
			//再去value表找出對應的key
			key = tables.Values[productsIndex]
		}
	}

	//最後看他是什麼牌型
	switch {
	case key == 0:
		return 0, key // 例外情況
	case key == 1:
		return 1, key // 皇家同花順
	case key <= 10:
		return 2, key // 同花順
	case key <= 166:
		return 3, key // 金剛
	case key <= 322:
		return 4, key // 葫蘆
	case key <= 1599:
		return 5, key // 同花
	case key <= 1609:
		return 6, key // 順子
	case key <= 2467:
		return 7, key // 三條
	case key <= 3325:
		return 8, key // 兩對
	case key <= 6185:
		return 9, key // 一對
	case key <= 7462:
		return 10, key // 高牌
	default:
		return 0, key // 例外情況
	}
}
