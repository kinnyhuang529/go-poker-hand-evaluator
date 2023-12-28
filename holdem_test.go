package evaluator

import (
	"reflect"
	"testing"
)

func Test_HoldemEvaluator(t *testing.T) {
	var hand1 = []string{"cT", "cJ", "c2", "cQ", "d3", "cK", "cA"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand1); cardType != 1 || cardStrength != 1 || !reflect.DeepEqual(highest, []string{"cT", "cJ", "cQ", "cK", "cA"}) {
		t.Error("Test_HoldemEvaluator: Royal Flush error! ")
		return
	}
	var hand2 = []string{"c3", "c4", "cT", "hK", "c5", "c6", "c7"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand2); cardType != 2 || cardStrength != 8 || !reflect.DeepEqual(highest, []string{"c3", "c4", "c5", "c6", "c7"}) {
		t.Error("Test_HoldemEvaluator: Straight Flush error! ")
		return
	}
	var hand3 = []string{"dT", "sT", "hT", "s3", "d2", "cT", "cA"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand3); cardType != 3 || cardStrength != 59 || !reflect.DeepEqual(highest, []string{"dT", "sT", "hT", "cT", "cA"}) {
		t.Error("Test_HoldemEvaluator: Four of a Kind error! ")
		return
	}
	var hand4 = []string{"d2", "cT", "dT", "c8", "hT", "c3", "d3"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand4); cardType != 4 || cardStrength != 225 || !reflect.DeepEqual(highest, []string{"cT", "dT", "hT", "c3", "d3"}) {
		t.Error("Test_HoldemEvaluator: Full House error! ")
		return
	}
	var hand5 = []string{"h3", "d5", "hA", "h8", "s8", "hT", "hK"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand5); cardType != 5 || cardStrength != 414 || !reflect.DeepEqual(highest, []string{"h3", "hA", "h8", "hT", "hK"}) {
		t.Error("Test_HoldemEvaluator: Flush error! ")
		return
	}
	var hand6 = []string{"c3", "d4", "dK", "h5", "s6", "c7", "h8"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand6); cardType != 6 || cardStrength != 1606 || !reflect.DeepEqual(highest, []string{"d4", "h5", "s6", "c7", "h8"}) {
		t.Error("Test_HoldemEvaluator: Straight error! ")
		return
	}
	var hand7 = []string{"c3", "d9", "h7", "d3", "s3", "cK", "cA"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand7); cardType != 7 || cardStrength != 2336 || !reflect.DeepEqual(highest, []string{"c3", "d3", "s3", "cK", "cA"}) {
		t.Error("Test_HoldemEvaluator: Three of a Kind error! ")
		return
	}
	var hand8 = []string{"c5", "d5", "d8", "c3", "cK", "d3", "cA"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand8); cardType != 8 || cardStrength != 3271 || !reflect.DeepEqual(highest, []string{"c5", "d5", "c3", "d3", "cA"}) {
		t.Error("Test_HoldemEvaluator: Two Pairs error! ")
		return
	}
	var hand9 = []string{"c5", "d5", "cT", "c8", "d6", "d3", "cA"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand9); cardType != 9 || cardStrength != 5334 || !reflect.DeepEqual(highest, []string{"c5", "d5", "cT", "c8", "cA"}) {
		t.Error("Test_HoldemEvaluator: One Pair error! ")
		return
	}
	var hand10 = []string{"c3", "c8", "d2", "cK", "dJ", "sT", "hA"}
	if cardType, cardStrength, highest := HoldemEvaluator(hand10); cardType != 10 || cardStrength != 6231 || !reflect.DeepEqual(highest, []string{"c8", "cK", "dJ", "sT", "hA"}) {
		t.Error("Test_HoldemEvaluator: High Card error! ")
		return
	}
	t.Log("Test_HoldemEvaluator: success! ")
}
