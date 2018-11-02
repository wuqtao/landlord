package doudizhu

import (
	"math/rand"
	"time"
	"chessSever/program/game"
	"chessSever/program/game/poker"
	"chessSever/program/game/player"
	"sync"
	"fmt"
	"chessSever/program/game/msg"
	"strconv"
	"errors"
	"encoding/json"
	"github.com/gorilla/websocket"
)

type Doudizhu struct {
	id int                               //游戏ID
	name string                          //游戏名称

	playerNum int						 //玩家数
	deckNum int                          //几副牌

	sync.RWMutex                         //操作playNum以及player时加锁

	IsPlaying  bool                      //是否正在游戏中
	pokerCards []*poker.PokerCard        //当前游戏中的所有的牌
	baseScore int						 //底分
	currMulti int                        //当前倍率
	CalledLoardNum  int                  //叫过地主的人数
	lastCards  *game.LastCardsType 		 //最后的出牌结构


	Players []*player.Player             //玩家数组
	lordIndex int                        //地主索引
	playerCards [][]*poker.PokerCard     //同桌不同玩家的牌的切片
	bottomCards []*poker.PokerCard       //底牌
	CurrPlayerIndex int                  //当前叫地主或者出牌人的index
	OutCardIndexs   []int                //出完牌的用户index
}

var originDoudizhu Doudizhu

func init(){
	//其他索引型属性不能在此赋值
	originDoudizhu = Doudizhu{
		name:game.GetGameName(game.GAME_ID_OF_DOUDOZHU),
		playerNum:3,
		deckNum:1,
		baseScore:10,
		currMulti:1,
	}
}

func GetDoudizhu(baseScore int,p *player.Player) game.IGame{

	newDou := originDoudizhu

	newDou.Lock()
	newDou.baseScore = baseScore
	newDou.pokerCards = []*poker.PokerCard{}
	newDou.Players = []*player.Player{}
	newDou.playerCards = [][]*poker.PokerCard{[]*poker.PokerCard{},[]*poker.PokerCard{},[]*poker.PokerCard{}}
	newDou.bottomCards = []*poker.PokerCard{}
	newDou.id = game.GetRoom().AddGame(game.GAME_ID_OF_DOUDOZHU,&newDou)
	newDou.Unlock()

	newDou.InitCards()

	newDou.addPlayer(p)
	return &newDou
}

//增加玩家
func (dou *Doudizhu) addPlayer(player *player.Player) error {

	if dou.IsPlaying{
		return errors.New("游戏进行中，无法加入游戏")
	}
	if len(dou.Players) > dou.playerNum{
		panic("player Num ")
		return errors.New("游戏数据出错")
	}
	if len(dou.Players) == dou.playerNum{
		dou.Lock()
		for i,p := range dou.Players{
			if p == nil{
				dou.Players[i] = player
				fmt.Println("有新玩家加入桌子"+strconv.Itoa(dou.id))
				dou.Unlock()
				dou.BroadCastMsg(player,msg.MSG_TYPE_OF_JOIN_TABLE,"玩家加入桌子")
				return nil
			}else{
				if(i == len(dou.Players)){
					dou.Unlock()
					return errors.New("该桌玩家已满")
				}
			}
		}
	}else{
		dou.Players = append(dou.Players,player)
		fmt.Println("有新玩家加入桌子"+strconv.Itoa(dou.id))
		dou.Unlock()
		dou.BroadCastMsg(player,msg.MSG_TYPE_OF_JOIN_TABLE,"玩家加入桌子")
		return nil
	}
	return nil
}
//移除玩家
func (dou *Doudizhu) removePlayer(player *player.Player) error{
	if dou.IsPlaying{
		return errors.New("游戏进行中，无法移除玩家")
	}
	dou.Lock()
	for i, p := range dou.Players {
		if p == player {
			dou.Players[i] = nil
			break
		}
	}
	dou.Unlock()
	dou.BroadCastMsg(player,msg.MSG_TYPE_OF_LEAVE_TABLE,"玩家离开桌子")
	fmt.Println("桌子"+strconv.Itoa(dou.id)+"移除玩家"+strconv.Itoa(player.User.Id))
	return nil
}

func (dou *Doudizhu) userReady(){

	userAllReady := false
	for _,p := range dou.Players{
		if p != nil{
			p.RLock()
			if p.IsReady{
				userAllReady = true
			}else{
				userAllReady = false
			}
			p.RUnlock()
		}
	}
	//用户都准备好了，则发牌
	if userAllReady {
		dou.Lock()
		fmt.Println("桌子"+strconv.Itoa(dou.id)+"的玩家都准备好了")
		dou.IsPlaying = true
		dou.Unlock()
		dou.DealCards()
	}
}

func (dou *Doudizhu) dealCards(){
	fmt.Println("开始发牌")
	dou.RLock()
	for i,currPlayer := range dou.Players{
		currPlayer.PokerCards = dou.playerCards[i]
		player.SendPlayerCards(currPlayer)
	}
	dou.RUnlock()
	dou.callLoard()
}

func (dou *Doudizhu) callLoard(){
	rand.Seed(time.Now().Unix())
	currUserIndex := rand.Int31n(int32(dou.playerNum-1))
	dou.nextCallLoard(int(currUserIndex))
}

func (dou *Doudizhu) userCallScore(currPlayer *player.Player,score int){
	dou.Lock()
	dou.CalledLoardNum++
	var i int
	var p *player.Player
	for i,p = range dou.Players{
		if p == currPlayer {
			break
		}
	}
	//直到第一个人二次抢地主结束
	if dou.CalledLoardNum == dou.playerNum+1 {

		if score != 0 {
			dou.lordIndex = i
		}
		dou.Unlock()
		dou.callLoardEnd()
	}else{
		if score != 0 {
			dou.lordIndex = i
		}
		dou.Unlock()
		dou.nextCallLoard(-1)
	}
	dou.BroadCastMsg(currPlayer,msg.MSG_TYPE_OF_CALL_SCORE,"用户叫地主")
}

func (dou *Doudizhu) callLoardEnd(){
	dou.Lock()
	dou.CurrPlayerIndex = 0
	dou.CalledLoardNum = 0
	dou.Unlock()
	fmt.Println("叫地主结束"+strconv.Itoa(dou.lordIndex)+"成为地主")
	currPlayer := dou.Players[dou.lordIndex]

	for _,card := range dou.bottomCards{
		currPlayer.PokerCards = append(currPlayer.PokerCards,card)
	}
	poker.CommonSort(dou.Players[dou.lordIndex].PokerCards)
	player.SendPlayerCards(dou.Players[dou.lordIndex])

	dou.BroadCastMsg(dou.Players[dou.lordIndex],msg.MSG_TYPE_OF_SEND_BOTTOM_CARDS,"发放底牌")
	fmt.Println("底牌发送完毕，开始游戏")
	dou.play(nil)
}

func (dou *Doudizhu) nextCallLoard(index int){

	var currPlayer *player.Player
	if index >= 0{
		dou.Lock()
		dou.CurrPlayerIndex = index
		currPlayer = dou.Players[dou.CurrPlayerIndex]
		dou.Unlock()
	}else{
		currPlayer = dou.GetNextLoard()
	}

	player.SendMsgToPlayer(currPlayer,msg.MSG_TYPE_OF_CALL_SCORE,"叫地主")
	//限制叫地主时间
	go func(){
		time.Sleep(time.Second*10)
		currPlayer.RLock()
		if !currPlayer.CalledScore{
			currPlayer.RUnlock()
			player.SendMsgToPlayer(currPlayer,msg.MSG_TYPE_OF_CALL_SCORE_TIME_OUT,"叫地主")
			currPlayer.CallScore(0)
		}else{
			currPlayer.RUnlock()
		}
	}()
}

func (dou *Doudizhu) play(currPlayer *player.Player){
	if currPlayer == nil{
		dou.CurrPlayerIndex = dou.lordIndex
		currPlayer = dou.Players[dou.lordIndex]
	}
	player.SendMsgToPlayer(currPlayer,msg.MSG_TYPE_OF_PLAY_CARD,"玩家出牌")
}

func (dou *Doudizhu) userPlayCard(p *player.Player,cardIndexs []int){
	//符合出牌规则才允许出牌
	if dou.GetCurrPlayerIndex(p) != dou.CurrPlayerIndex{
		p.PlayCardError("还没到您的出牌次序")
		return
	}

	cards := []*poker.PokerCard{}
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
		cards = append(cards,p.PokerCards[index])
	}
	p.RUnlock()

	lastCards,err := dou.MatchRoles(dou.GetCurrPlayerIndex(p),cards)
	if err == nil{

		if  dou.lastCards == nil || lastCards.PlayerIndex == dou.lastCards.PlayerIndex ||
			(lastCards.CardsType == dou.lastCards.CardsType &&
				lastCards.CardMinAndMax["min"] > dou.lastCards.CardMinAndMax["min"] &&
				lastCards.CardMinAndMax["max"] > dou.lastCards.CardMinAndMax["min"]){

			if lastCards.PlayerCardIndexs == nil{
				lastCards.PlayerCardIndexs = []int{}
			}

			if p.PlayedCardIndexs == nil{
				p.PlayedCardIndexs = []int{}
			}

			for _,index := range cardIndexs{
				lastCards.PlayerCardIndexs = append(lastCards.PlayerCardIndexs,index)
				p.PlayedCardIndexs = append(p.PlayedCardIndexs,index)
			}

			isBomb := false
			dou.Lock()
			dou.lastCards = lastCards
			if dou.lastCards.CardsType == POKERS_TYPE_COMMON_BOMB ||
				dou.lastCards.CardsType == POKERS_TYPE_JOKER_BOMB{
				isBomb = true
				dou.currMulti *= 2
			}
			dou.Unlock()
			if(isBomb){
				dou.BroadCastMsg(p,msg.MSG_TYPE_OF_SCORE_CHANGE,"分数加倍")
			}
			//出牌成功，给前端提示
			p.PlayCardSuccess()

			dou.BroadCastMsg(p,msg.MSG_TYPE_OF_PLAY_CARD,"玩家出牌")
			//玩家的牌全部出完了
			if len(p.PlayedCardIndexs) == len(p.PokerCards) {
				if dou.OutCardIndexs == nil{
					dou.OutCardIndexs = []int{}
				}

				currIndex := dou.GetCurrPlayerIndex(p)
				dou.OutCardIndexs = append(dou.OutCardIndexs,currIndex)

				if currIndex == dou.lordIndex{
					dou.gameOver()
					return
				}else{
					if len(dou.OutCardIndexs) == 2{
						dou.gameOver()
						return
					}
				}
			}
			//下一个玩家出牌
			dou.play(dou.GetNextPlayer())

		}else{
			p.PlayCardError("出牌必须大于上一家")
		}

	}else{
		p.PlayCardError(err.Error())
	}
}

func (dou *Doudizhu) gameOver(){
	if len(dou.OutCardIndexs) == 1 {
		dou.BroadCastMsg(nil,msg.MSG_TYPE_OF_GAME_OVER,"游戏结束,地主胜利")
	}else{
		if dou.OutCardIndexs[1] == dou.lordIndex{
			dou.BroadCastMsg(nil,msg.MSG_TYPE_OF_GAME_OVER,"游戏结束,地主胜利")
		}else{
			dou.BroadCastMsg(nil,msg.MSG_TYPE_OF_GAME_OVER,"游戏结束,农民胜利")
		}
	}
}

func (dou *Doudizhu) userPassCard(currPlayer *player.Player){
	//之前出牌是当前玩家则不能过牌，第一个出牌玩家也不能过牌
	if dou.lastCards != nil && dou.GetCurrPlayerIndex(currPlayer) != dou.lastCards.PlayerIndex{
		currPlayer.PlayCardSuccess()
		dou.BroadCastMsg(currPlayer,msg.MSG_TYPE_OF_PASS,"用户过牌")
		dou.play(dou.GetNextPlayer())
	}else{
		currPlayer.PlayCardError("第一个出牌的玩家不能过牌")
	}
}

func (dou *Doudizhu) GetNextPlayer() *player.Player{
	dou.Lock()
	defer dou.Unlock()
	if(dou.CurrPlayerIndex >= dou.playerNum-1){
		dou.CurrPlayerIndex = 0
	}else{
		dou.CurrPlayerIndex++
	}

	return dou.Players[dou.CurrPlayerIndex]
}

func (dou *Doudizhu) GetNextLoard() *player.Player{
	dou.Lock()
	defer dou.Unlock()
	if(dou.CurrPlayerIndex >= dou.playerNum-1){
		dou.CurrPlayerIndex = 0
	}else{
		dou.CurrPlayerIndex++
	}

	return dou.Players[dou.CurrPlayerIndex]
}

func (dou *Doudizhu) GetCurrPlayerIndex(currPlayer *player.Player) int {
	dou.RLock()
	defer dou.RUnlock()
	for i,p := range dou.Players{
		if(p == currPlayer){
			return i
		}
	}
	return -1
}

func (dou *Doudizhu) BroadCastMsg(player *player.Player,msgType int,hints string){

	newMsg := msg.NewBraodCastMsg()
	newMsg.SubMsgType = msgType

	dou.RLock()
	defer dou.RUnlock()

	if player != nil{
		newMsg.PlayerId = player.User.Id
		for i,p := range dou.Players{
			if p != nil{
				newMsg.PlayerIndexIdDic["id"+strconv.Itoa(p.User.Id)] = i
			}
		}
	}

	switch msgType{
	case msg.MSG_TYPE_OF_READY:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"已准备"
	case msg.MSG_TYPE_OF_UN_READY:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"取消准备"
	case msg.MSG_TYPE_OF_JOIN_TABLE:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"加入游戏"
	case msg.MSG_TYPE_OF_LEAVE_TABLE:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"离开游戏"
	case msg.MSG_TYPE_OF_PLAY_CARD:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"出牌"
		for _,card := range dou.lastCards.Cards{
			newMsg.Cards = append(newMsg.Cards,card)
		}
	case msg.MSG_TYPE_OF_PASS:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"过牌"
	case msg.MSG_TYPE_OF_CALL_SCORE:
		newMsg.Msg = strconv.Itoa(player.User.Id)+"叫地主"
		newMsg.Score = dou.currMulti
	case msg.MSG_TYPE_OF_SCORE_CHANGE:
		newMsg.Msg = "基础变动"
		newMsg.Score = dou.currMulti
	case msg.MSG_TYPE_OF_SEND_BOTTOM_CARDS:
		newMsg.Msg = "发放底牌"
		newMsg.Cards = dou.bottomCards
		newMsg.PlayerId = player.User.Id
	case msg.MSG_TYPE_OF_GAME_OVER:
		newMsg.Msg = "游戏结束，结算积分"
		newMsg.Score = dou.currMulti
		for _,index := range dou.OutCardIndexs{
			if index == dou.lordIndex{
				newMsg.SettleInfoDic["id"+strconv.Itoa(dou.Players[index].User.Id)] = "+"+strconv.Itoa(dou.currMulti*dou.baseScore*2)
			}else{
				newMsg.SettleInfoDic["id"+strconv.Itoa(dou.Players[index].User.Id)] = "+"+strconv.Itoa(dou.currMulti*dou.baseScore)
			}
		}

		for i,player := range dou.Players{
			_,ok := newMsg.SettleInfoDic["id"+strconv.Itoa(player.User.Id)]
			if !ok{
				if i == dou.lordIndex{
					newMsg.SettleInfoDic["id"+strconv.Itoa(dou.Players[i].User.Id)] = "-"+strconv.Itoa(dou.currMulti*dou.baseScore*2)
				}else{
					newMsg.SettleInfoDic["id"+strconv.Itoa(dou.Players[i].User.Id)] = "-"+strconv.Itoa(dou.currMulti*dou.baseScore)
				}
			}
		}
	}

	msgJson,err := json.Marshal(newMsg)
	if err != nil {
		panic(err.Error())
	}

	for _,player := range dou.Players{
		if player != nil{
			player.Conn.WriteMessage(websocket.TextMessage,msgJson)
		}
	}
}

func (dou *Doudizhu)GetPlayerNum() int{
	return dou.playerNum
}

func (dou *Doudizhu)GetPokerCards() []*poker.PokerCard{
	return dou.pokerCards
}

func (dou *Doudizhu)GetGameName() string{
	return dou.name
}

func (dou *Doudizhu)GetGameID() int{
	return dou.id
}

func (dou *Doudizhu)GetDeckNum() int{
	return dou.deckNum
}

func (dou Doudizhu)GetLastCards() *game.LastCardsType{
	return dou.lastCards
}

//初始化游戏中的牌
func (dou *Doudizhu)InitCards(){
	dou.Lock()
	defer dou.Unlock()

	for i:=0;i<dou.deckNum;i++{
		deck := poker.CreateDeck()
		for i,_ := range deck.Cards{
			dou.pokerCards = append(dou.pokerCards,&deck.Cards[i])
		}
	}
}

//洗牌
func (dou *Doudizhu)ShuffleCards(){
	dou.Lock()
	defer dou.Unlock()

	rand.Seed(time.Now().Unix())
	for i := len(dou.pokerCards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		dou.pokerCards[i], dou.pokerCards[num] = dou.pokerCards[num], dou.pokerCards[i]
	}
}

//发牌
func (dou *Doudizhu)DealCards(){

	dou.ShuffleCards()

	dou.Lock()
	dou.playerCards[0] = dou.pokerCards[:17]
	dou.playerCards[1] = dou.pokerCards[17:34]
	dou.playerCards[2] = dou.pokerCards[34:51]
	dou.bottomCards = dou.pokerCards[51:]
	dou.Unlock()

	dou.sortPlayerCards()
}

func (dou *Doudizhu)Hint() []*poker.PokerCard{
	return []*poker.PokerCard{}
}

func (dou *Doudizhu)CompareCards(cardsNow []poker.PokerDeck,lastCards []poker.PokerCard) bool{
	return false
}
//检查出牌是否符合规则
func (dou *Doudizhu) MatchRoles(currPlayerIndex int,pokers []*poker.PokerCard) (*game.LastCardsType,error){
	return CheckRules(currPlayerIndex,pokers)
}
//获取玩家的牌
func (dou *Doudizhu)GetPlayerCards(index int) []*poker.PokerCard{
	return dou.playerCards[index]
}
//对玩家手中扑克牌，按照从小到大排序
func (dou *Doudizhu)sortPlayerCards(){
	dou.Lock()
	defer dou.Unlock()
	for _,cards := range dou.playerCards{
		poker.CommonSort(cards)
	}
}
//获取底牌
func (dou *Doudizhu) GetBottomCards() []*poker.PokerCard{
	return dou.bottomCards
}

