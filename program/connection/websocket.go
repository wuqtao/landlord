package connection

import (
	"sync"
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	sync.RWMutex
	Conn *websocket.Conn
}

func NewWebSocketConnection(conn *websocket.Conn) *WebSocketConnection{
	return &WebSocketConnection{
		Conn:conn,
	}
}

func (conn *WebSocketConnection)SendMsg(msg []byte){
	conn.Lock()
	defer conn.Unlock()
	conn.Conn.WriteMessage(websocket.TextMessage,msg)
}

func (conn *WebSocketConnection)SendMsgWithType(msgType int,msg []byte){
	conn.Lock()
	defer conn.Unlock()
	conn.Conn.WriteMessage(msgType,msg)
}
