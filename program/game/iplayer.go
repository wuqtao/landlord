package game

import (
	"chessSever/program/model"
	"chessSever/program/game/poker"
)

type IPlayer interface {
	JoinGame(gameType int,gameId int)   							//加入游戏
	CreateGame(gameID int,baseScore int)  							//创建游戏
	LeaveGame()												 		//离开游戏
	SayToOthers(msg []byte)											//给同一游戏的其他玩家发消息
	SayToAnother(id int,msg []byte)                                 //给同意游戏的特定索引玩家发消息
	ResolveMsg(msgB []byte) error									//处理客户端消息
	Ready()														    //玩家准备
	UnReady()														//玩家取消准备
	CallScore(score int)											//抢地主
	PlayCards(cards []int)											//出牌
	Pass()	                                                        //过牌
	PlayCardSuccess(cardIndexs []int)												//出牌成功
	PlayCardError(err string)										//出牌错误
	HintCards()														//提示出牌

	GetPlayerUser() *model.User
	GetIndex() int
	GetReadyStatus() bool
	GetAutoStatus() bool
	GetPlayedCardIndexs() []int
	GetPlayerCards(indexs []int) []*poker.PokerCard

	SetPokerCards(cards []*poker.PokerCard)
	StartCallScore()
	StartPlay()
	IsOutOfCards() bool  											//是否出完牌
	SendMsg(msg []byte)
}
