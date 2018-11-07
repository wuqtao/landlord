package game

import (
	"sync"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
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
//将game加入房间，并返回game的id
func (r Room)AddGame(gameType int,game IGame) int{
	list,ok := r[gameType]
	if !ok {
		r[gameType] = &GameList{
			list:make(map[int]IGame),
		}
		list = r[gameType]
	}
	//todo 暂时用数量来表示

	list.Lock()
	gameCount := len(list.list)
	logrus.Info("新游戏加入房间:"+GetGameName(gameType)+strconv.Itoa(gameType)+strconv.Itoa(gameCount))
	list.list[gameCount] = game
	list.Unlock()

	return gameCount
}
//根据游戏类型和gameId查找game
func (r Room)GetGame(gameType int,gameID int) (IGame,error){
	list,ok := r[gameType]
	if ok{
		game,ok := list.list[gameID]
		if ok {
			return game,nil
		}else{
			logrus.Error("不存在该游戏")
			return nil,errors.New("不存在该游戏")
		}
	}else{
		logrus.Error("不存在该类型的游戏")
		return nil,errors.New("不存在该类型的游戏")
	}
}