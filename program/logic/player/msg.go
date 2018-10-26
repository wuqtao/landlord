package player

import (
	"chessSever/program/logic/game/poker"
	"encoding/json"
	"chessSever/program/data"
)

const(
	MSG_TYPE_OF_READY                  = iota //准备
	MSG_TYPE_OF_UN_READY                      //取消准备
	MSG_TYPE_OF_JOIN_TABLE                    //加入桌子
	MSG_TYPE_OF_LEAVE_TABLE                   //离开桌子

	MSG_TYPE_OF_HINT       //提示
	MSG_TYPE_OF_PLAY_CARD  //出牌
	MSG_TYPE_OF_PASS       //过牌

	MSG_TYPE_OF_AUTO                 //托管
	MSG_TYPE_OF_SEND_CARD            //发牌
	MSG_TYPE_OF_CALL_SCORE           //抢地主叫分
	MSG_TYPE_OF_CONFIRM              //客户端出牌等操作确认信息
	MSG_TYPE_OF_CALL_SCORE_TIME_OUT  //叫地主超时
	MSG_TYPE_OF_PLAY_ERROR           //出牌错误
	TYPE_OF_PLAY_CARD_SUCCESS        //出牌成功
	TYPE_OF_TABLE_BRODCAST           //桌子广播消息
	MSG_TYPE_OF_SCORE_CHANGE         //牌局分数变化
	MSG_TYPE_OF_SETTLE_SCORE         //结算玩家分数
	MSG_TYPE_OF_GAME_OVER            //游戏结束
	MSG_TYPE_OF_LOGIN                //登陆消息
)
type SendCard struct {
	Index int          //标志当前牌在用户所有牌中的索引位置
	Card poker.PokerCard
}
//发送给客户端的消息类型
type SendCardMsg struct{
	MsgType int
	Cards []*SendCard
}

type Msg struct {
	MsgType int
	Msg string
}


func newSendCardMsg(cards []*poker.PokerCard) ([]byte,error){
	cardMsg := SendCardMsg{
		MSG_TYPE_OF_SEND_CARD,
		[]*SendCard{},
	}
	for i,card := range cards{
		sendCard := SendCard{}
		sendCard.Index =i
		sendCard.Card = *card
		cardMsg.Cards = append(cardMsg.Cards,&sendCard)
	}
	return json.Marshal(cardMsg)
}

func newCallScoreMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_CALL_SCORE,
		"",
	}
	return json.Marshal(msg)
}

func newCallScoreTimeOutMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_CALL_SCORE_TIME_OUT,
		"",
	}
	return json.Marshal(msg)
}

func newPlayCardMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_PLAY_CARD,
		"",
	}
	return json.Marshal(msg)
}

func newPlayCardsErrorMsg(error string) ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_PLAY_ERROR,
		error,
	}
	return json.Marshal(msg)
}

func newPlayCardSuccessMsg() ([]byte,error){
	msg := Msg{
		TYPE_OF_PLAY_CARD_SUCCESS,
		"",
	}
	return json.Marshal(msg)
}

type LoginMsg struct{
	MsgType int
	Msg string
	*data.User
}

func NewLoginMsg(loginMsg string) LoginMsg{
	return LoginMsg{
		MSG_TYPE_OF_LOGIN,
		loginMsg,
		&data.User{},
	}
}

type BroadCastMsg struct{
	MsgType          int
	SubMsgType       int
	Msg              string
	Cards            []*poker.PokerCard
	Score            int
	PlayerId         int
	SettleInfoDic    map[int]string
	PlayerIndexIdDic map[int]int
}
func newBraodCastMsg() BroadCastMsg{
	msg := BroadCastMsg{
		TYPE_OF_PLAY_CARD_SUCCESS,
		-1,
		"",
		[]*poker.PokerCard{},
		-1,
		-1,
		make(map[int]string),
		make(map[int]int),
		}
	return msg
}
/*
	确认消息
	{
		"msgType":MSG_TYPE_OF_CONFIRM
		"data"{

		}
	}
	发牌消息形式
	{
		"msgType":MSG_TYPE_OF_SEND_CARD,
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
		"msgType":MSG_TYPE_OF_SEND_CARD,
		"data":{
			pokerCars[{
				"carIndex":,    //当前玩家手中的index
			},]
		}
	}


 */
