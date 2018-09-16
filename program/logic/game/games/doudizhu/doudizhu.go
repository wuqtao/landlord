package doudizhu

import (
	"chessSever/program/logic/game/games"
)

var doudizhu *games.Game

func init(){
	doudizhu = &games.Game{
		Id:1,
		Name:"斗地主",
		PlayerNum:3,
		DeckNum:1,
	}
}

func GetDoudizhu() *games.Game{
	return doudizhu
}
