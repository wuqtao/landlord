package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"chessSever/player"
	"chessSever/game"
	"fmt"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

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
	stopChan := make(chan int)
	go func() {
		m.Lock()
		userID++
		nowId := userID
		m.Unlock()
		player := player.NewPlayer(nowId,strconv.Itoa(nowId),con,"1headPic")
		if player.Id == 1{
			player.CreateTable(&game.Game{})
			fmt.Println("createTable")
		}else{
			fmt.Println("joinTable")
			player.JoinTable("table1")
		}
		//
		//
		//for {
		//	mt, message, err := con.ReadMessage()
		//	if err != nil {
		//		log.Println("read:", err)
		//
		//		break
		//	}
		//	log.Printf("recv: %s", message)
		//	err = con.WriteMessage(mt, message)
		//	if err != nil {
		//		log.Println("write:", err)
		//		break
		//	}
		//}
		//stopChan <- 1
	}()
	//go func() {
		//for i := 0; i < 10; i++ {
		//	err := con.WriteMessage(websocket.TextMessage, []byte(""+string(i)))
		//	if err != nil {
		//		log.Println("write:", err)
		//		break
		//	}
		//	time.Sleep(time.Second * 1)
		//}
		//stopChan <- 2
	//}()
	sig := <-stopChan
	log.Println("stop with unkonwn signal" + strconv.Itoa(sig))
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
