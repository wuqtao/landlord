package poker

//从小到大对扑克牌排序
func CommonSort(pokers []*PokerCard){
	BubbleSortCardsMax2Min(pokers,CardCommonCompare)
}
//使用冒泡排序法，对给定的扑克牌，使用给定的规则进项从小到大排序
func BubbleSortCardsMax2Min(cards []*PokerCard,maxCard func (card1 *PokerCard,card2 *PokerCard) bool){
	length := len(cards)
	for i :=0;i<length;i++{
		for j := i;j < length;j++{
			if !maxCard(cards[i],cards[j]){
				cards[i],cards[j] = cards[j],cards[i]
			}
		}
	}
}

//使用冒泡排序法，对给定的扑克牌，使用给定的规则进项从小到大排序
func BubbleSortCardsMin2Max(cards []*PokerCard,maxCard func (card1 *PokerCard,card2 *PokerCard) bool){
	length := len(cards)
	for i :=0;i<length;i++{
		for j := i;j < length;j++{
			if maxCard(cards[i],cards[j]){
				cards[i],cards[j] = cards[j],cards[i]
			}
		}
	}
}

