package set

import (
	"chessSever/program/game/poker/card"
)

//定义一副扑克牌
type PokerDeck struct {
	Cards [54]card.PokerCard
}

var originDeck PokerDeck

func init(){
	originDeck = CreateOriginDeck()
}

//创建原始扑克牌，后续只需要复制即可，不用再运算获得
func CreateOriginDeck() PokerDeck{
	deck := PokerDeck{}
	for i := 0;i<52;i++ {
		shang := i/4;
		yu := i%4;
		suit := ""
		switch(yu){
		case 0:
			suit = card.PokerSuitClub
		case 1:
			suit = card.PokerSuitDiamond
		case 2:
			suit = card.PokerSuitHeart
		case 3:
			suit = card.PokerSuitSpade
		}
		pokerValue := 0
		pokerName := ""
		switch(shang){
		case 0:
			pokerValue = card.POKER_VALUE_THREE
			pokerName = card.POKER_SYMBOL_THREE
		case 1:
			pokerValue = card.POKER_VALUE_FOUR
			pokerName = card.POKER_SYMBOL_FOUR
		case 2:
			pokerValue = card.POKER_VALUE_FIVE
			pokerName = card.POKER_SYMBOL_FIVE
		case 3:
			pokerValue = card.POKER_VALUE_SIX
			pokerName = card.POKER_SYMBOL_SIX
		case 4:
			pokerValue = card.POKER_VALUE_SEVEN
			pokerName = card.POKER_SYMBOL_SEVEN
		case 5:
			pokerValue = card.POKER_VALUE_EIGHT
			pokerName = card.POKER_SYMBOL_EIGHT
		case 6:
			pokerValue = card.POKER_VALUE_NINE
			pokerName = card.POKER_SYMBOL_NINE
		case 7:
			pokerValue = card.POKER_VALUE_TEN
			pokerName = card.POKER_SYMBOL_TEN
		case 8:
			pokerValue = card.POKER_VALUE_JACK
			pokerName = card.POKER_SYMBOL_JACK
		case 9:
			pokerValue = card.POKER_VALUE_QUEEN
			pokerName = card.POKER_SYMBOL_QUEEN
		case 10:
			pokerValue = card.POKER_VALUE_KING
			pokerName = card.POKER_SYMBOL_KING
		case 11:
			pokerValue = card.POKER_VALUE_ACE
			pokerName = card.POKER_SYMBOL_ACE
		case 12:
			pokerValue = card.POKER_VALUE_TWO
			pokerName = card.POKER_SYMBOL_TWO
		}
		deck.Cards[i] = card.PokerCard{
			pokerValue,
			suit,
			pokerName,
		}
	}
	deck.Cards[52] = card.PokerCard{
		card.POKER_VALUE_BLACK_JOKER,
		card.PokerSuitJoker,
		card.POKER_SYMBOL_BLACK_JOKER,
	}

	deck.Cards[53] = card.PokerCard{
		card.POKER_VALUE_RED_JOKER,
		card.PokerSuitJoker,
		card.POKER_SYMBOL_RED_JOKER,
	}
	return deck
}

//每个游戏桌子都有单独的扑克牌，防止洗牌等冲突
func CreateDeck() PokerDeck{
	copyDeck := originDeck
	return copyDeck
}
