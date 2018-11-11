package poker

/**
	定义扑克牌花色和数字
*/
//定义扑克牌值
const (
	POKER_VALUE_THREE = iota
	POKER_VALUE_FOUR
	POKER_VALUE_FIVE
	POKER_VALUE_SIX
	POKER_VALUE_SEVEN
	POKER_VALUE_EIGHT
	POKER_VALUE_NINE
	POKER_VALUE_TEN
	POKER_VALUE_JACK
	POKER_VALUE_QUEEN
	POKER_VALUE_KING
	POKER_VALUE_ACE
	POKER_VALUE_TWO
	POKER_VALUE_BLACK_JOKER
	POKER_VALUE_RED_JOKER
)
//定义扑克牌符号
const (
	POKER_SYMBOL_THREE = "3"
	POKER_SYMBOL_FOUR = "4"
	POKER_SYMBOL_FIVE = "5"
	POKER_SYMBOL_SIX = "6"
	POKER_SYMBOL_SEVEN = "7"
	POKER_SYMBOL_EIGHT = "8"
	POKER_SYMBOL_NINE = "9"
	POKER_SYMBOL_TEN = "10"
	POKER_SYMBOL_JACK = "J"
	POKER_SYMBOL_QUEEN = "Q"
	POKER_SYMBOL_KING = "K"
	POKER_SYMBOL_ACE = "A"
	POKER_SYMBOL_TWO = "2"
	POKER_SYMBOL_BLACK_JOKER = "BlackJoker"
	POKER_SYMBOL_RED_JOKER = "RedJoker"
)
//定义扑克牌花色
const(
	PokerSuitDiamond = "Diamond" //方片
	PokerSuitHeart   = "Heart"   //红桃
	PokerSuitSpade   = "Spade"   //黑桃
	PokerSuitClub    = "Club"    //梅花
	PokerSuitJoker   = "Joker"   //大小王无花色
)
//定义扑克牌
type PokerCard struct {
	CardValue int			//card值用于排序比较
	CardSuit string			//card花色
	CardName string			//card显示的字符
}
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






