package poker

type PokerRecorder map[string]int

var originRecorder PokerRecorder

func init(){
	originRecorder = make(map[string]int)

	originRecorder[POKER_SYMBOL_THREE] = 0
	originRecorder[POKER_SYMBOL_FOUR] = 0
	originRecorder[POKER_SYMBOL_FIVE] = 0
	originRecorder[POKER_SYMBOL_SIX] = 0
	originRecorder[POKER_SYMBOL_SEVEN] = 0
	originRecorder[POKER_SYMBOL_EIGHT] = 0
	originRecorder[POKER_SYMBOL_NINE] = 0
	originRecorder[POKER_SYMBOL_TEN] = 0
	originRecorder[POKER_SYMBOL_JACK] = 0
	originRecorder[POKER_SYMBOL_QUEEN] = 0
	originRecorder[POKER_SYMBOL_KING] = 0
	originRecorder[POKER_SYMBOL_TWO] = 0
	originRecorder[POKER_SYMBOL_BLACK_JOKER] = 0
	originRecorder[POKER_SYMBOL_RED_JOKER] = 0
}

func newPokerRecorder() PokerRecorder {
	newRecorder := originRecorder
	return newRecorder
}
//初始化玩家记牌器，发牌后调用
func InitRecorder(playerPokers []map[int]*PokerCard) PokerRecorder{
	recorder := newPokerRecorder()
	for _,pokerCards := range playerPokers{
		for _,pokerCard := range pokerCards{
			recorder[pokerCard.CardName]++
		}
	}
	return recorder
}
//更新玩家记牌器,玩家出牌后调用
func UpdateRecorder(recorder PokerRecorder,cards []*PokerCard){
	for _,card := range cards{
		recorder[card.CardName]--
	}
}
