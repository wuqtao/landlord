package player

import "sync"

type Room struct {
	tables map[string]*Table
	sync.RWMutex
}

var room *Room = nil

//获得全局room对象
func getRoom() *Room {
	if room != nil {
		return room
	} else {
		return &Room{
			tables: make(map[string]*Table),
		}
	}
}

func (r *Room) getTable(key string) *Table {
	return r.tables[key]
}

func (r *Room) addTable(key string, table *Table) {
	r.RLock()
	r.tables[key] = table
	r.RUnlock()
}

func (r *Room) removeTable(key string) {
	r.RLock()
	delete(r.tables, key)
	r.RUnlock()
}

func (r *Room) tablesCounts() int {
	return len(r.tables)
}
