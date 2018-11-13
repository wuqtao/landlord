package card

//常规比较，按照值大小比较,如果第一个牌比第二个牌的值大，返回true
func IsFirstValueBigger(card1 *PokerCard,card2 *PokerCard) bool{
	if card1.CardValue > card2.CardValue{
		return true
	}else{
		return false
	}
}


