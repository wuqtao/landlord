package poker

//定义玩家的扑克牌分析器map的索引为poker的value,value为改值得扑克牌在玩家牌中的索引
type PokerAnalyzer map[int]PokerSet
//创建新的扑克牌分析器
func NewPokerAnalyzer() PokerAnalyzer{
	analyzer := PokerAnalyzer{}
   	analyzer[POKER_VALUE_THREE] = PokerSet{}
	analyzer[POKER_VALUE_FOUR] = PokerSet{}
	analyzer[POKER_VALUE_FIVE] = PokerSet{}
	analyzer[POKER_VALUE_SIX] = PokerSet{}
	analyzer[POKER_VALUE_SEVEN] = PokerSet{}
	analyzer[POKER_VALUE_EIGHT] = PokerSet{}
	analyzer[POKER_VALUE_NINE] = PokerSet{}
	analyzer[POKER_VALUE_TEN] = PokerSet{}
	analyzer[POKER_VALUE_JACK] = PokerSet{}
	analyzer[POKER_VALUE_QUEEN] = PokerSet{}
	analyzer[POKER_VALUE_KING] = PokerSet{}
	analyzer[POKER_VALUE_ACE] = PokerSet{}
	analyzer[POKER_VALUE_TWO] = PokerSet{}
	analyzer[POKER_VALUE_BLACK_JOKER] = PokerSet{}
	analyzer[POKER_VALUE_RED_JOKER] = PokerSet{}
	return analyzer
}
//根据给定的扑克集初始化分析器
func (ana PokerAnalyzer) InitAnalyzer(pokers PokerSet){
	pokers.DoOnEachPokerPoker(func(index int,card *PokerCard){
		ana[card.CardValue].AddPokers(PokerSet{card})
	})
}
//根据给定的扑克集更新记牌器,出牌时调用
func (ana PokerAnalyzer) UpdateAnalyzer(pokers PokerSet){
	pokers.DoOnEachPokerPoker(func(index int,card *PokerCard){
		ana[card.CardValue],_ = ana[card.CardValue].DelPokers(PokerSet{card})
	})
}
