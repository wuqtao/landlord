package player

type PlayerUserDic map[int]*Player

var playerUserDic PlayerUserDic

func init(){
	playerUserDic = make(PlayerUserDic)
}

func GetPlayer(userID int) *Player{
	p,ok := playerUserDic[userID]
	if ok {
		return p
	}else{
		return nil
	}
}

func SetPlayer(userId int,p *Player){
	playerUserDic[userId] = p
}


