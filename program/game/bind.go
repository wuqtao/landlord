package game

import (
	"sync"
	"errors"
)

type PlayerGameDic struct {
	sync.RWMutex
	Dic map[IPlayer]IGame
}

var dic PlayerGameDic

func init(){
	dic = PlayerGameDic{
		Dic: make(map[IPlayer]IGame),
	}
}

func BindPlayerGame(p IPlayer,game IGame){
	dic.Lock()
	dic.Dic[p] = game
	dic.Unlock()
}

func GetPlayerGame(p IPlayer) (IGame,error){
	dic.RLock()
	defer dic.RUnlock()
	game,ok := dic.Dic[p]
	if ok {
		return game,nil
	}else{
		return nil,errors.New("该player没有关联的game")
	}
}