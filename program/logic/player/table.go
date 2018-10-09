package player

import (
	"chessSever/program/logic/game/games"
	"sync"
	"chessSever/program/logic/game/poker"
	"chessSever/program/logic/game"
	"strconv"
	"fmt"
	"errors"
	"time"
	"math/rand"
	"github.com/gorilla/websocket"
	"log"
)

/*
	定义游戏桌对象
*/
type Table struct {
	Key          string     				//桌子key,用于从room索引中查找桌子
	Players      []*Player        			//玩家数组
	Game         games.IGame 				//该桌玩的游戏
	sync.RWMutex            				//操作playNum以及player时加锁
	CurrPokerCards []*poker.PokerCard  		//当前出的牌
	CurrPalyerIndex int 					//当前出牌的玩家数组index
	IsPlaying    bool                       //是否正在游戏中
	CurrLoardIndex int                     //当前叫地主或者地主的Index
}
//创建桌子
func newTable(player *Player, gameName string) *Table {

	currGame := game.GetGame(gameName)
	table := Table{
		Game: currGame,
		Key:  "table" + strconv.Itoa(time.Now().Nanosecond()),//桌子的key要保证唯一且好找，所以用时间戳，
		Players:make([]*Player,currGame.GetPlayerNum()),
		IsPlaying:false,
	}
	fmt.Println("创建新桌子"+"table" + strconv.Itoa(player.Id))
	//桌子加入房间
	table.joinRoom()
	//将创建者加入桌子
	table.addPlayer(player)
	return &table
}
//加入房间
func (t *Table) joinRoom() {
	GetRoom().addTable(t.Key, t)
}
//销毁桌子
func (t *Table) destory() {
	t.Lock()
	if len(t.Players) >= 0 {
		for _, p := range t.Players {
			p.LeaveTable()
		}
	}
	GetRoom().removeTable(t.Key)
	fmt.Println("桌子"+t.Key+"销毁")
	t.Unlock()
}
//增加玩家
func (t *Table) addPlayer(player *Player) error {
	t.Lock()
	defer t.Unlock()
	if(len(t.Players) >= t.Game.GetPlayerNum()){
		for i,p := range t.Players{
			if p == nil{
				t.Players[i] = player
				player.Lock()
				player.Table = t
				player.Unlock()
				fmt.Println(t.Key+"有新玩家加入")
				return nil
			}else{
				if(i == len(t.Players)){
					return errors.New("该桌玩家已满")
				}
			}
		}
		return errors.New("该桌玩家已满")
	}else{
		t.Players = append(t.Players,player)
		player.Lock()
		player.Table = t
		player.Unlock()
		fmt.Println(t.Key+"有新玩家加入")
		return nil
	}
}
//移除玩家
func (t *Table) removePlayer(player *Player) {
	t.Lock()
	for i, p := range t.Players {
		if p == player {
			t.Players[i] = nil
			break
		}
	}
	fmt.Println("桌子"+t.Key+"移除玩家"+strconv.Itoa(player.Id))
	t.Unlock()
}

func (t *Table) userReady(){
	t.Lock()
	userAllReady := false
	for _,p := range t.Players{
		if p != nil && p.IsReady{
			userAllReady = true
		}else{
			userAllReady = false
		}
	}
	//用户都准备好了，则发牌
	if userAllReady {
		fmt.Println(t.Key+"的玩家都准备好了")
		t.IsPlaying = true
		t.Game.DealCards()
		t.dealCards()
	}
	t.Unlock()
}

func (t *Table) dealCards(){
	fmt.Println("开始发牌")
	for i,player := range t.Players{
		player.PokerCards = t.Game.GetPlayerCards(i)
		sendPlayerCards(player)
	}
	t.callLoard()
}

func (t *Table) callLoard(){
	rand.Seed(time.Now().Unix())
	currUserIndex := rand.Int31n(int32(t.Game.GetPlayerNum()-1))
	t.CurrLoardIndex = int(currUserIndex)

	callScoreMsg,err := newCallScoreMsg()
	if err == nil{
		t.Players[currUserIndex].Conn.WriteMessage(websocket.TextMessage,callScoreMsg)
		fmt.Println(strconv.Itoa(t.Players[currUserIndex].Id)+"开始叫地主")
	}else{
		log.Fatal(err.Error())
	}
}

func (t *Table) callLoardEnd(player *Player){
	t.Lock()
	defer t.Unlock()
	for i,p := range t.Players{
		if p == player{
			t.CurrLoardIndex = i
		}
	}
	fmt.Println("叫地主结束"+strconv.Itoa(t.CurrLoardIndex)+"成为地主")
	currPlayer := t.Players[t.CurrLoardIndex]
	for _,card := range t.Game.GetBottomCards(){
		currPlayer.PokerCards = append(currPlayer.PokerCards,card)
	}
	sendPlayerCards(player)
	fmt.Println("底牌发送完毕，开始游戏")
	t.play()
}

func (t *Table) nextCallLoard(){
	player := t.GetNextLoard()
	callScoreMsg,err := newCallScoreMsg()
	if err == nil{
		player.Conn.WriteMessage(websocket.TextMessage,callScoreMsg)
		fmt.Println(strconv.Itoa(player.Id)+"开始叫地主")
	}else{
		log.Fatal(err.Error())
	}
}

func (t *Table) play(){

}

func (t *Table) GetNextPlayer() *Player{
	t.Lock()
	defer t.Unlock()
	if(t.CurrPalyerIndex >= t.Game.GetPlayerNum()){
		t.CurrPalyerIndex = 0
	}else{
		t.CurrPalyerIndex++
	}

	return t.Players[t.CurrPalyerIndex]
}

func (t *Table) GetNextLoard() *Player{
	t.Lock()
	defer t.Unlock()
	if(t.CurrLoardIndex >= t.Game.GetPlayerNum()){
		t.CurrLoardIndex = 0
	}else{
		t.CurrLoardIndex++
	}

	return t.Players[t.CurrLoardIndex]
}






