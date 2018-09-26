package player

import (
	"chessSever/program/logic/game/games"
	"sync"
	"chessSever/program/logic/game/poker"
	"chessSever/program/logic/game"
	"strconv"
	"fmt"
	"errors"
)

/*
	定义游戏桌对象
*/
type Table struct {
	Key          string     				//桌子key,用于从room索引中查找桌子
	Players      []*Player  				//玩家数组
	Game         games.IGame 				//该桌玩的游戏
	sync.RWMutex            				//操作playNum以及player时加锁
	CurrPokerCards []*poker.PokerCard  		//当前出的牌
	CurrPalyerIndex int 					//当前出牌的玩家切片index

}
//创建桌子
func newTable(player *Player, gameName string) *Table {
	game := game.GetGame(gameName)
	table := Table{
		Game: game,
		Key:  "table" + strconv.Itoa(player.Id),
	}
	fmt.Println("创建新桌子"+"table" + strconv.Itoa(player.Id))
	//桌子加入房间
	table.joinRoom()
	//将创建者加入桌子
	table.addPlayer(player)
	return &table
}
//加入房间
func (t *Table) joinRoom() {
	getRoom().addTable(t.Key, t)
}
//销毁桌子
func (t *Table) destory() {
	t.Lock()
	if len(t.Players) >= 0 {
		for _, p := range t.Players {
			p.LeaveTable()
		}
	}
	getRoom().removeTable(t.Key)
	fmt.Println("桌子"+t.Key+"销毁")
	t.Unlock()
}
//增加玩家
func (t *Table) addPlayer(player *Player) error {
	t.Lock()
	var err error = nil
	if len(t.Players) >= t.Game.GetPlayerNum() {
		err = errors.New("该卓玩家已经满了")
		fmt.Println("桌子"+t.Key+"已经满了")
	} else {
		t.Players = append(t.Players, player)
		fmt.Println("桌子"+t.Key+"新加入一个玩家"+strconv.Itoa(player.Id)+"，当前玩家数是"+strconv.Itoa(len(t.Players)))
	}
	t.Unlock()
	return err
}
//移除玩家
func (t *Table) removePlayer(player *Player) {
	t.Lock()
	for i, p := range t.Players {
		if p == player {
			t.Lock()
			t.Players[i] = nil
			break
		}
	}
	fmt.Println("桌子"+t.Key+"移除玩家"+strconv.Itoa(player.Id)+"，当前玩家数是"+strconv.Itoa(len(t.Players)))
	t.Unlock()
	if len(t.Players) == 0 {
		t.destory()
	}
}


