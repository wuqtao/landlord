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
	"encoding/json"
	"fmt"
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
	go player.HandlerUserMsg(&wg,con,currPlayer)

	wg.Wait()
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w,r,"/pages/index.html",301)
}
type A struct{
	a int
	b string
}
type Host struct {
	IP string
	Name string
	a []A
}

func main() {

	m := Host{
		Name:"Sky",
		IP:"192.168.23.92",
		a:[]A{A{1,"11"},{2,"22"}},
	}
	fmt.Println(m)
	b, err := json.Marshal(m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}


	fmt.Println("json:", string(b))
}
//func main() {
	//flag.Parse()
	//log.SetFlags(0)
	//http.HandleFunc("/echo", echo)
	//http.HandleFunc("/", home)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./views/static"))))
	//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./views/pages"))))
	//log.Fatal(http.ListenAndServe(*addr, nil))
//}
