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
	"chessSever/program/util"
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
	CurrLoardScore int 						//当前地主分数
	CalledLoardNum int                     //叫过地主的人数
	CurrCallLoardIndex int 					//当前叫地主index
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
		t.Unlock()
		t.Game.DealCards()
		t.dealCards()
	}else{
		t.Unlock()
	}

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
	t.nextCallLoard(int(currUserIndex))
}

func (t *Table) userCallScore(player *Player,score int){
	t.Lock()
	t.CalledLoardNum++
	var i int
	var p *Player
	for i,p = range t.Players{
		if p == player {
			break
		}
	}

	if score == 3{
		t.CurrLoardScore = score
		t.CurrLoardIndex = i
		t.Unlock()
		t.callLoardEnd()
	}else{
		if t.CalledLoardNum >= t.Game.GetPlayerNum() {
			t.Unlock()
			t.callLoardEnd()
		}else{
			if score > t.CurrLoardScore{
				t.CurrLoardScore = score
				t.CurrLoardIndex = i
				t.Unlock()
				t.nextCallLoard(-1)
			}else{
				t.Unlock()
				t.nextCallLoard(-1)
			}
		}
	}
}

func (t *Table) callLoardEnd(){
	t.Lock()
	t.CurrCallLoardIndex = 0
	t.CalledLoardNum = 0
	t.Unlock()
	fmt.Println("叫地主结束"+strconv.Itoa(t.CurrLoardIndex)+"成为地主")
	currPlayer := t.Players[t.CurrLoardIndex]
	for _,card := range t.Game.GetBottomCards(){
		currPlayer.PokerCards = append(currPlayer.PokerCards,card)
	}
	util.BubbleSortCards(t.Players[t.CurrLoardIndex].PokerCards,poker.CardCommonCompare)
	sendPlayerCards(t.Players[t.CurrLoardIndex])
	fmt.Println("底牌发送完毕，开始游戏")
	t.play()
}

func (t *Table) nextCallLoard(index int){

	var player *Player
	if index >= 0{
		t.Lock()
		t.CurrCallLoardIndex = index
		player = t.Players[t.CurrCallLoardIndex]
		t.Unlock()
	}else{
		player = t.GetNextLoard()
	}

	callScoreMsg,err := newCallScoreMsg()
	if err == nil{
		player.Conn.WriteMessage(websocket.TextMessage,callScoreMsg)
		fmt.Println(strconv.Itoa(player.Id)+"开始叫地主")
	}else{
		log.Fatal(err.Error())
	}
	//限制叫地主时间
	go func(player *Player) {
		time.Sleep(time.Second*10)
		player.RLock()
		if !player.CalledScore{
			player.RUnlock()
			callScoreTimeOutMsg,err := newCallScoreTimeOutMsg()
			if err == nil{
				player.Conn.WriteMessage(websocket.TextMessage,callScoreTimeOutMsg)
				fmt.Println(strconv.Itoa(player.Id)+"叫地主超时")
			}else{
				log.Fatal(err.Error())
			}
			player.callScore(0)
		}else{
			player.RUnlock()
		}
	}(player)
}

func (t *Table) play(){
	msg ,err := newPlayCardMsg()
	if err != nil{
		panic(err.Error())
	}
	t.Players[t.CurrLoardIndex].Conn.WriteMessage(websocket.TextMessage,msg)
}

func (t *Table) userPlayCard(p *Player,cardIndexs []int){
	//符合出牌规则才允许出牌
	cards := []*poker.PokerCard{}
	p.RLock()
	for _,index := range cardIndexs{
		cards = append(cards,p.PokerCards[index])
	}
	p.RUnlock()
	if t.Game.IsMatchRoles(cards){

	}else{

	}
}

func (t *Table) GetNextPlayer() *Player{
	t.Lock()
	defer t.Unlock()
	if(t.CurrCallLoardIndex >= t.Game.GetPlayerNum()-1){
		t.CurrPalyerIndex = 0
	}else{
		t.CurrPalyerIndex++
	}

	return t.Players[t.CurrPalyerIndex]
}

func (t *Table) GetNextLoard() *Player{
	t.Lock()
	defer t.Unlock()
	if(t.CurrCallLoardIndex >= t.Game.GetPlayerNum()-1){
		t.CurrCallLoardIndex = 0
	}else{
		t.CurrCallLoardIndex++
	}

	return t.Players[t.CurrCallLoardIndex]
}






