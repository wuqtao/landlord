package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"chessSever/program/logic/player"
	"chessSever/program/logic/game/games"
	"fmt"
	"runtime/debug"
)

var addr = flag.String("addr", "localhost:9999", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var userID int

var m sync.RWMutex

func echo(w http.ResponseWriter, r *http.Request) {

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer con.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	//暂时用变量模拟用户登陆，后续从数据库读取用户信息，实例化用户，游戏过程中用redis来暂存游戏信息，用户推出后持久化到数据库
	m.Lock()
	userID++
	nowId := userID
	m.Unlock()

	currPlayer := player.NewPlayer(nowId,strconv.Itoa(nowId),con,"1headPic")
	if currPlayer.Id == 1{
		currPlayer.CreateTable(games.DouDiZhu)
	}else{
		currPlayer.JoinTable(player.GetRoom().GetAllTable()[0])
	}
	//启动一个goroutine监听该客户端发来的消息
	go func() {
		defer wg.Done()
		defer func(){
			if p := recover();p != nil{
				fmt.Printf("panic recover! p: %v", p)
				debug.PrintStack()
			}
		}()
		for{
			msgType,msg,err := con.ReadMessage()
			if err == nil{
				switch msgType {
				case websocket.TextMessage:
					//同桌用户交流，包含对话流程和出牌流程
					currPlayer.ResolveMsg(msg)
				case websocket.CloseMessage:
					fmt.Println("链接关闭")
					break
					//离开桌子流程，后续包含断线保持，自动出牌
				default:

				}
			}else{
				break
			}
		}
	}()

	wg.Wait()
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w,r,"/pages/index.html",301)
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./views/pages"))))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
