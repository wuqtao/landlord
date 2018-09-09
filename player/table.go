package player

import (
	"chessSever/game"
	"errors"
	"strconv"
	"sync"
)

/*
	定义游戏桌对象
*/
type Table struct {
	key          string     //桌子key,用于从room索引中查找桌子
	players      []*Player  //玩家数组
	game         *game.Game //该桌玩的游戏
	sync.RWMutex            //操作playNum以及player时加锁
}

func newTable(player *Player, game *game.Game) *Table {
	table := Table{
		game: game,
		key:  "table" + strconv.Itoa(player.id),
	}
	table.joinRoom()
	table.addPlayer(player)
	return &table
}
func (t *Table) joinRoom() {
	getRoom().addTable(t.key, t)
}

func (t *Table) destory() {
	if len(t.players) >= 0 {
		for _, p := range t.players {
			p.leaveTable()
		}
	}
	getRoom().removeTable(t.key)
}

func (t *Table) addPlayer(player *Player) error {
	if len(t.players) >= t.game.PlayerNum {
		return errors.New("该卓玩家已经满了")
	} else {
		t.RLock()
		t.players = append(t.players, player)
		t.RUnlock()
		return nil
	}
}

func (t *Table) removePlayer(palyer *Player) {
	for i, p := range t.players {
		if p == palyer {
			t.RLock()
			t.players[i] = nil
			palyer.tableKey = ""
			t.RUnlock()
			break
		}
	}

	if len(t.players) == 0 {
		t.destory()
	}
}
