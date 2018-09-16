package player

import (
	"errors"
	"strconv"
	"sync"
	"fmt"
	"chessSever/program/logic/game/games"
	"chessSever/program/logic/game/poker"
	"chessSever/program/logic/util"
)

/*
	定义游戏桌对象
*/
type Table struct {
	key          string     //桌子key,用于从room索引中查找桌子
	players      []*Player  //玩家数组
	game         *games.Game //该桌玩的游戏
	sync.RWMutex            //操作playNum以及player时加锁
	PokerCards []poker.PokerCard
}

func newTable(player *Player, game *games.Game) *Table {
	table := Table{
		game: game,
		key:  "table" + strconv.Itoa(player.Id),
	}
	porkerCards := []poker.PokerCard{}
	//根据游戏类型，初始化该桌子的扑克牌数组
	for i:=0;i<table.game.DeckNum;i++{
		for _,card := range poker.CreateDeck().Cards{
			porkerCards = append(porkerCards,*card)
		}
	}
	table.PokerCards = porkerCards

	fmt.Println("创建新桌子"+"table" + strconv.Itoa(player.Id))
	table.joinRoom()
	table.addPlayer(player)
	return &table
}
func (t *Table) joinRoom() {
	getRoom().addTable(t.key, t)
}

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

func (t *Table) addPlayer(player *Player) error {
	t.Lock()
	var err error = nil
	if len(t.players) >= t.game.PlayerNum {
		err = errors.New("该卓玩家已经满了")
		fmt.Println("桌子"+t.key+"已经满了")
	} else {
		t.players = append(t.players, player)
		fmt.Println("桌子"+t.key+"新加入一个玩家"+strconv.Itoa(player.Id)+"，当前玩家数是"+strconv.Itoa(len(t.players)))
	}
	t.Unlock()
	return err
}

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
//洗牌
func (t *Table) flushPokerCards(){
	util.Random(t.PokerCards,len(t.PokerCards))
}

//发牌
func (t *Table) dipatchPokerCards(){

}


