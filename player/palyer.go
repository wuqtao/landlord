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
	Id       int             //注册用户id
	NickName string          //用户昵称
	Conn     *websocket.Conn //用户socket链接
	HeadPic  string          //用户头像
	Table    *Table          //桌子索引
	sync.RWMutex
}

func NewPlayer(id int, nickName string, conn *websocket.Conn, headPic string) *Player {
	player := Player{
		Id:       id,
		NickName: nickName,
		HeadPic:  headPic,
		Conn:     conn,
	}
	return &player
}

//按照桌号加入牌桌
func (p *Player) JoinTable(key string) error {
	err := getRoom().getTable(key).addPlayer(p)
	return err
}

//开牌桌
func (p *Player) CreateTable(game *game.Game) {
	table := newTable(p, game)
	table.addPlayer(p)
	p.Lock()
	p.Table = table
	p.Unlock()
}

func (p *Player) LeaveTable() {
	p.Table.removePlayer(p)
	p.Lock()
	p.Table = nil
	p.Unlock()
}
