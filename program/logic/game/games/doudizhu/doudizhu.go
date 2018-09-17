package doudizhu

import (
	"chessSever/program/logic/game/games"
	"math/rand"
	"time"
	"chessSever/program/logic/game/poker"
)

type Doudizhu struct {
	id int
	name string
	playerNum int
	deckNum int
	pokerCards []poker.PokerCard
}

func GetDoudizhu() games.Game{
	return Doudizhu{
		1,
		"斗地主",
		3,
		1,
		[]poker.PokerCard{},
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
func (dou Doudizhu)DealCards(){

}
