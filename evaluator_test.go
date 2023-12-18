package evaluator

import (
	"testing"
)

func Test_Evaluator(t *testing.T) {
	var hand1 = []string{"cT", "cJ", "cQ", "cK", "cA"}
	if cardType, cardStrength := Evaluator(hand1); cardType != 1 || cardStrength != 1 {
		t.Error("Test_Evaluator: 皇家同花順錯誤")
		return
	}
	var hand2 = []string{"c3", "c4", "c5", "c6", "c7"}
	if cardType, cardStrength := Evaluator(hand2); cardType != 2 || cardStrength != 8 {
		t.Error("Test_Evaluator: 同花順錯誤")
		return
	}
	var hand3 = []string{"dT", "sT", "hT", "cT", "cA"}
	if cardType, cardStrength := Evaluator(hand3); cardType != 3 || cardStrength != 59 {
		t.Error("Test_Evaluator: 金剛錯誤")
		return
	}
	var hand4 = []string{"cT", "dT", "hT", "c3", "d3"}
	if cardType, cardStrength := Evaluator(hand4); cardType != 4 || cardStrength != 225 {
		t.Error("Test_Evaluator: 葫蘆錯誤")
		return
	}
	var hand5 = []string{"h3", "hA", "h8", "hT", "hK"}
	if cardType, cardStrength := Evaluator(hand5); cardType != 5 || cardStrength != 414 {
		t.Error("Test_Evaluator: 同花錯誤")
		return
	}
	var hand6 = []string{"c3", "d4", "h5", "s6", "c7"}
	if cardType, cardStrength := Evaluator(hand6); cardType != 6 || cardStrength != 1607 {
		t.Error("Test_Evaluator: 順子錯誤")
		return
	}
	var hand7 = []string{"c3", "d3", "s3", "cK", "cA"}
	if cardType, cardStrength := Evaluator(hand7); cardType != 7 || cardStrength != 2336 {
		t.Error("Test_Evaluator: 三條錯誤")
		return
	}
	var hand8 = []string{"c5", "d5", "c3", "d3", "cA"}
	if cardType, cardStrength := Evaluator(hand8); cardType != 8 || cardStrength != 3271 {
		t.Error("Test_Evaluator: 兩對錯誤")
		return
	}
	var hand9 = []string{"c5", "d5", "cT", "d3", "cA"}
	if cardType, cardStrength := Evaluator(hand9); cardType != 9 || cardStrength != 5338 {
		t.Error("Test_Evaluator: 一對錯誤")
		return
	}
	var hand10 = []string{"c3", "c8", "d2", "sT", "hA"}
	if cardType, cardStrength := Evaluator(hand10); cardType != 10 || cardStrength != 6589 {
		t.Error("Test_Evaluator: 高牌錯誤")
		return
	}
}
