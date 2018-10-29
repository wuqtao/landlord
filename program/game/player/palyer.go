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
)

/**
定义游戏玩家对象
*/
type Player struct {
	User *model.User
	Conn     *websocket.Conn //用户socket链接
	Table    *Table          //桌子
	sync.RWMutex
	PokerCards       []*poker.PokerCard //玩家手里的扑克牌0

	Index            int                //在桌子上的索引
	IsReady          bool               //是否准备
	IsAuto           bool               //是否托管
	CallScore        int                //用户叫地主分数
	CalledScore      bool               //用户是否已经叫地主
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
func (p *Player) JoinTableByKey(key string){
	err := GetRoom().getTable(key).addPlayer(p)
	if err != nil{
		fmt.Println(err.Error())
	}
}

func (p *Player) JoinTable(table *Table){
	err := table.addPlayer(p)
	if err != nil{
		fmt.Println(err.Error())
	}
	p.Lock()
	p.Table = table
	p.Unlock()
}
//开牌桌
func (p *Player) CreateTable(gameID int){
	table := newTable(p, gameID)
	p.Lock()
	p.Table = table
	p.Unlock()
}

func (p *Player) LeaveTable() {
	p.Table.removePlayer(p)
	p.Lock()
	p.Table = nil
	p.Unlock()
}
//用户跟该桌所有人说话
func (p *Player) SayToTable(msg []byte){
	p.Table.RLock()
	for _,po := range p.Table.Players{
		if po != p {
			po.Conn.WriteMessage(websocket.TextMessage,msg)
		}
	}
	p.Table.RUnlock()
}
//用户跟该桌某一个说话
func (p *Player) SayToAnother(id int,msgB []byte){
	p.Table.RLock()
	for _,po := range p.Table.Players{
		if po.User.Id == id {
			po.Conn.WriteMessage(websocket.TextMessage,msgB)
		}
	}
	p.Table.RUnlock()
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
			p.unReady()
		case msg.MSG_TYPE_OF_READY:
			p.Ready()
		case msg.MSG_TYPE_OF_PLAY_CARD:
			cardIndex := gjson.Get(string(msgB),"Data.CardIndex").Array()
			cards := []int{}
			for _,card := range cardIndex{
				cards = append(cards,int(card.Int()))
			}
			p.playCards(cards)
		case msg.MSG_TYPE_OF_PASS:
			p.pass()
		case msg.MSG_TYPE_OF_LEAVE_TABLE:

		case msg.MSG_TYPE_OF_JOIN_TABLE:

		case msg.MSG_TYPE_OF_HINT:

		case msg.MSG_TYPE_OF_CALL_SCORE:
			score,_ := strconv.Atoi(gjson.Get(string(msgB),"Data.Score").String())
			p.callScore(score)

		default:
			p.Conn.WriteMessage(msgType,msgB)
	}

	return nil
}

func (p *Player)Ready(){
	p.Lock()
	if(p.Table != nil){
		p.IsReady = true
		fmt.Println(strconv.Itoa(p.User.Id)+"is ready")
		p.Unlock()
		p.Table.userReady()
		p.Table.BroadCastMsg(p,msg.MSG_TYPE_OF_READY,"玩家准备")
	}else{
		p.Unlock()
	}
}

func (p *Player)unReady(){
	p.Lock()
	p.IsReady = false
	p.Unlock()
	p.Table.BroadCastMsg(p,msg.MSG_TYPE_OF_UN_READY,"玩家取消准备")
}

func (p *Player)callScore(score int){
	p.Lock()
	p.CallScore = score
	p.CalledScore = true
	p.Unlock()
	p.Table.userCallScore(p,score)
}
//出牌
func (p *Player)playCards(cards []int){
	p.Table.userPlayCard(p,cards)
}
//过牌
func (p *Player)pass(){
	p.Table.userPassCard(p)
}
//出牌成功
func (p *Player)playCardSuccess(){
	SendMsgToPlayer(p,msg.MSG_TYPE_OF_PLAY_CARD_SUCCESS,"用户出牌成功")
}

//出牌出错
func (p *Player)playCardError(error string){
	SendMsgToPlayer(p,msg.MSG_TYPE_OF_PLAY_ERROR,error)
}
//提示出牌
func(p *Player) hintCards(){

}

