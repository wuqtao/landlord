package player

import (
	"errors"
	"strconv"
	"sync"
	"fmt"
	"chessSever/program/logic/game/games"
	"chessSever/program/logic/game/poker"
)

/*
	定义游戏桌对象
*/
type Table struct {
	key          string     				//桌子key,用于从room索引中查找桌子
	players      []*Player  				//玩家数组
	game         games.Game 				//该桌玩的游戏
	sync.RWMutex            				//操作playNum以及player时加锁
	CurrPokerCards []*poker.PokerCard  		//当前出的牌
	CurrPalyerIndex int 					//当前出牌的玩家切片index
}
//创建桌子
func newTable(player *Player, game *games.Game) *Table {
	table := Table{
		game: game,
		key:  "table" + strconv.Itoa(player.Id),
	}
	fmt.Println("创建新桌子"+"table" + strconv.Itoa(player.Id))
	//桌子加入房间
	table.joinRoom()
	//将创建者加入桌子
	table.addPlayer(player)
	return &table
}
func (t *Table) joinRoom() {
	getRoom().addTable(t.key, t)
}
//销毁桌子
func (t *Table) destory() {
	t.Lock()
	if len(t.players) >= 0 {
		for _, p := range t.players {
			p.LeaveTable()
		}
	}
	getRoom().removeTable(t.key)
	fmt.Println("桌子"+t.key+"销毁")
	t.Unlock()
}
//增加玩家
func (t *Table) addPlayer(player *Player) error {
	t.Lock()
	var err error = nil
	if len(t.players) >= t.game.GetPlayerNum() {
		err = errors.New("该卓玩家已经满了")
		fmt.Println("桌子"+t.key+"已经满了")
	} else {
		t.players = append(t.players, player)
		fmt.Println("桌子"+t.key+"新加入一个玩家"+strconv.Itoa(player.Id)+"，当前玩家数是"+strconv.Itoa(len(t.players)))
	}
	t.Unlock()
	return err
}
//移除玩家
func (t *Table) removePlayer(player *Player) {
	t.Lock()
	for i, p := range t.players {
		if p == player {
			t.Lock()
			t.players[i] = nil
			break
		}
	}
	fmt.Println("桌子"+t.key+"移除玩家"+strconv.Itoa(player.Id)+"，当前玩家数是"+strconv.Itoa(len(t.players)))
	t.Unlock()
	if len(t.players) == 0 {
		t.destory()
	}
}


