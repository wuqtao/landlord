const 	TypeOfReady = 0;    //准备
const   TypeOfUnReady =1;	 	  //取消准备
const   TypeOfJoinTable	= 2;      //加入桌子
const  TypeOfLeaveTable = 3;    //离开桌子
const  TypeOfHint       =  4;     //提示
const  TypeOfPlayCard   =  5;  //出牌
const  TypeOfPass       =  6;  //过牌
const  TypeOfAuto		= 7;  //托管
const  TypeOfSendCard	= 8;  //发牌
const  TypeOfCallScore   = 9;  //抢地主叫分
const    TypeOfConfirm   = 10;   //客户端出牌等操作确认信息


var print = function(message) {
    var d = document.createElement("div");
    d.innerHTML = message;
    $("#output").append(d);
};
var ws;
$(function(){
    $("#open").bind('click',function(){
        if (ws) {
            return false;
        }
        ws = new WebSocket("ws://localhost:9999/echo");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            data = JSON.parse(evt.data)
            switch(data.MsgType){
                case TypeOfReady:
                    break;
                case TypeOfUnReady:
                    break;
                case TypeOfJoinTable:
                    break;
                case TypeOfLeaveTable:
                    break;
                case TypeOfHint:
                    break;
                case TypeOfPlayCard:
                    break;
                case TypeOfPass:
                    break;
                case TypeOfAuto:
                    break;
                case TypeOfSendCard:
                    console.log(data)
                    break;
                case TypeOfCallScore:
                    $("#divScore").show()
                    break;
                case TypeOfConfirm:
                    break;
                default:
                    console.log(evt.data)
            }
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    });
    $("#send").bind('click',function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + $("#input").val());
        ws.send(input.value);
        return false;
    });
    $("#close").bind('click',function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    });

    $("#ready").bind('click',function () {
        var readyMsg = {}
        readyMsg.msgType = TypeOfReady;
        ws.send(JSON.stringify(readyMsg))
        return false;
    });

    $("unReady").bind('click',function () {

    })
})

function callScore() {
    msg = {};
    msg.MsgType = TypeOfCallScore;
    msg.Data = {
        "score":$("#score").val()
    }
    ws.send(JSON.stringify(msg))
}