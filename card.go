package evaluator

import "math/rand"

// NewDeck Build a deck.
func NewDeck() [][]string {
	deck := [][]string{
		{"d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "dT", "dJ", "dQ", "dK", "dA"}, //方塊 258~270
		{"s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "sT", "sJ", "sQ", "sK", "sA"}, //梅花 514~526
		{"h2", "h3", "h4", "h5", "h6", "h7", "h8", "h9", "hT", "hJ", "hQ", "hK", "hA"}, //紅心 770~782
		{"c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "cT", "cJ", "cQ", "cK", "cA"}, //黑桃 1026~1038
	}
	return deck
}

// ShuffleDeck shuffle deck.
func ShuffleDeck(deck [][]string) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			suitRandIndex := rand.Intn(3)
			valueRandIndex := rand.Intn(11)
			tmp := deck[i][j]
			deck[i][j] = deck[suitRandIndex][valueRandIndex]
			deck[suitRandIndex][valueRandIndex] = tmp
		}
	}
}

// DrawCards draw cards.
func DrawCards(deck [][]string, count int) []string {
	hand := make([]string, count)

	for i := 0; i < count; i++ {
		// 將第一張牌加入手牌
		hand[i] = deck[0][0]

		// 從牌組中移除已抽取的牌
		deck[0] = deck[0][1:]

		// 如果該花色的牌已經抽完，則從牌組中移除該花色
		if len(deck[0]) == 0 {
			deck = deck[1:]
		}
	}

	return hand
}
