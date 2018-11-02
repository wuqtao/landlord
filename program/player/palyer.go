package player

import (
	"github.com/gorilla/websocket"
	"sync"
	"strconv"
	"github.com/tidwall/gjson"
	"fmt"
	"chessSever/program/game/poker"
	"chessSever/program/model"
	"chessSever/program/game/msg"
	"chessSever/program/proxy"
	"chessSever/program/game/games"
)

/**
定义游戏玩家对象
*/
type Player struct {
	User *model.User
	Conn  *websocket.Conn //用户socket链接
	sync.RWMutex
	PokerCards       []*poker.PokerCard //玩家手里的扑克牌0

	Index            int                //在桌子上的索引
	IsReady          bool               //是否准备
	IsAuto           bool               //是否托管
	PlayedCardIndexs []int              //已经出牌的ID
}

func NewPlayer(user *model.User,conn *websocket.Conn) *Player {
	player := Player{
		User:user,
		Conn:conn,
	}
	return &player
}

//按照桌号加入牌桌
func (p *Player) JoinGame(gameType int,gameId int){
	game,err := proxy.GetGameRoom().GetGame(gameType,gameId)
	if err != nil{
		fmt.Println(err.Error())
	}else{
		err := game.AddPlayer(p)
		if err != nil{
			println(err.Error())
		}
	}
}
//开牌桌
func (p *Player) CreateGame(gameID int,baseScore int){
	 err := games.NewGame(gameID,baseScore).AddPlayer(p)
	 if err != nil{
		println(err.Error())
	 }
}

func (p *Player) LeaveGame() {
	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		err := game.RemovePlayer(p)
		if err != nil{
			println(err.Error())
		}
	}else{
		println("未加入任何桌子")
	}

}
//用户跟该桌所有人说话
func (p *Player) SayToOthers(msg []byte){

	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.SayToOthers(p,msg)
	}else{
		println("未加入桌子")
	}
}
//用户跟该桌某一个说话
func (p *Player) SayToAnother(id int,msg []byte){
	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.SayToAnother(p,id,msg)
	}else{
		println("未加入桌子")
	}
}

func (p *Player)ResolveMsg(msgB []byte) error{
	fmt.Println(string(msgB))
	msgType,err := strconv.Atoi(gjson.Get(string(msgB),"MsgType").String())
	if err != nil{
		p.Conn.WriteMessage(websocket.TextMessage,msgB)
		return err
	}

	switch msgType {
		case msg.MSG_TYPE_OF_AUTO:

		case msg.MSG_TYPE_OF_UN_READY:
			p.UnReady()
		case msg.MSG_TYPE_OF_READY:
			p.Ready()
		case msg.MSG_TYPE_OF_PLAY_CARD:
			cardIndex := gjson.Get(string(msgB),"Data.CardIndex").Array()
			cards := []int{}
			for _,card := range cardIndex{
				cards = append(cards,int(card.Int()))
			}
			p.PlayCards(cards)
		case msg.MSG_TYPE_OF_PASS:
			p.pass()
		case msg.MSG_TYPE_OF_LEAVE_TABLE:

		case msg.MSG_TYPE_OF_JOIN_TABLE:

		case msg.MSG_TYPE_OF_HINT:

		case msg.MSG_TYPE_OF_CALL_SCORE:
			score,_ := strconv.Atoi(gjson.Get(string(msgB),"Data.Score").String())
			p.CallScore(score)

		default:
			p.Conn.WriteMessage(msgType,msgB)
	}

	return nil
}

func (p *Player)Ready(){
	p.Lock()
	p.IsReady = true
	p.Unlock()

	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.PlayerReady(p)
	}else{
		println("未加入桌子")
	}
}

func (p *Player) UnReady(){
	p.Lock()
	p.IsReady = false
	p.Unlock()

	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.PlayerUnReady(p)
	}else{
		println("未加入桌子")
	}
}

func (p *Player) CallScore(score int){
	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.PlayerCallScore(p,score)
	}else{
		println("未加入桌子")
	}
}
//出牌
func (p *Player) PlayCards(cards []int){
	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.PlayerPlayCards(p,cards)
	}else{
		println("未加入桌子")
	}
}
//过牌
func (p *Player)pass(){
	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.PlayerPassCard(p)
	}else{
		println("未加入桌子")
	}
}
//出牌成功
func (p *Player) PlayCardSuccess(){
	SendMsgToPlayer(p,msg.MSG_TYPE_OF_PLAY_CARD_SUCCESS,"用户出牌成功")
}

//出牌出错
func (p *Player) PlayCardError(error string){
	SendMsgToPlayer(p,msg.MSG_TYPE_OF_PLAY_ERROR,error)
}
//提示出牌
func(p *Player) HintCards(){
	game,ok := proxy.GetPlayerGameDic().Dic[p]
	if ok {
		game.HintCards(p)
	}else{
		println("未加入桌子")
	}
}

