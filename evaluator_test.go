package evaluator

import (
	"testing"
)

func Test_Evaluator(t *testing.T) {
	var hand1 = []string{"cT", "cJ", "cQ", "cK", "cA"}
	if Evaluator(hand1) != 1 {
		t.Error("Test_Evaluator: 皇家同花順錯誤")
		return
	}
	var hand2 = []string{"c3", "c4", "c5", "c6", "c7"}
	if Evaluator(hand2) != 2 {
		t.Error("Test_Evaluator: 同花順錯誤")
		return
	}
	var hand3 = []string{"dT", "sT", "hT", "cT", "cA"}
	if Evaluator(hand3) != 3 {
		t.Error("Test_Evaluator: 金剛錯誤")
		return
	}
	var hand4 = []string{"cT", "dT", "hT", "c3", "d3"}
	if Evaluator(hand4) != 4 {
		t.Error("Test_Evaluator: 葫蘆錯誤")
		return
	}
	var hand5 = []string{"h3", "hA", "h8", "hT", "hK"}
	if Evaluator(hand5) != 5 {
		t.Error("Test_Evaluator: 同花錯誤")
		return
	}
	var hand6 = []string{"c3", "d4", "h5", "s6", "c7"}
	if Evaluator(hand6) != 6 {
		t.Error("Test_Evaluator: 順子錯誤")
		return
	}
	var hand7 = []string{"c3", "d3", "s3", "cK", "cA"}
	if Evaluator(hand7) != 7 {
		t.Error("Test_Evaluator: 三條錯誤")
		return
	}
	var hand8 = []string{"c5", "d5", "c3", "d3", "cA"}
	if Evaluator(hand8) != 8 {
		t.Error("Test_Evaluator: 兩對錯誤")
		return
	}
	var hand9 = []string{"c5", "d5", "cT", "d3", "cA"}
	if Evaluator(hand9) != 9 {
		t.Error("Test_Evaluator: 一對錯誤")
		return
	}
	var hand10 = []string{"c3", "c8", "d2", "sT", "hA"}
	if Evaluator(hand10) != 10 {
		t.Error("Test_Evaluator: 高牌錯誤")
		return
	}
}
