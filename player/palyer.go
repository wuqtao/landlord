package player

import (
	"chessSever/game"
	"github.com/gorilla/websocket"
	"sync"
)

/**
定义游戏玩家对象
*/
type Player struct {
	id       int             //注册用户id
	nickName string          //用户昵称
	conn     *websocket.Conn //用户socket链接
	headPic  string          //用户头像
	table    *Table          //桌子索引
	sync.RWMutex
}

func newPlayer(id int, nickName string, conn *websocket.Conn, headPic string) *Player {
	player := Player{
		id:       id,
		nickName: nickName,
		headPic:  headPic,
		conn:     conn,
	}
	return &player
}

//按照桌号加入牌桌
func (p *Player) joinTable(key string) error {
	err := getRoom().getTable(key).addPlayer(p)
	return err
}

//开牌桌
func (p *Player) createTable(game *game.Game) {
	table := newTable(p, game)
	table.addPlayer(p)
	p.RLock()
	p.table = table
	p.RUnlock()
}

func (p *Player) leaveTable() {
	p.table.removePlayer(p)
	p.RLock()
	p.table = nil
	p.RUnlock()
}
