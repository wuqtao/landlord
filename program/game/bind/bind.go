package bind

import (
	"sync"
	"chessSever/program/player"
	"chessSever/program/game"
	"errors"
)

type PlayerGameDic struct {
	sync.RWMutex
	Dic map[*player.Player]game.IGame
}

var dic PlayerGameDic

func init(){
	dic = PlayerGameDic{
		Dic: make(map[*player.Player]game.IGame),
	}
}

func BindPlayerGame(p *player.Player,game game.IGame){
	dic.Lock()
	dic.Dic[p] = game
	dic.Unlock()
}

func GetPlayerGame(p *player.Player) (game.IGame,error){
	dic.RLock()
	defer dic.RUnlock()
	game,ok := dic.Dic[p]
	if ok {
		return game,nil
	}else{
		return nil,errors.New("该player没有关联的game")
	}
}

func GetDic() *PlayerGameDic{
	return &dic
}