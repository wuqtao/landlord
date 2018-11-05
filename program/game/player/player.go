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
	"chessSever/program/game/games"
	"chessSever/program/game"
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

func (p *Player) GetPlayerUser() *model.User{
	return p.User
}

func (p *Player) GetIndex() int{
	return p.Index
}

func (p *Player) SetIndex(index int){
	p.Lock()
	p.Index = index
	p.Unlock()
}

func (p *Player) GetReadyStatus() bool{
	return p.IsReady
}

func (p *Player) GetAutoStatus() bool{
	return p.IsAuto
}

func (p *Player) GetPlayedCardIndexs() []int{
	return p.PlayedCardIndexs
}

func (p *Player) GetPlayerCards(indexs []int) []*poker.PokerCard{
	if len(indexs) > 0{
		temCards := []*poker.PokerCard{}
		for _,i := range indexs{
			temCards = append(temCards,p.PokerCards[i])
		}
		return temCards
	}else{
		return p.PokerCards
	}
}

func (p *Player) SetPokerCards(cards []*poker.PokerCard){
	p.Lock()
	p.PokerCards = cards
	p.Unlock()
	msg,err := msg.NewSendCardMsg(cards)
	if err == nil{
		p.SendMsg(msg)
	}else{
		fmt.Println(err.Error())
	}
}

func (p *Player) StartCallScore(){
	msg,err := msg.NewCallScoreMsg()
	if err == nil{
		p.SendMsg(msg)
	}else{
		fmt.Println(err.Error())
	}
}

func (p *Player) StartPlay(){
	msg,err := msg.NewPlayCardMsg()
	if err == nil{
		p.SendMsg(msg)
	}else{
		fmt.Println(err.Error())
	}
}

//按照桌号加入牌桌
func (p *Player) JoinGame(gameType int,gameId int){
	game,err := game.GetRoom().GetGame(gameType,gameId)
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
func (p *Player) CreateGame(gameType int,baseScore int){
	 err := games.NewGame(gameType,baseScore).AddPlayer(p)
	 if err != nil{
		println(err.Error())
	 }
}

func (p *Player) LeaveGame() {
	game,err := game.GetPlayerGame(p)
	if err == nil {
		err := game.RemovePlayer(p)
		if err != nil{
			println(err.Error())
		}
	}else{
		println(err.Error())
	}
}
//用户跟该桌所有人说话
func (p *Player) SayToOthers(msg []byte){

	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.SayToOthers(p,msg)
	}else{
		//todo
	}
}
//用户跟该桌某一个说话
func (p *Player) SayToAnother(id int,msg []byte){
	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.SayToAnother(p,id,msg)
	}else{
		//todo
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
			go p.UnReady()
		case msg.MSG_TYPE_OF_READY:
			go p.Ready()
		case msg.MSG_TYPE_OF_PLAY_CARD:
			cardIndex := gjson.Get(string(msgB),"Data.CardIndex").Array()
			cards := []int{}
			for _,card := range cardIndex{
				cards = append(cards,int(card.Int()))
			}
			go p.PlayCards(cards)
		case msg.MSG_TYPE_OF_PASS:
			go p.Pass()
		case msg.MSG_TYPE_OF_LEAVE_TABLE:

		case msg.MSG_TYPE_OF_JOIN_TABLE:

		case msg.MSG_TYPE_OF_HINT:

		case msg.MSG_TYPE_OF_CALL_SCORE:
			score,_ := strconv.Atoi(gjson.Get(string(msgB),"Data.Score").String())
			go p.CallScore(score)

		default:
			go p.Conn.WriteMessage(msgType,msgB)
	}

	return nil
}

func (p *Player)Ready(){
	p.Lock()
	p.IsReady = true
	p.Unlock()

	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.PlayerReady(p)
	}else{
		msg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
		if err1 == nil{
			p.SendMsg(msg)
		}
		fmt.Println(err.Error())
	}
}

func (p *Player) UnReady(){
	p.Lock()
	p.IsReady = false
	p.Unlock()

	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.PlayerUnReady(p)
	}else{
		msg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
		if err1 == nil{
			p.SendMsg(msg)
		}
		fmt.Println(err.Error())
	}
}

func (p *Player) CallScore(score int){
	game,err := game.GetPlayerGame(p)
	if err == nil{
		game.PlayerCallScore(p,score)
	}else{
		msg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
		if err1 == nil{
			p.SendMsg(msg)
		}
		fmt.Println(err.Error())
	}
}
//出牌
func (p *Player) PlayCards(cardIndexs []int){

	p.RLock()
	for _,index := range cardIndexs{
		//判断是否是之前出过的牌
		if p.PlayedCardIndexs != nil {
			for _,playedIndex := range p.PlayedCardIndexs{
				if index == playedIndex {
					p.PlayCardError("出牌中包含已出的牌")
					p.RUnlock()
					return
				}
			}
		}
	}
	p.RUnlock()

	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.PlayerPlayCards(p,cardIndexs)
	}else{
		msg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
		if err1 == nil{
			p.SendMsg(msg)
		}
		fmt.Println(err.Error())
	}
}
//过牌
func (p *Player)Pass(){
	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.PlayerPassCard(p)
	}else{
		msg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
		if err1 == nil{
			p.SendMsg(msg)
		}
		fmt.Println(err.Error())
	}
}
//出牌成功
func (p *Player) PlayCardSuccess(cardIndexs []int){

	if p.PlayedCardIndexs == nil{
		p.PlayedCardIndexs = []int{}
	}

	for _,index := range cardIndexs{
		p.PlayedCardIndexs = append(p.PlayedCardIndexs,index)
	}

	SendMsgToPlayer(p,msg.MSG_TYPE_OF_PLAY_CARD_SUCCESS,"用户出牌成功")
}

func (p *Player)IsOutOfCards() bool{
	return len(p.PlayedCardIndexs) == len(p.PokerCards)
}
//出牌出错
func (p *Player) PlayCardError(error string){
	SendMsgToPlayer(p,msg.MSG_TYPE_OF_PLAY_ERROR,error)
}
//提示出牌
func(p *Player) HintCards(){
	game,err := game.GetPlayerGame(p)
	if err == nil {
		game.HintCards(p)
	}else{
		msg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
		if err1 == nil{
			p.SendMsg(msg)
		}
		fmt.Println(err.Error())
	}
}

func (p *Player) SendMsg(msg []byte){
	p.Conn.WriteMessage(websocket.TextMessage,msg)
}

