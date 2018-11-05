package game

import (
	"sync"
	"errors"
	"fmt"
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
	defer dic.Unlock()
	_,ok := dic.Dic[p]
	if ok {
		fmt.Println("该玩家已绑定游戏，绑定失败")
	}else{
		dic.Dic[p] = game
	}
}

func UnbindPlayerGame(p IPlayer,game IGame){
	dic.Lock()
	defer dic.Unlock()
	currGame,ok := dic.Dic[p]
	if ok {
		if currGame == game{
			delete(dic.Dic,p)
		}else{
			fmt.Println("玩家已绑定游戏，不是当前给定的游戏，解绑失败")
		}
	}else{
		fmt.Println("玩家未绑定游戏，解绑失败")
	}
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