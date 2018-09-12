package game

type Game struct {
	Id        int    //游戏ID
	Name      string //游戏name
	PlayerNum int    //游戏参与人数
}

var doudizhu *Game;
var baohuang *Game;
var shengji *Game;
var zhajinhua *Game;

//下面定义几个游戏范例
func GetDoudizhu() *Game{
	if(doudizhu != nil){
		return doudizhu
	}
	doudizhu = &Game{
		1,
		"斗地主",
		3,
	}
	return doudizhu
}
//扎金花，这个游戏可以多人玩,暂定5人
func GetZhajinhua() *Game{
	if(zhajinhua != nil){
		return zhajinhua
	}
	zhajinhua = &Game{
		2,
		"扎金花",
		5,
	}
	return zhajinhua
}
//升级
func GetShengji() *Game{
	if(shengji != nil){
		return shengji
	}
	shengji = &Game{
		3,
		"升级",
		4,
	}
	return shengji
}
//保皇
func GetBaohuang() *Game{
	if(baohuang != nil){
		return baohuang
	}
	baohuang = &Game{
		4,
		"斗地主",
		3,
	}
	return baohuang
}



