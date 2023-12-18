package evaluator

// 7張牌所有排列組合
var t7c5 = [21][7]int{
	{0, 1, 2, 3, 4, 5, 6},
	{0, 1, 2, 3, 5, 4, 6},
	{0, 1, 2, 3, 6, 4, 5},
	{0, 1, 2, 4, 5, 3, 6},
	{0, 1, 2, 4, 6, 3, 5},
	{0, 1, 2, 5, 6, 3, 4},
	{0, 1, 3, 4, 5, 2, 6},
	{0, 1, 3, 4, 6, 2, 5},
	{0, 1, 3, 5, 6, 2, 4},
	{0, 1, 4, 5, 6, 2, 3},
	{0, 2, 3, 4, 5, 1, 6},
	{0, 2, 3, 4, 6, 1, 5},
	{0, 2, 3, 5, 6, 1, 4},
	{0, 2, 4, 5, 6, 1, 3},
	{0, 3, 4, 5, 6, 1, 2},
	{1, 2, 3, 4, 5, 0, 6},
	{1, 2, 3, 4, 6, 0, 5},
	{1, 2, 3, 5, 6, 0, 4},
	{1, 2, 4, 5, 6, 0, 3},
	{1, 3, 4, 5, 6, 0, 2},
	{2, 3, 4, 5, 6, 0, 1},
}

// HoldemEvaluator 7 cards to calculate card type and strength
func HoldemEvaluator(input []string) (int, int) {
	//牌數一定要是7張
	if len(input) != 7 {
		return 0, 0
	}

	var tmpInput []string
	tmpInput = make([]string, 5)
	//牌力數值越小牌型越大，故取名min，而非最小牌型
	var minCardType = 10       //最小的牌型
	var minCardStrength = 7462 //最小的牌力

	for _, item := range t7c5 {
		tmpInput[0] = input[item[0]]
		tmpInput[1] = input[item[1]]
		tmpInput[2] = input[item[2]]
		tmpInput[3] = input[item[3]]
		tmpInput[4] = input[item[4]]

		cardType, cardStrength := Evaluator(tmpInput)
		if cardStrength < minCardStrength {
			minCardStrength = cardStrength
			minCardType = cardType
		}
	}
	return minCardType, minCardStrength
}
