package poker

//定义玩家的扑克牌分析器map的索引为poker的value,value为改值得扑克牌在玩家牌中的索引
type PokerAnalyzer map[int][]int

func NewPokerAnalyzer() PokerAnalyzer{
	analyzer := PokerAnalyzer{}
   	analyzer[POKER_VALUE_THREE] = []int{}
	analyzer[POKER_VALUE_FOUR] = []int{}
	analyzer[POKER_VALUE_FIVE] = []int{}
	analyzer[POKER_VALUE_SIX] = []int{}
	analyzer[POKER_VALUE_SEVEN] = []int{}
	analyzer[POKER_VALUE_EIGHT] = []int{}
	analyzer[POKER_VALUE_NINE] = []int{}
	analyzer[POKER_VALUE_TEN] = []int{}
	analyzer[POKER_VALUE_JACK] = []int{}
	analyzer[POKER_VALUE_QUEEN] = []int{}
	analyzer[POKER_VALUE_KING] = []int{}
	analyzer[POKER_VALUE_ACE] = []int{}
	analyzer[POKER_VALUE_TWO] = []int{}
	analyzer[POKER_VALUE_BLACK_JOKER] = []int{}
	analyzer[POKER_VALUE_RED_JOKER] = []int{}
	return analyzer
}
