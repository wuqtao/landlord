package analyzer

import (
	"chessSever/program/game/poker/set"
	"chessSever/program/game/poker/card"
)

//定义玩家的扑克牌分析器map的索引为poker的value,value为改值得扑克牌在玩家牌中的索引
type PokerAnalyzer map[int]set.PokerSet
//创建新的扑克牌分析器
func NewPokerAnalyzer() PokerAnalyzer{
	analyzer := PokerAnalyzer{}
	return analyzer
}
//根据给定的扑克集初始化分析器
func (ana PokerAnalyzer) InitAnalyzer(){
	ana[card.POKER_VALUE_THREE] = set.PokerSet{}
	ana[card.POKER_VALUE_FOUR] = set.PokerSet{}
	ana[card.POKER_VALUE_FIVE] = set.PokerSet{}
	ana[card.POKER_VALUE_SIX] = set.PokerSet{}
	ana[card.POKER_VALUE_SEVEN] = set.PokerSet{}
	ana[card.POKER_VALUE_EIGHT] = set.PokerSet{}
	ana[card.POKER_VALUE_NINE] = set.PokerSet{}
	ana[card.POKER_VALUE_TEN] = set.PokerSet{}
	ana[card.POKER_VALUE_JACK] = set.PokerSet{}
	ana[card.POKER_VALUE_QUEEN] = set.PokerSet{}
	ana[card.POKER_VALUE_KING] = set.PokerSet{}
	ana[card.POKER_VALUE_ACE] = set.PokerSet{}
	ana[card.POKER_VALUE_TWO] = set.PokerSet{}
	ana[card.POKER_VALUE_BLACK_JOKER] = set.PokerSet{}
	ana[card.POKER_VALUE_RED_JOKER] = set.PokerSet{}
}
//根据给定的扑克集更新记牌器,出牌时调用
func (ana PokerAnalyzer) RemovePokerSet(pokers set.PokerSet){
	pokers.DoOnEachPokerCard(func(index int,card *card.PokerCard){
		ana[card.CardValue],_ = ana[card.CardValue].DelPokers(set.PokerSet{card})
	})
}

func (ana PokerAnalyzer) AddPokerSet(pokers set.PokerSet){
	pokers.DoOnEachPokerCard(func(index int,card *card.PokerCard){
		ana[card.CardValue] = ana[card.CardValue].AddPokers(set.PokerSet{card})
	})
}
