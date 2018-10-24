package games

import (
	"chessSever/program/logic/game/poker"
)

const (
	DouDiZhu="doudizhu"
	ShengJi = "shengji"
	BaoHuang ="baohuang"
	Zhajinhua = "zhajinhua"
)

//游戏使用接口类型，便于实现多态
type IGame interface {
	GetPlayerNum() int                                                              //获取游戏玩家数量
	GetPokerCards() []*poker.PokerCard                                              //获取游戏的扑克牌
	GetGameName() string                                                            //获取游戏名称
	GetGameID() int                                                                 //获取游戏id
	GetDeckNum() int                                                                //获取游戏牌的付数
	GetLastCards() []*poker.PokerCard                                               //获取上一次出牌
	InitCards()                                                                     //初始化游戏中的牌
	ShuffleCards()                                                                  //洗牌
	Hint()	[]*poker.PokerCard                                                    //提示有效的出牌
	CompareCards(cardsNow []poker.PokerDeck,lastCards []poker.PokerCard) bool       //比较牌型大小
	MatchRoles(currPlayerIndex int,pokers []*poker.PokerCard) (*LastCardsType,error) //是否符合出牌规则
	DealCards()                                                                     //发牌
	GetPlayerCards(index int) []*poker.PokerCard                                    //根据玩家索引获取玩家的牌
	GetBottomCards() []*poker.PokerCard                                             //获取底牌
}

type LastCardsType struct{
	PlayerCardIndexs []int		 	//扑克牌在出牌玩家中的index
	PlayerIndex int              	//出牌的玩家ID
	CardsType int				 	//牌的类型
	Cards []*poker.PokerCard	     //出的牌
	CardMinAndMax map[string]int		     //出的牌的map,min为主体牌的最小值，max为有主体牌的最大值，比如三第一，三为主体牌，带一为辅助牌
}

func NewLastCards(playerIndex int,cardsType int,cards []*poker.PokerCard,minValue int,maxValue int) *LastCardsType {
	lastCards :=  &LastCardsType{
		PlayerIndex:playerIndex,
		CardsType:cardsType,
		Cards:cards,
		CardMinAndMax:make(map[string]int),
	}
	lastCards.CardMinAndMax["min"] = minValue
	lastCards.CardMinAndMax["max"] = maxValue
	return lastCards
}