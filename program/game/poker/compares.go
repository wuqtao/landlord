package poker

//常规比较，按照值大小比较
func CardCommonCompare(card1 *PokerCard,card2 *PokerCard) bool{
	if card1.CardValue > card2.CardValue{
		return true
	}else{
		return false
	}
}


