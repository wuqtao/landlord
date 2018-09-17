package games

import (
	doudizhu2 "chessSever/program/logic/game/games/doudizhu"
	"chessSever/program/logic/game/poker"
)
//游戏使用接口类型，便于实现多态
type Game interface {
	GetPlayerNum() int                      //获取游戏玩家数量
	GetPokerCards() []poker.PokerCard      //获取游戏的扑克牌
	GetGameName() string                    //获取游戏名称
	GetGameID() int                         //获取游戏id
	GetDeckNum() int                        //获取游戏牌的付数
	ShuffleCards()                          //洗牌
	DealCards()                             //发牌
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





