package evaluator

import (
	"reflect"
	"testing"
)

var TestDeck = [][]string{
	{"d2", "d3", "d4", "d5", "d6", "d7", "d8", "d9", "dT", "dJ", "dQ", "dK", "dA"}, //方塊 258~270
	{"s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "sT", "sJ", "sQ", "sK", "sA"}, //梅花 514~526
	{"h2", "h3", "h4", "h5", "h6", "h7", "h8", "h9", "hT", "hJ", "hQ", "hK", "hA"}, //紅心 770~782
	{"c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "cT", "cJ", "cQ", "cK", "cA"}, //黑桃 1026~1038
}

func Test_NewDeck(t *testing.T) {
	deck := NewDeck()

	if !reflect.DeepEqual(deck, TestDeck) {
		t.Error("Test_NewDeck: Deck content error! ")
		return
	}
	t.Log("Test_NewDeck: success! ")
}

func Test_ShuffleDeck(t *testing.T) {

	deck := NewDeck()
	ShuffleDeck(deck)
	if reflect.DeepEqual(deck, TestDeck) {
		t.Error("Test_ShuffleDeck: shuffle deck error! ")
		return
	}
	t.Log("Test_ShuffleDeck: success! ")
}

func Test_DrawCards(t *testing.T) {
	deck := NewDeck()
	hand := DrawCards(deck, 5)
	if len(hand) != 5 {
		t.Error("Test_DrawCards: Wrong hand card length! ")
		return
	}
	for _, h := range hand {
		check := false
		for _, d := range TestDeck {
			for _, c := range d {
				if h == c {
					check = true
				}
			}
		}
		if !check {
			t.Error("Test_DrawCards: Wrong hand card! ")
			return
		}
	}
	t.Log("Test_DrawCards: success! ")
}
