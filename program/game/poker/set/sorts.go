package set

import "chessSever/program/game/poker/card"

//从小到大对扑克牌排序
func SortAsc(pokers PokerSet){
	BubbleSortCardsMin2Max(pokers,card.IsFirstValueBigger)
}
//从大到小对扑克牌排序
func SortDesc(pokers PokerSet){
	BubbleSortCardsMax2Min(pokers,card.IsFirstValueBigger)
}
//使用冒泡排序法，对给定的扑克牌，使用给定的规则进项从小到大排序
func BubbleSortCardsMax2Min(cards PokerSet,maxCard func (card1 *card.PokerCard,card2 *card.PokerCard) bool){
	length := cards.GetLength()
	for i :=0;i<length;i++{
		for j := i;j < length;j++{
			if !maxCard(cards[i],cards[j]){
				cards[i],cards[j] = cards[j],cards[i]
			}
		}
	}
}

//使用冒泡排序法，对给定的扑克牌，使用给定的规则进项从小到大排序
func BubbleSortCardsMin2Max(cards PokerSet,maxCard func (card1 *card.PokerCard,card2 *card.PokerCard) bool){
	length := cards.GetLength()
	for i :=0;i<length;i++{
		for j := i;j < length;j++{
			if maxCard(cards[i],cards[j]){
				cards[i],cards[j] = cards[j],cards[i]
			}
		}
	}
}

