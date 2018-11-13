package msg

import (
	"encoding/json"
	"strconv"
	"chessSever/program/game/poker/card"
	"chessSever/program/game/poker/set"
)

const(
	MSG_TYPE_OF_READY    = iota //准备
	MSG_TYPE_OF_UN_READY           //取消准备
	MSG_TYPE_OF_JOIN_TABLE         //加入桌子
	MSG_TYPE_OF_LEAVE_TABLE        //离开桌子

	MSG_TYPE_OF_HINT       			//提示
	MSG_TYPE_OF_PLAY_CARD  			//出牌
	MSG_TYPE_OF_PASS       			//过牌

	MSG_TYPE_OF_AUTO                 //托管
	MSG_TYPE_OF_SEND_CARD            //发牌
	MSG_TYPE_OF_CALL_SCORE           //抢地主叫分
	MSG_TYPE_OF_CONFIRM              //客户端出牌等操作确认信息
	MSG_TYPE_OF_CALL_SCORE_TIME_OUT  //叫地主超时
	MSG_TYPE_OF_PLAY_ERROR           //出牌错误
	MSG_TYPE_OF_PLAY_CARD_SUCCESS    //出牌成功
	MSG_TYPE_OF_TABLE_BRODCAST       //桌子广播消息
	MSG_TYPE_OF_SCORE_CHANGE         //牌局分数变化
	MSG_TYPE_OF_SETTLE_SCORE         //结算玩家分数
	MSG_TYPE_OF_GAME_OVER            //游戏结束
	MSG_TYPE_OF_LOGIN                //登陆消息
	MSG_TYPE_OF_SEND_BOTTOM_CARDS    //发底牌
	MSG_TYPE_OF_TIME_TICKER          //倒计时数
)
type SendCard struct {
	Index int          //标志当前牌在用户所有牌中的索引位置
	Card card.PokerCard
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


func NewSendCardMsg(cards set.PokerSet) ([]byte,error){
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

func NewCallScoreMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_CALL_SCORE,
		"",
	}
	return json.Marshal(msg)
}
func NewTimeCountMsg(second int) ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_CALL_SCORE,
		strconv.Itoa(second),
	}
	return json.Marshal(msg)
}
func NewCallScoreTimeOutMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_CALL_SCORE_TIME_OUT,
		"",
	}
	return json.Marshal(msg)
}

func NewPlayCardMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_PLAY_CARD,
		"",
	}
	return json.Marshal(msg)
}

func NewPlayCardsErrorMsg(error string) ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_PLAY_ERROR,
		error,
	}
	return json.Marshal(msg)
}

func NewPlayCardSuccessMsg() ([]byte,error){
	msg := Msg{
		MSG_TYPE_OF_PLAY_CARD_SUCCESS,
		"",
	}
	return json.Marshal(msg)
}

type LoginMsg struct{
	MsgType int
	Msg string
	ID int
}

func NewLoginMsg(userID int,loginMsg string) ([]byte,error){
	newMsg :=LoginMsg{
		MSG_TYPE_OF_LOGIN,
		loginMsg,
		-1,
	}
	newMsg.ID = userID
	return json.Marshal(newMsg)
}

type BroadCastMsg struct{
	MsgType          int
	SubMsgType       int
	Msg              string
	Cards            set.PokerSet
	CardsIndex       []int
	Score            int
	PlayerId         int
	SettleInfoDic    map[string]string
	PlayerIndexIdDic map[string]int
}
func NewBraodCastMsg() BroadCastMsg{
	msg := BroadCastMsg{
		MSG_TYPE_OF_TABLE_BRODCAST,
		-1,
		"",
		set.PokerSet{},
		[]int{},
		-1,
		-1,
		make(map[string]string),
		make(map[string]int),
		}
	return msg
}
/*
	确认消息
	{
		"msgType":MSG_TYPE_OF_CONFIRM
		"model"{

		}
	}
	发牌消息形式
	{
		"msgType":MSG_TYPE_OF_SEND_CARD,
		"model":{
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
		"model":{
			pokerCars[{
				"carIndex":,    //当前玩家手中的index
			},]
		}
	}


 */
