package util

import "chessSever/program/logic/game/poker"

func BubbleSortCards(cards []*poker.PokerCard,maxCard func (card1 *poker.PokerCard,card2 *poker.PokerCard) bool){
	length := len(cards)
	for i :=0;i<length;i++{
		for j := i;j < length;j++{
			if maxCard(cards[i],cards[j]){
				cards[i],cards[j] = cards[j],cards[i]
			}
		}
	}
}


