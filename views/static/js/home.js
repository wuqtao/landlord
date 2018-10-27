const MSG_TYPE_OF_READY = 0;    //准备
const MSG_TYPE_OF_UN_READY =1;	 	  //取消准备
const MSG_TYPE_OF_JOIN_TABLE	= 2;      //加入桌子
const MSG_TYPE_OF_LEAVE_TABLE = 3;    //离开桌子
const MSG_TYPE_OF_HINT       =  4;     //提示
const MSG_TYPE_OF_PLAY_CARD   =  5;  //出牌
const MSG_TYPE_OF_PASS       =  6;  //过牌
const MSG_TYPE_OF_AUTO		= 7;  //托管
const MSG_TYPE_OF_SEND_CARD	= 8;  //发牌
const MSG_TYPE_OF_CALL_SCORE   = 9;  //抢地主叫分
const MSG_TYPE_OF_CONFIRM   = 10;   //客户端出牌等操作确认信息
const MSG_TYPE_OF_CALL_SCORE_TIME_OUT = 11; //叫地主超时
const MSG_TYPE_OF_PLAY_ERROR = 12          //出牌错误
const TYPE_OF_PLAY_CARD_SUCCESS  = 13      //出牌成功
const TYPE_OF_TABLE_BRODCAST = 14          //游戏桌子广播消息
const MSG_TYPE_OF_SCORE_CHANGE = 15        //牌局分数变化
const MSG_TYPE_OF_SETTLE_SCORE = 16    //结算玩家分数
const MSG_TYPE_OF_GAME_OVER = 17           //游戏结束
const MSG_TYPE_OF_LOGIN = 18           //玩家登陆成功
const MSG_TYPE_OF_SEND_BOTTOM_CARDS = 19 //发底牌
var ws;
var lastPlayCardIndex = [];
var currPlayerId = -1;
var currPlayerIndex = -1;
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
                case MSG_TYPE_OF_LOGIN:
                    currPlayerId = data.ID
                    break;
                case MSG_TYPE_OF_HINT:
                    break;
                case MSG_TYPE_OF_PLAY_CARD:
                    $("#divPlay").show();
                    break;
                case MSG_TYPE_OF_SEND_CARD:
                    $("#userCards").html('');
                    $.each(data.Cards,function(i,o){
                        $("#userCards").append(String.format($("#tempPlayCard").html(),o.Card.CardName,o.Card.CardSuit,o.Index));
                    })
                    break;
                case MSG_TYPE_OF_CALL_SCORE:
                    $("#divScore").show()
                    print("请叫分");
                    break;
                case MSG_TYPE_OF_CALL_SCORE_TIME_OUT:
                    $("#divScore").hide();
                    break;
                case MSG_TYPE_OF_PLAY_ERROR:
                    alert(data.Msg);
                    break;
                case TYPE_OF_PLAY_CARD_SUCCESS:
                    $("#divPlay").hide()
                    alert();
                    for (i=0;i<lastPlayCardIndex.length;i++){
                        $("#cardDiv"+lastPlayCardIndex[i]).remove();
                    }
                    break;
                case TYPE_OF_TABLE_BRODCAST:
                    switch(data.SubMsgType){
                        case MSG_TYPE_OF_READY:
                            // console.log(data);
                            if(data.PlayerId == currPlayerId){
                                print("已准备");
                            }

                            $("#PlayerMsg"+data.PlayerId).html("已准备");
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_UN_READY:
                            if(data.PlayerId == currPlayerId){
                                print("取消准备");
                            }
                            $("#PlayerMsg"+data.PlayerId).html("取消准备");
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_LEAVE_TABLE:
                        case MSG_TYPE_OF_JOIN_TABLE:
                            $("#rightPlayer").html('');
                            $("#leftPlayer").html('');
                            for (x in data.PlayerIndexIdDic){
                                if(x == "id"+currPlayerId){
                                    currPlayerIndex = data.PlayerIndexIdDic[x];
                                    break;
                                }
                            }
                            for (x in data.PlayerIndexIdDic){
                                if(data.PlayerIndexIdDic[x] == currPlayerIndex){
                                    continue;
                                }
                                id = x.replace("id",'');
                                if (data.PlayerIndexIdDic[x] > currPlayerIndex){
                                    if(data.PlayerIndexIdDic[x] == currPlayerIndex+1){
                                        $("#rightPlayer").html(String.format($("#tempDivPlayerInfo").html(),id));
                                        $("#PlayerName"+id).html("玩家"+x);
                                        $("#PlayerMsg"+id).html("玩家"+x+"加入游戏");
                                    }else{
                                        $("#leftPlayer").html(String.format($("#tempDivPlayerInfo").html(),id));
                                        $("#PlayerName"+id).html("玩家"+x);
                                        $("#PlayerMsg"+id).html("玩家"+x+"加入游戏");
                                    }

                                }else{
                                    if(data.PlayerIndexIdDic[x] == currPlayerIndex-1){
                                        $("#leftPlayer").html(String.format($("#tempDivPlayerInfo").html(),id));
                                        $("#PlayerName"+id).html("玩家"+x);
                                        $("#PlayerMsg"+id).html("玩家"+x+"加入游戏");
                                    }else{
                                        $("#rightPlayer").html(String.format($("#tempDivPlayerInfo").html(),id));
                                        $("#PlayerName"+id).html("玩家"+x);
                                        $("#PlayerMsg"+id).html("玩家"+x+"加入游戏");
                                    }
                                }
                            }
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_SEND_BOTTOM_CARDS:
                            $.each(data.Cards,function(i,o){
                                $("#LoardCards"+data.PlayerId).append(String.format($("#tempCard").html(),o.CardName,o.CardSuit));
                            })
                            break;
                        case MSG_TYPE_OF_PLAY_CARD:
                            $("#PlayerMsg"+data.PlayerId).html("玩家"+data.PlayerId+"出牌");
                            $("#divPlayCards").html('');
                            $.each(data.Cards,function(i,o){
                                $("#divPlayCards").append(String.format($("#tempCard").html(),o.CardName,o.CardSuit));
                            })
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_PASS:
                            $("#PlayerMsg"+data.PlayerId).html("玩家"+data.PlayerId+"过牌");
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_CALL_SCORE:
                            $("#PlayerMsg"+data.PlayerId).html("玩家"+data.PlayerId+"叫地主"+data.Score+"分");
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_SCORE_CHANGE:
                            print("底分变为"+data.Score)
                            console.log(data);
                            break;
                        case MSG_TYPE_OF_GAME_OVER:
                            print("游戏结束")
                            console.log(data);
                        default:
                            print("未知消息子类型");
                            console.log(data);
                    }
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
    readyMsg.MsgType = MSG_TYPE_OF_READY;
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
    msg.MsgType = MSG_TYPE_OF_CALL_SCORE;
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
    msg = {};
    cardIndex = [];
    lastPlayCardIndex = [];
    msg.MsgType = MSG_TYPE_OF_PASS;
    ws.send(JSON.stringify(msg));
}

function playCards(){
    msg = {};
    msg.MsgType = MSG_TYPE_OF_PLAY_CARD;
    cardIndex = [];
    lastPlayCardIndex = []
    $("#userCards").find('div').each(function(i,o){
        if($(o).hasClass('chooseYes')){
             cardIndex.push($($(o).find('input')[0]).val());
             lastPlayCardIndex.push($($(o).find('input')[0]).val());
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
    cardIndex = [];
    lastPlayCardIndex = [];
    $("#divPlay").show();
    msg = {};
    msg.MsgType = MSG_TYPE_OF_HINT;
    ws.send(JSON.stringify(msg));
}