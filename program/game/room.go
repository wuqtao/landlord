package game

import (
	"sync"
	"errors"
)

type GameList struct {
	list map[int]IGame
	sync.RWMutex
}

type Room map[int]*GameList    //分游戏类型，每种游戏单独切片

var room Room = nil

func init(){
	room = Room{}
}
//获得全局room单例对象
func GetRoom() Room {
	return room
}

func (r Room)AddGame(gameID int,game IGame) int{
	list,ok := r[gameID]
	if !ok {
		r[gameID] = &GameList{
			list:make(map[int]IGame),
		}
		list = r[gameID]
	}
	//todo 暂时用数量来表示
	gameCount := len(list.list)

	list.Lock()
	list.list[gameCount] = game
	list.Unlock()

	return gameCount
}

func (r Room)GetGame(gameType int,gameID int) (IGame,error){
	list,ok := r[gameType]
	if ok{
		game,ok := list.list[gameID]
		if ok {
			return game,nil
		}else{
			return nil,errors.New("不存在该游戏")
		}
	}else{
		return nil,errors.New("不存在该类型的游戏")
	}
}