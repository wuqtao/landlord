package player

import (
	"fmt"
	"landlord/program/game/msg"
	"log"
)

func SendPlayerCards(curPlayer *Player){
	json,err := msg.NewSendCardMsg(curPlayer.PokerCards)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	curPlayer.SendMsg(json)
}

func SendMsgToPlayer(p *Player,msgType int,hints string){

	var newMsg []byte
	var err error
	switch msgType {
		case msg.MSG_TYPE_OF_CALL_SCORE:
			newMsg,err = msg.NewCallScoreMsg()
		case msg.MSG_TYPE_OF_CALL_SCORE_TIME_OUT:
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
		p.SendMsg(newMsg)
	}else{
		log.Fatal(err.Error())
	}
}