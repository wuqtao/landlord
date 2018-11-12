package poker

/**
	定义扑克牌花色、显示牌型、值以及扑克牌
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






