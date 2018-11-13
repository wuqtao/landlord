package recorder

import (
	"strconv"
	"chessSever/program/game/poker/card"
	"chessSever/program/game/poker/set"
)
//记牌器，记录玩家之外带各牌待出的数量
type PokerRecorder map[string]int

func NewPokerRecorder() PokerRecorder {
	recorder := make(map[string]int)
	return recorder
}
//初始化玩家记牌器，发牌后调用
func (re PokerRecorder)InitRecorder(){
	re[card.POKER_SYMBOL_THREE] = 0
	re[card.POKER_SYMBOL_FOUR] = 0
	re[card.POKER_SYMBOL_FIVE] = 0
	re[card.POKER_SYMBOL_SIX] = 0
	re[card.POKER_SYMBOL_SEVEN] = 0
	re[card.POKER_SYMBOL_EIGHT] = 0
	re[card.POKER_SYMBOL_NINE] = 0
	re[card.POKER_SYMBOL_TEN] = 0
	re[card.POKER_SYMBOL_JACK] = 0
	re[card.POKER_SYMBOL_QUEEN] = 0
	re[card.POKER_SYMBOL_KING] = 0
	re[card.POKER_SYMBOL_ACE] = 0
	re[card.POKER_SYMBOL_TWO] = 0
	re[card.POKER_SYMBOL_BLACK_JOKER] = 0
	re[card.POKER_SYMBOL_RED_JOKER] = 0
}
//增加记录器
func (re PokerRecorder)AddPokerSet(playerPokers ...set.PokerSet){
	for _,pokerSet := range playerPokers{
		pokerSet.DoOnEachPokerCard(func(index int,card *card.PokerCard){
			re[card.GetCardName()]++
		})
	}
}
//更新玩家记牌器,玩家出牌后调用
func (re PokerRecorder)RemovePokerSet(cards set.PokerSet){
	cards.DoOnEachPokerCard(func(index int,card *card.PokerCard){
		re[card.GetCardName()]--
	})
}
//根据牌的顺序从大到小排序的记牌器json对象
func (re PokerRecorder) SequenceJsonEncode() string{
	jsonString := ""
	jsonString += "{"
	jsonString += "'"+card.POKER_SYMBOL_RED_JOKER+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_RED_JOKER))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_BLACK_JOKER+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_BLACK_JOKER))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_TWO+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_TWO))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_ACE+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_ACE))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_KING+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_KING))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_QUEEN+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_QUEEN))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_JACK+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_JACK))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_TEN+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_TEN))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_NINE+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_NINE))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_EIGHT+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_EIGHT))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_SEVEN+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_SEVEN))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_SIX+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_SIX))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_FIVE+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_FIVE))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_FOUR+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_FOUR))
	jsonString += ","
	jsonString += "'"+card.POKER_SYMBOL_THREE+"'"+":"+strconv.Itoa(re.GetPokerNum(card.POKER_SYMBOL_THREE))
	jsonString += "}"
	return jsonString
}

func (re PokerRecorder) GetPokerNum(key string) int{
	num,ok := re[key]
	if ok {
		return num
	}else{
		return 0
	}
}


