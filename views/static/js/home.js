const TypeOfReady = 0;    //准备
const TypeOfUnReady =1;	 	  //取消准备
const TypeOfJoinTable	= 2;      //加入桌子
const TypeOfLeaveTable = 3;    //离开桌子
const TypeOfHint       =  4;     //提示
const TypeOfPlayCard   =  5;  //出牌
const TypeOfPass       =  6;  //过牌
const TypeOfAuto		= 7;  //托管
const TypeOfSendCard	= 8;  //发牌
const TypeOfCallScore   = 9;  //抢地主叫分
const TypeOfConfirm   = 10;   //客户端出牌等操作确认信息
const TypeOfCallScoreTimeOut = 11; //叫地主超时

var ws;

function print(message) {
    var d = document.createElement("div");
    d.innerHTML = message;
    $("#output").append(d);
};

function openConnection(){
        if (ws) {
            return false;
        }
        ws = new WebSocket("ws://localhost:8888/echo");
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
                    $("#divPlay").show();
                    break;
                case TypeOfPass:
                    break;
                case TypeOfAuto:
                    break;
                case TypeOfSendCard:
                    $("#userCards").html('');
                    $.each(data.Cards,function(i,o){
                        $("#userCards").append(String.format($("#tempCard").html(),o.Card.CardName,o.Card.CardSuit,o.Index));
                    })
                    break;
                case TypeOfCallScore:
                    $("#divScore").show()
                    print("请叫分");
                    break;
                case TypeOfConfirm:
                    break;
                case TypeOfCallScoreTimeOut:
                    $("#divScore").hide();
                    break;
                default:
                    console.log(evt.data)
            }
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
}


function closeConnection() {
    if (!ws) {
        return false;
    }
    ws.close();
}

function send(evt) {
    if (!ws) {
        return false;
    }
    ws.send(input.value);
}

function ready() {
    var readyMsg = {}
    readyMsg.MsgType = TypeOfReady;
    ws.send(JSON.stringify(readyMsg));
}

function unReady(){
    
}

function chooseCard(obj){
    if($(obj).hasClass('chooseYes')){
        $(obj).removeClass('chooseYes');
    }else{
        $(obj).addClass('chooseYes');
    }
}

function sendScore() {
    msg = {};
    msg.MsgType = TypeOfCallScore;
    msg.Data = {
        "Score":$("#score").val()
    }
    if(!ws){
        return false
    }
    ws.send(JSON.stringify(msg));
    $("#divScore").hide();
}

function pass(){
    $("#divPlay").show();
    msg = {};
    msg.MsgType = TypeOfPass;
    ws.send(JSON.stringify(msg));
}

function playCards(){
    $("#divPlay").show();
    msg = {};
    msg.MsgType = TypeOfPlayCard;
    cardIndex = [];
    $("#userCards").find('div').each(function(i,o){
        if($(o).hasClass('chooseYes')){
             cardIndex.push($($(o).find('input')[0]).val());
        }
    })
    if(cardIndex.length == 0){
        alert('请选择牌，再点击出牌');
    }
    msg.Data = {
        'CardIndex':cardIndex
    }
    ws.send(JSON.stringify(msg));
}

function cardHints(){
    $("#divPlay").show();
    msg = {};
    msg.MsgType = TypeOfHint;
    ws.send(JSON.stringify(msg));
}