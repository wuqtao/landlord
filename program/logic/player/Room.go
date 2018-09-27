package player

import (
	"sync"
	"fmt"
	"strconv"
)

type Room struct {
	tables map[string]*Table
	sync.RWMutex
}

var room *Room = nil

//获得全局room单例对象
func GetRoom() *Room {
	if room != nil {
		return room
	} else {
		room = &Room{
			tables: make(map[string]*Table),
		}
		return room
	}
}

func (r *Room) getTable(key string) *Table {
	return r.tables[key]
}

func (r *Room) addTable(key string, table *Table) {
	r.Lock()
	r.tables[key] = table
	fmt.Println("添加桌子"+key+"后，当前房间桌子数量为"+strconv.Itoa(len(r.tables)))
	r.Unlock()
}

func (r *Room) removeTable(key string) {
	r.Lock()
	delete(r.tables, key)
	fmt.Println("移除桌子"+key+"后，当前房间桌子数量为"+strconv.Itoa(r.tablesCounts()))
	r.Unlock()
}

func (r *Room) tablesCounts() int {
	return len(r.tables)
}

func (r *Room) GetAllTable() []*Table{

	if len(room.tables) == 0{
		return nil
	}

	tables := []*Table{}

	for _,table := range room.tables{
		tables = append(tables,table)
	}
	return tables
}
