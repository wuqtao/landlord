package doudizhu

import (
	"chessSever/program/logic/game/games"
	"math/rand"
	"time"
	"chessSever/program/logic/game/poker"
	"chessSever/program/logic/player"
	"errors"
)

type Doudizhu struct {
	id int
	name string
	playerNum int
	deckNum int
	pokerCards []poker.PokerCard
	lastCards []poker.PokerCard
	lastPlayerIndex int
	table *player.Table
}

func GetDoudizhu(table *player.Table) games.Game{
	return Doudizhu{
		1,
		"斗地主",
		3,
		1,
		[]poker.PokerCard{},
		[]poker.PokerCard{},
		2,
		table,
	}
}

func (dou Doudizhu)GetPlayerNum() int{
	return dou.playerNum
}

func (dou Doudizhu)GetPokerCards() []poker.PokerCard{
	return dou.pokerCards
}

func (dou Doudizhu)GetGameName() string{
	return dou.name
}

func (dou Doudizhu)GetGameID() int{
	return dou.id
}

func (dou Doudizhu)GetDeckNum() int{
	return dou.deckNum
}
//洗牌
func (dou Doudizhu)ShuffleCards(){
	rand.Seed(time.Now().Unix())
	for i := len(dou.pokerCards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		dou.pokerCards[i], dou.pokerCards[num] = dou.pokerCards[num], dou.pokerCards[i]
	}
}
//发牌
func (dou Doudizhu)DealCards() error{

	if(len(dou.table.Players) != dou.playerNum){
		return errors.New("玩家数不满足游戏规则")
	}

	for i:=0;i<len(dou.pokerCards);i++{
		yu := i%dou.playerNum
		dou.table.Players[yu].PokerCards = append(dou.table.Players[yu].PokerCards,dou.pokerCards[i])
	}

	return nil
}

func (dou Doudizhu)GetLastCards() []poker.PokerCard{
	return dou.lastCards
}

func (dou Doudizhu)GetLastPlayer() int{
	return dou.lastPlayerIndex
}

func (dou Doudizhu)GetNextPlayer() int{
	return dou.playerNum-1-dou.lastPlayerIndex
}

func (dou Doudizhu)Hint() []poker.PokerCard{
	return []poker.PokerCard{}
}

func (dou Doudizhu)CompareCards(cardsNow []poker.PokerDeck,lastCards []poker.PokerCard) bool{
	return false
}

func (dou Doudizhu)DisCard(cards []poker.PokerCard){

}

func (dou Doudizhu)Rules(cards []poker.PokerCard) bool{
	return true
}