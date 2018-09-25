package games

import (
	doudizhu2 "chessSever/program/logic/game/games/doudizhu"
	"chessSever/program/logic/game/poker"
	"chessSever/program/logic/player"
)

const (
	DouDiZhu="doudizhu"
	ShengJi = "shengji"
	BaoHuang ="baohuang"
	Zhajinhua = "zhajinhua"
)
//游戏使用接口类型，便于实现多态
type Game interface {
	GetPlayerNum() int                      //获取游戏玩家数量
	GetPokerCards() []*poker.PokerCard      //获取游戏的扑克牌
	GetGameName() string                    //获取游戏名称
	GetGameID() int                         //获取游戏id
	GetDeckNum() int                        //获取游戏牌的付数
	GetLastCards() []*poker.PokerCard    	//获取上一次出牌
	initCards()								//初始化游戏中的牌
	shuffleCards()                          //洗牌
	DealCards(table *player.Table)    		//给在座的玩家发牌
	Hint()	[]*poker.PokerCard				//提示有效的出牌
	CompareCards(cardsNow []poker.PokerDeck,lastCards []poker.PokerCard) bool 	//比较牌型大小
	IsMatchRoles() bool                     //是否符合出牌规则
}

func GetGame(name string) Game {
	switch name {
	case DouDiZhu:
		return doudizhu2.GetDoudizhu()
	case ShengJi:
		return nil
	case BaoHuang:
		return nil
	case Zhajinhua:
		return nil
	default:
		return nil
	}
}
