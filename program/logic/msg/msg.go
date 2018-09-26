package msg

const(
	TypeOfReady = iota    //准备
	TypeOfUnReady		  //取消准备
	TypeOfJoinTable	      //加入桌子
	TypeOfLeaveTable      //离开桌子

	TypeOfHint            //提示
	TypeOfPlay			  //出牌
	TypeOfPass			  //过牌

	TypeOfAuto			  //托管
)


