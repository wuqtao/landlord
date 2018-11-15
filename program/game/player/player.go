package player

import (
	"github.com/gorilla/websocket"
	"sync"
	"strconv"
	"github.com/tidwall/gjson"
	"fmt"
	"chessSever/program/model"
	"chessSever/program/game/msg"
	"chessSever/program/game/games"
	"chessSever/program/game"
	"github.com/sirupsen/logrus"
	"time"
	"chessSever/program/game/poker/set"
	"chessSever/program/game/poker/recorder"
	"chessSever/program/game/poker/analyzer"
)

/**
定义游戏玩家对象
*/
type Player struct {
	User *model.User
	Conn  *websocket.Conn 				//用户socket链接
	sync.RWMutex
	PokerCards       set.PokerSet 		//玩家手里的扑克牌0

	Index            int                //在桌子上的索引
	IsReady          bool               //是否准备
	IsAuto           bool               //是否托管
	PlayedCardIndexs []int              //已经出牌的ID
	callScoreChan    chan int           //叫地主通道
	playCardsChan    chan []int 		//出牌的索引切片通道
	stopTimeChan	 chan byte			//停止倒计时的通道

	IsOnline bool                       //是否在线，用于断线重连
	UpLineTime time.Time
	OffLine time.Time
	PokerRecorder recorder.PokerRecorder
	PokerAnalyzer analyzer.PokerAnalyzer

	UseablePokerSets []set.PokerSet
	CurrHintSetIndex int
}

func NewPlayer(user *model.User,conn *websocket.Conn) *Player {
	player := Player{
		User:user,
		Conn:conn,
		PlayedCardIndexs:[]int{},
		callScoreChan:make(chan int),
		playCardsChan:make(chan []int),
		stopTimeChan:make(chan byte),
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

func (p *Player) GetPlayerCards(indexs []int) set.PokerSet{
	if indexs != nil && len(indexs) > 0{
		temCards := set.PokerSet{}
		for _,i := range indexs{
			temCards = append(temCards,p.PokerCards[i])
		}
		return temCards
	}else{
		return p.PokerCards
	}
}

func (p *Player) SetPokerCards(cards set.PokerSet){

	p.Lock()
	p.PokerCards = cards
	p.Unlock()
	logrus.Debug("发牌给玩家"+strconv.Itoa(p.GetPlayerUser().Id),cards)
	msg,err := msg.NewSendCardMsg(cards)
	if err == nil{
		p.SendMsg(msg)
	}else{
		fmt.Println(err.Error())
	}
}

func (p *Player) StartCallScore(){
	currMsg,err := msg.NewCallScoreMsg()
	if err == nil{
		p.SendMsg(currMsg)

		go func(){
			score := <-p.callScoreChan
			game,err := game.GetPlayerGame(p)
			if err == nil{
				game.PlayerCallScore(p,score)
			}else{
				currMsg,err1 := msg.NewPlayCardsErrorMsg(err.Error())
				if err1 == nil{
					p.SendMsg(currMsg)
				}
				fmt.Println(err.Error())
			}
		}()
		//启动定时器,限制叫地主时间，过时自动不叫
		go func(){
			//给玩家发送定时消息
			game,err := game.GetPlayerGame(p)
			if err == nil{
				second := 7
				for {
					select {
						case <-p.stopTimeChan:
							fmt.Println("用户叫地主，定时器退出")
							return
						default:
							game.BroadCastMsg(p,msg.MSG_TYPE_OF_TIME_TICKER,strconv.Itoa(second))
							second--
							if second <= 0{
								p.callScoreChan<-0
								return
							}
							time.Sleep(time.Second)
					}
				}
			}else{
				fmt.Println("未获得用户game")
			}
		}()
	}else{
		fmt.Println(err.Error())
	}
}

func (p *Player) StartPlay(){
	currMsg,err := msg.NewPlayCardMsg()
	if err == nil{
		p.Lock()
		currGame,err := game.GetPlayerGame(p)
		if err == nil{
			 lastCards := currGame.GetLastCard()
			 //如果上家没有出牌或者上次是当前玩家出牌，提示可用最小牌即可,否则根据上轮出牌给出可用的扑克牌
			 if lastCards == nil || lastCards.PlayerIndex == p.Index{
				 p.UseablePokerSets = []set.PokerSet{p.PokerAnalyzer.GetMinPlayableCards()}
			 }else{
			 	p.UseablePokerSets = p.PokerAnalyzer.GetUseableCards(lastCards.PokerSetTypeInfo)
			 }
			 //重新分析完牌型，将当前提示的牌型索引重置为0
			 p.CurrHintSetIndex = 0
		}
		//如果有牌可以出，发送出牌消息，否则发送要不起消息
		if len(p.UseablePokerSets) > 0 && p.UseablePokerSets[0].GetLength() > 0{
			p.SendMsg(currMsg)
		}else{
			//todo
		}

		p.Unlock()
		go func(){
			cardIndexs := <-p.playCardsChan
			if len(cardIndexs) == 0{
				currGame.PlayerPassCard(p)
			}else{
				currGame.PlayerPlayCards(p,cardIndexs)
			}
		}()
		//启动定时器,限制叫地主时间，过时自动不叫
		go func(){
			//给玩家发送定时消息
			second := 3
			for {
				select {
					case <-p.stopTimeChan:
						fmt.Println("玩家出牌，定时器结束")
						return
					default:
						currGame.BroadCastMsg(p,msg.MSG_TYPE_OF_TIME_TICKER,strconv.Itoa(second))
						second--
						if second <= 0{
							p.autoPlay(currGame)
							return
						}
						time.Sleep(time.Second)
				}
			}
		}()
	}else{
		fmt.Println(err.Error())
	}
}

func (p *Player)autoPlay(currGame game.IGame){
	if len(p.UseablePokerSets) > 0 {
		indexs,err := p.PokerCards.GetPokerIndexs(p.UseablePokerSets[0])
		if err == nil{
			p.playCardsChan<- indexs
		}else{
			fmt.Println(err.Error())
		}
	}else{
		p.playCardsChan<- []int{}
	}
}

func fiterCardIndex(cardIndexs []int,playedCardIndexs []int) []int{
	//检测待出牌切片中牌是否已经出过
	for j,index := range cardIndexs {
		for _,playedIndex := range playedCardIndexs{
			if index == playedIndex{
				cardIndexs[j] = -1
				break
			}
		}
	}
	//重新整理待出的牌
	playIndexs := []int{}
	for _,index := range cardIndexs {
		if index != -1{
			playIndexs = append(playIndexs,index)
		}
	}
	return playIndexs
}

func (p *Player) CallScore(score int){
	p.stopTimeChan<-1
	p.callScoreChan<-score
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

	fmt.Println("玩家出牌"+strconv.Itoa(p.GetPlayerUser().Id))

	p.stopTimeChan<-1
	p.playCardsChan<-cardIndexs
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

func (p *Player) SetPokerRecorder(recorder recorder.PokerRecorder){
	p.PokerRecorder = recorder
}
func (p *Player) SetPokerAnalyzer(analyzer analyzer.PokerAnalyzer){
	p.PokerAnalyzer = analyzer
}