package games

import (
	doudizhu2 "chessSever/program/logic/game/games/doudizhu"
)

type Game struct {
	Id        int    //游戏ID
	Name      string //游戏name
	PlayerNum int    //游戏参与人数
	DeckNum int      //几副牌
}

var doudizhu *Game
var baohuang *Game
var shengji *Game
var zhajinhua *Game

func GetGame(name string) *Game {
	switch name {
	case "doudizhu":
		return doudizhu2.GetDoudizhu()
	case "shengji":
		return nil
	case "baohuang":
		return nil
	case "zhajinhua":
		return nil
	default:
		return nil
	}
}





