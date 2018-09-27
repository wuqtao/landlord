package msg

const(
	TypeOfReady = iota    //准备
	TypeOfUnReady		  //取消准备
	TypeOfJoinTable	      //加入桌子
	TypeOfLeaveTable      //离开桌子

	TypeOfHint      //提示
	TypeOfPlayCard  //出牌
	TypeOfPass      //过牌

	TypeOfAuto			  //托管
	TypeOfSendCard		  //发牌
	TypeOfCallScore       //抢地主叫分
	TypeOfConfirm         //客户端出牌等操作确认信息
)
/*
	确认消息
	{
		"msgType":TypeOfConfirm
		"data"{

		}
	}
	发牌消息形式
	{
		"msgType":TypeOfSendCard,
		"data":{
			pokerCars[{
				"carIndex":,    //当前玩家手中的index
				"carNum":,      //牌的数字
				"carColor":,    //牌的花色
				""
			},]
		}
	}

	出牌消息
	{
		"msgType":TypeOfSendCard,
		"data":{
			pokerCars[{
				"carIndex":,    //当前玩家手中的index
			},]
		}
	}


 */

