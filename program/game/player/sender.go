package player

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"chessSever/program/game/msg"
	"log"
)

func sendPlayerCards(curPlayer *Player){
	json,err := msg.NewSendCardMsg(curPlayer.PokerCards)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("给玩家"+strconv.Itoa(curPlayer.User.Id)+"发牌")
	curPlayer.Conn.WriteMessage(websocket.TextMessage,json)
}

func SendMsgToPlayer(p *Player,msgType int,hints string){

	var newMsg []byte
	var err error
	switch msgType {
		case msg.MSG_TYPE_OF_CALL_SCORE:
			fmt.Println(strconv.Itoa(p.User.Id)+"开始叫地主")
			newMsg,err = msg.NewCallScoreMsg()
		case msg.MSG_TYPE_OF_CALL_SCORE_TIME_OUT:
			fmt.Println(strconv.Itoa(p.User.Id)+"叫地主超时")
			newMsg,err = msg.NewCallScoreTimeOutMsg()
		case msg.MSG_TYPE_OF_PLAY_CARD:
			newMsg,err = msg.NewPlayCardMsg()
		case msg.MSG_TYPE_OF_PLAY_ERROR:
			newMsg,err = msg.NewPlayCardsErrorMsg(hints)
		case msg.MSG_TYPE_OF_PLAY_CARD_SUCCESS:
			newMsg,err = msg.NewPlayCardSuccessMsg()
		case msg.MSG_TYPE_OF_LOGIN:
			newMsg,err = msg.NewLoginMsg(p.User.Id,"登陆成功")
		default:
			return
	}

	if err == nil{
		p.Conn.WriteMessage(websocket.TextMessage,newMsg)
	}else{
		log.Fatal(err.Error())
	}
}