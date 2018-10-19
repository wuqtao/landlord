package poker

//判断给定的扑克牌切片中是否存在3个一样大小的牌,以及不一样的牌的index数组
func HasThreeSameValueCard(pokers []*PokerCard) (bool,[]int){

	if len(pokers) < 3 || len(pokers) > 5{
		return false,nil
	}

	indexs := []int{}
	num := 0
	currValue := -1
	for i,card := range pokers{
		if i == 0 {
			currValue = card.CardValue
			num++
		}else{
			if card.CardValue == currValue{
				num++
			}else{
				switch num {
				case 1:
					indexs = append(indexs,i-1)
					currValue = card.CardValue
					num = 1
				case 2:
					indexs = append(indexs,i-1)
					indexs = append(indexs,i-2)
					currValue = card.CardValue
					num = 1
				case 3:
					indexs = append(indexs,i)
				case 4:
					indexs = append(indexs,i)
				}

			}
		}
	}

	if num >= 3{
		return true,indexs
	}else{
		return false,indexs
	}
}
//是否是对子
func IsPair(pokers []*PokerCard) bool{
	if len(pokers) != 2 {
		return false
	}

	if(pokers[0].CardValue == pokers[1].CardValue){
		return true
	}else{
		return false
	}
}
//是否是王炸
func IsJokerBomb(pokers []*PokerCard) bool{

	if len(pokers) != 2{
		return false
	}
	if(pokers[0].CardValue == PokerBlackJoker && pokers[1].CardValue == PokerRedJoker) ||
		(pokers[1].CardValue == PokerBlackJoker && pokers[0].CardValue == PokerRedJoker){
		return true
	}else{
		return false
	}
}
//是否是炸弹
func IsCommonBomb(pokers []*PokerCard) bool{

	if len(pokers) != 4{
		return false
	}
	if pokers[0].CardValue == pokers[1].CardValue  && pokers[2].CardValue == pokers[3].CardValue &&
		pokers[0].CardValue == pokers[2].CardValue{
		return true
	}else{
		return false
	}
}
//是否是三代一或三代二
func IsThreePlus(pokers []*PokerCard) bool{

	is,indexs := HasThreeSameValueCard(pokers)
	if !is{
		return false
	}else{
		if len(indexs) == 1 || len(indexs) == 0{
			return true
		}else{
			if pokers[indexs[0]].CardValue == pokers[indexs[1]].CardValue{
				return true
			}else{
				return false
			}
		}
	}
}
//是否是一条龙
func IsDragon(pokers []*PokerCard) bool{
	if len(pokers) < 5 {
		return false
	}

	pokers = CommonSort(pokers)
	//如果最大的牌是王，则不是一条龙
	if pokers[len(pokers)-1].CardValue == PokerRedJoker || pokers[len(pokers)-1].CardValue == PokerRedJoker{
		return false
	}
	tempValue := -1
	for _,card := range pokers{
		if card.CardValue == tempValue+1{
			tempValue++
		}else{
			return false
		}
	}
	return true
}
//从小到大对扑克牌排序
func CommonSort(pokers []*PokerCard) []*PokerCard{
	len := len(pokers)
	for i:=0;i<len;i++{
		for j:=0;j<len-i;j++{
			if pokers[j].CardValue > pokers[j+1].CardValue{
				pokers[j+1],pokers[j] = pokers[j],pokers[j+1]
			}
		}
	}
	return pokers
}