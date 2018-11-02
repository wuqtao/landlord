package game

import (
	"chessSever/program/game/poker"
)

const (
	GAME_ID_OF_DOUDOZHU = iota
	GAME_ID_OF_SHENGJI
	GAME_ID_OF_BAOHUANG
	GAME_ID_OF_ZHAJINHUA
)

var gameIDNameDic map[int]string

func init(){
	gameIDNameDic = make(map[int]string)
	gameIDNameDic[GAME_ID_OF_DOUDOZHU] = "斗地主"
	gameIDNameDic[GAME_ID_OF_SHENGJI] = "升级"
	gameIDNameDic[GAME_ID_OF_BAOHUANG] = "保皇"
	gameIDNameDic[GAME_ID_OF_ZHAJINHUA] = "斗地主"
}

func GetGameName(gameID int) string{
	name,ok := gameIDNameDic[gameID]
	if ok{
		return name
	}else{
		return "未定义游戏名称"
	}
}

//游戏使用接口类型，便于实现多态
type IGame interface {
	GetPlayerNum() int                                                              //获取游戏玩家数量
	GetPokerCards() []*poker.PokerCard                                              //获取游戏的扑克牌
	GetGameName() string                                                            //获取游戏名称
	GetGameID() int                                                                 //获取游戏id
	GetDeckNum() int                                                                //获取游戏牌的付数
	GetLastCards() *LastCardsType                                                  //获取上一次出牌
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
	PlayerCardIndexs []int		 	//扑克牌在出牌玩家所有牌中的index
	PlayerIndex int              	//出牌的玩家index
	CardsType int				 	//牌的类型
	Cards []*poker.PokerCard	    //出的牌
	CardMinAndMax map[string]int    //出的牌的map,min为主体牌的最小值，max为有主体牌的最大值，比如三带一，三为主体牌，带一为辅助牌
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