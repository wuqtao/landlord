package player

import (
	"chessSever/program/logic/game/poker"
	"encoding/json"
	"fmt"
)

const(
	TypeOfReady = iota    //准备
	TypeOfUnReady		  //取消准备
	TypeOfJoinTable	      //加入桌子
	TypeOfLeaveTable      //离开桌子

	TypeOfHint      //提示
	TypeOfPlayCard  //出牌
	TypeOfPass      //过牌

	TypeOfAuto			  //托管
	TypeOfSendCard		  //发牌
	TypeOfCallScore       //抢地主叫分
	TypeOfConfirm         //客户端出牌等操作确认信息
)
type SendCard struct {
	index int          //标志当前牌在用户所有牌中的索引位置
	card poker.PokerCard
}
//发送给客户端的消息类型
type CardMsg struct{
	msgType int
	cards []SendCard
}

func newSendCardMsg(cards []*poker.PokerCard) ([]byte,error){
	cardMsg := CardMsg{
		TypeOfSendCard,
		[]SendCard{},
	}
	for i,card := range cards{
		sendCard := SendCard{}
		sendCard.index =i
		sendCard.card = *card
		cardMsg.cards = append(cardMsg.cards,sendCard)
	}
	fmt.Println(cardMsg)
	json,err := json.Marshal(cardMsg)
	fmt.Println(string(json))
	return json,err
}
/*
	确认消息
	{
		"msgType":TypeOfConfirm
		"data"{

		}
	}
	发牌消息形式
	{
		"msgType":TypeOfSendCard,
		"data":{
			pokerCars[{
				"carIndex":,    //当前玩家手中的index
				"carNum":,      //牌的数字
				"carColor":,    //牌的花色
				""
			},]
		}
	}

	出牌消息
	{
		"msgType":TypeOfSendCard,
		"data":{
			pokerCars[{
				"carIndex":,    //当前玩家手中的index
			},]
		}
	}


 */
