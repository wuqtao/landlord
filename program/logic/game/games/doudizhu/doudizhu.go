package doudizhu

import (
	"math/rand"
	"time"
	"chessSever/program/logic/game/poker"
	"chessSever/program/logic/game"
)

type Doudizhu struct {
	id int
	name string                           //游戏名称
	playerNum int						  //玩家数
	deckNum int                           //几副牌
	pokerCards []*poker.PokerCard         //当前游戏中的所有的牌
	lastCards []*poker.PokerCard          //最后一次出的牌
	currMulti int                         //当前倍率
	lordIndex int                         //地主索引
	lastThreeCards  []*poker.PokerCard    //最后三张底牌
	playerCards [][]*poker.PokerCard      //同桌不同玩家的牌的切片
	bottomCards []*poker.PokerCard        //底牌
}

func GetDoudizhu() game.IGame{
	dou := Doudizhu{
		id:1,
		name:"斗地主",
		playerNum:3,
		deckNum:1,
		pokerCards:[]*poker.PokerCard{},
		lastCards:[]*poker.PokerCard{},
	}
	dou.InitCards()
	return &dou
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

func (dou Doudizhu)GetLastCards() []*poker.PokerCard{
	return dou.lastCards
}

//初始化游戏中的牌
func (dou *Doudizhu)InitCards(){
	for i:=0;i<dou.deckNum;i++{
		deck := poker.CreateDeck()
		for _,card := range deck.Cards{
			dou.pokerCards = append(dou.pokerCards,&card)
		}
	}
}

//洗牌
func (dou *Doudizhu)ShuffleCards(){
	rand.Seed(time.Now().Unix())
	for i := len(dou.pokerCards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		dou.pokerCards[i], dou.pokerCards[num] = dou.pokerCards[num], dou.pokerCards[i]
	}
}

//发牌
func (dou *Doudizhu)DealCards(){

	dou.playerCards = make([][]*poker.PokerCard,dou.GetPlayerNum())
	for i:=0;i<len(dou.playerCards);i++{
		dou.playerCards[i] = []*poker.PokerCard{}
	}

	dou.ShuffleCards()
	dou.bottomCards = make([]*poker.PokerCard,3)
	dou.bottomCards[0] = dou.pokerCards[0]
	dou.bottomCards[1] = dou.pokerCards[1]
	dou.bottomCards[2] = dou.pokerCards[2]

	for i:=3; i<len(dou.pokerCards);i++  {
		yu := i%dou.playerNum
		dou.playerCards[yu] = append(dou.playerCards[yu],dou.pokerCards[i])
	}

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
	for _,cards := range dou.playerCards{
		poker.CommonSort(cards)
	}
}
//获取底牌
func (dou *Doudizhu) GetBottomCards() []*poker.PokerCard{
	return dou.bottomCards
}

