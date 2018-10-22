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
	Key                string            //桌子key,用于从room索引中查找桌子
	Players         []*Player            //玩家数组
	Game            games.IGame          //该桌玩的游戏
	sync.RWMutex                         //操作playNum以及player时加锁
	CurrPokerCards  []*poker.PokerCard   //当前出的牌
	CurrPalyerIndex int                  //当前出牌的玩家数组index
	IsPlaying       bool                 //是否正在游戏中
	LoardIndex      int                  //当前地主的Index
	CurrLoardScore  int                  //当前地主分数
	CalledLoardNum  int                  //叫过地主的人数
	CurrPlayerIndex int                  //当前叫地主或者出牌人的index
	LastCards       *games.LastCardsType //最后的出牌结构

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
		t.LoardIndex = i
		t.Unlock()
		t.callLoardEnd()
	}else{
		if t.CalledLoardNum >= t.Game.GetPlayerNum() {
			t.Unlock()
			t.callLoardEnd()
		}else{
			if score > t.CurrLoardScore{
				t.CurrLoardScore = score
				t.LoardIndex = i
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
	t.CurrPlayerIndex = 0
	t.CalledLoardNum = 0
	t.Unlock()
	fmt.Println("叫地主结束"+strconv.Itoa(t.LoardIndex)+"成为地主")
	currPlayer := t.Players[t.LoardIndex]
	for _,card := range t.Game.GetBottomCards(){
		currPlayer.PokerCards = append(currPlayer.PokerCards,card)
	}
	util.BubbleSortCards(t.Players[t.LoardIndex].PokerCards,poker.CardCommonCompare)
	sendPlayerCards(t.Players[t.LoardIndex])
	fmt.Println("底牌发送完毕，开始游戏")
	t.play(nil)
}

func (t *Table) nextCallLoard(index int){

	var player *Player
	if index >= 0{
		t.Lock()
		t.CurrPlayerIndex = index
		player = t.Players[t.CurrPlayerIndex]
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

func (t *Table) play(player *Player){
	msg ,err := newPlayCardMsg()
	if err != nil{
		panic(err.Error())
	}
	if player == nil{
		t.CurrPalyerIndex = t.LoardIndex
		t.Players[t.LoardIndex].Conn.WriteMessage(websocket.TextMessage,msg)
	}else{
		player.Conn.WriteMessage(websocket.TextMessage,msg)
	}
}

func (t *Table) userPlayCard(p *Player,cardIndexs []int){
	//符合出牌规则才允许出牌
	cards := []*poker.PokerCard{}
	p.RLock()
	for _,index := range cardIndexs{
		cards = append(cards,p.PokerCards[index])
	}
	p.RUnlock()

	lastCards,err := t.Game.MatchRoles(t.GetCurrPlayerIndex(p),cards)
	if err == nil{

		if (lastCards.PlayerIndex == t.LastCards.PlayerIndex) ||
			(lastCards.CardsType == t.LastCards.CardsType &&
			lastCards.CardMinAndMax["min"] > t.LastCards.CardMinAndMax["min"] &&
			lastCards.CardMinAndMax["max"] > t.LastCards.CardMinAndMax["min"]){

				//出牌成功，给前段提示
				msg ,err := newPlayCardsErrorMsg("出牌必须大于上一家")
				if err != nil{
					panic(err.Error())
				}
				t.Players[t.LastCards.PlayerIndex].Conn.WriteMessage(websocket.TextMessage,msg)

				t.Lock()
				t.LastCards = lastCards
				t.Unlock()

				//下一个玩家出牌
				t.play(t.GetNextPlayer())
		}else{
			msg ,err := newPlayCardsErrorMsg("出牌必须大于上一家")
			if err != nil{
				panic(err.Error())
			}
			t.Players[lastCards.PlayerIndex].Conn.WriteMessage(websocket.TextMessage,msg)
		}
	}else{
		fmt.Println(err.Error())
		msg ,err := newPlayCardsErrorMsg(err.Error())
		if err != nil{
			panic(err.Error())
		}
		t.Players[lastCards.PlayerIndex].Conn.WriteMessage(websocket.TextMessage,msg)
		//用户出牌错误提示
	}
}

func (t *Table) GetNextPlayer() *Player{
	t.Lock()
	defer t.Unlock()
	if(t.CurrPlayerIndex >= t.Game.GetPlayerNum()-1){
		t.CurrPalyerIndex = 0
	}else{
		t.CurrPalyerIndex++
	}

	return t.Players[t.CurrPalyerIndex]
}

func (t *Table) GetNextLoard() *Player{
	t.Lock()
	defer t.Unlock()
	if(t.CurrPlayerIndex >= t.Game.GetPlayerNum()-1){
		t.CurrPlayerIndex = 0
	}else{
		t.CurrPlayerIndex++
	}

	return t.Players[t.CurrPlayerIndex]
}

func (t *Table) GetCurrPlayerIndex(player *Player) int {
	t.RLock()
	defer t.RUnlock()
	for i,p := range t.Players{
		if(p == player){
			return i
		}
	}
	return -1
}






