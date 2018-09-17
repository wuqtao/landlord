package games

import (
	doudizhu2 "chessSever/program/logic/game/games/doudizhu"
	"chessSever/program/logic/game/poker"
)
//游戏使用接口类型，便于实现多态
type Game interface {
	GetPlayerNum() int                      //获取游戏玩家数量
	GetPokerCards() []poker.PokerCard       //获取游戏的扑克牌
	GetGameName() string                    //获取游戏名称
	GetGameID() int                         //获取游戏id
	GetDeckNum() int                        //获取游戏牌的付数
	GetLastCards() []poker.PokerCard    	//获取上一次出牌
	GetLastPlayer() int                     //获取上一次出牌的玩家索引
	GetNextPlayer() int 					//获取下一把出牌的玩家索引
	ShuffleCards()                          //洗牌
	DealCards()                             //给在座的玩家发牌
	Hint()	[]poker.PokerCard				//提示有效的出牌
	CompareCards(cardsNow []poker.PokerDeck,lastCards []poker.PokerCard) bool 	//比较牌型大小
	DisCard(cards []poker.PokerCard)        //出牌
	Rules(cards []poker.PokerCard) bool     //出牌规则
}

func GetGame(name string) Game {
	switch name {
	case "doudizhu":
		return doudizhu2.GetDoudizhu()
	case "shengji":
		return nil
	case "baohuang":
		return nil
	case "zhajinhua":
		return nil
	default:
		return nil
	}
}
