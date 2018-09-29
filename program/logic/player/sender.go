package player

import (
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
)

func sendPlayerCards(curPlayer *Player){
	json,err := newSendCardMsg(curPlayer.PokerCards)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("给玩家"+strconv.Itoa(curPlayer.Id)+"发牌")
	curPlayer.Conn.WriteMessage(websocket.TextMessage,json)
	fmt.Println(json)
}
