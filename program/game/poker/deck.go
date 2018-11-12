package poker

//定义一副扑克牌
type PokerDeck struct {
	Cards [54]PokerCard
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
			suit = PokerSuitClub
		case 1:
			suit = PokerSuitDiamond
		case 2:
			suit = PokerSuitHeart
		case 3:
			suit = PokerSuitSpade
		}
		pokerValue := 0
		pokerName := ""
		switch(shang){
		case 0:
			pokerValue = POKER_VALUE_THREE
			pokerName = POKER_SYMBOL_THREE
		case 1:
			pokerValue = POKER_VALUE_FOUR
			pokerName = POKER_SYMBOL_FOUR
		case 2:
			pokerValue = POKER_VALUE_FIVE
			pokerName = POKER_SYMBOL_FIVE
		case 3:
			pokerValue = POKER_VALUE_SIX
			pokerName = POKER_SYMBOL_SIX
		case 4:
			pokerValue = POKER_VALUE_SEVEN
			pokerName = POKER_SYMBOL_SEVEN
		case 5:
			pokerValue = POKER_VALUE_EIGHT
			pokerName = POKER_SYMBOL_EIGHT
		case 6:
			pokerValue = POKER_VALUE_NINE
			pokerName = POKER_SYMBOL_NINE
		case 7:
			pokerValue = POKER_VALUE_TEN
			pokerName = POKER_SYMBOL_TEN
		case 8:
			pokerValue = POKER_VALUE_JACK
			pokerName = POKER_SYMBOL_JACK
		case 9:
			pokerValue = POKER_VALUE_QUEEN
			pokerName = POKER_SYMBOL_QUEEN
		case 10:
			pokerValue = POKER_VALUE_KING
			pokerName = POKER_SYMBOL_KING
		case 11:
			pokerValue = POKER_VALUE_ACE
			pokerName = POKER_SYMBOL_ACE
		case 12:
			pokerValue = POKER_VALUE_TWO
			pokerName = POKER_SYMBOL_TWO
		}
		deck.Cards[i] = PokerCard{
			pokerValue,
			suit,
			pokerName,
		}
	}
	deck.Cards[52] = PokerCard{
		POKER_VALUE_BLACK_JOKER,
		PokerSuitJoker,
		POKER_SYMBOL_BLACK_JOKER,
	}

	deck.Cards[53] = PokerCard{
		POKER_VALUE_RED_JOKER,
		PokerSuitJoker,
		POKER_SYMBOL_RED_JOKER,
	}
	return deck
}

//每个游戏桌子都有单独的扑克牌，防止洗牌等冲突
func CreateDeck() PokerDeck{
	copyDeck := originDeck
	return copyDeck
}
