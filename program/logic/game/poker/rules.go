package poker

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
//是否是连对
func IsMultiPairs(pokers []*PokerCard) bool{

	if len(pokers)%2 != 0 || len(pokers) < 6 {
		return false
	}

	pokers = CommonSort(pokers)

	currValue := -1

	for i,card := range pokers{
		if i == 0 {
			currValue = card.CardValue
		}else{
			if i % 2 == 1{
				if card.CardValue != currValue{
					return false
				}
			}else{
				if card.CardValue == currValue+1{
					currValue=card.CardValue
				}else{
					return false
				}
			}
		}
	}
	return true
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

//判断给定的扑克牌切片中是否存在3个一样大小的牌,以及不一样的牌的index数组,
//返回一样的三张牌数量和不同的数的索引
//调用前需要先排序
func CheckThreeSameValueCard(pokers []*PokerCard) (int,[]int){

	if len(pokers) < 3{
		return 0,nil
	}
	sameNum := 0
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
				case 3:
					indexs = append(indexs,i)
					num = 1
					sameNum++
					currValue = card.CardValue
					num = 1
				}
			}
		}
	}

	if num >= 3{
		return sameNum,indexs
	}else{
		return sameNum,indexs
	}
}

//判断给定的扑克牌切片中是否存在4个一样大小的牌,以及不一样的牌的index数组
func CheckFourSameValueCard(pokers []*PokerCard) (int,[]int){

	if len(pokers) < 4{
		return 0,nil
	}

	sameNum := 0
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
					indexs = append(indexs, i-1)
					currValue = card.CardValue
					num = 1
				case 2:
					indexs = append(indexs, i-1)
					indexs = append(indexs, i-2)
					currValue = card.CardValue
					num = 1
				case 3:
					indexs = append(indexs, i-1)
					indexs = append(indexs, i-2)
					indexs = append(indexs, i-3)
					currValue = card.CardValue
					num = 1
				case 4:
					indexs = append(indexs, i)
					sameNum++
					currValue = card.CardValue
					num = 1
				}
			}
		}
	}

	if num >= 4{
		return sameNum,indexs
	}else{
		return sameNum,indexs
	}
}
//是否是三代一或三代二
func IsThreePlus(pokers []*PokerCard) bool{
	if len(pokers) < 3 || len(pokers) >5{
		return false
	}

	pokers = CommonSort(pokers)
	num,indexs := CheckThreeSameValueCard(pokers)
	if num !=1{
		return false
	}else{
		if len(indexs) == 1 || len(indexs) == 0{
			return true
		}else{
			num := checkUnsameNum(pokers,indexs)
			if num == 1{
				return true
			}else{
				return false
			}
		}
	}
}
//是否是四代一或者四代二
func IsFourPlus(pokers []*PokerCard) bool{

	if len(pokers) !=4 && len(pokers) != 6 && len(pokers) != 8{
		return false
	}

	pokers = CommonSort(pokers)
	num,indexs := CheckFourSameValueCard(pokers)
	if num != 1{
		return false
	}else{
		if len(indexs) == 2 || len(indexs) == 0{
			return true
		}else{
			if len(indexs) == 3 {
				return false
			}else{
				num := checkUnsameNum(pokers,indexs)
				if num == 1{
					return true
				}else{
					return false
				}
			}
		}
	}
	return true
}
//是否多个三带一，或三代二，或不带
func IsMultiThreePlus(pokers []*PokerCard) bool{

	if len(pokers) < 6 {
		return false
	}

	pokers = CommonSort(pokers)
	num,indexs := CheckThreeSameValueCard(pokers)
	//不可能是多连
	if num < 2{
		return false
	}
	//多连不带或者各带一个
	if num == len(indexs) || len(indexs) == 0{
		return true
	}
	//不可能符合多连
	if len(indexs) != 2*num {
		return false
	}
	//数量符合三代二，然后判断牌型是否符合
	unsameNum := checkUnsameNum(pokers,indexs)

	if unsameNum == num{
		return true
	}else{
		return false
	}
}

//是否多个四带一或四代二，或不带
func IsMultiFourPlus(pokers []*PokerCard) bool{

	if len(pokers) < 8 {
		return false
	}

	pokers = CommonSort(pokers)
	num,indexs := CheckFourSameValueCard(pokers)
	//四个相同牌的数量小于二的话可能是四代几，但不是多个
	if num < 2{
		return false
	}
	//蠢四张牌连牌，或者四张连拍各带一个
	if num == len(indexs) || len(indexs) == 0{
		return true
	}
	//不可能符合四代二
	if len(indexs) != 2*num {
		return false
	}
	//数量符合四代二，然后判断牌型是否是四代二
	unsameNum := checkUnsameNum(pokers,indexs)

	if unsameNum == num{
		return true
	}else{
		return false
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
//判断一组牌中，不同数字的数量
func checkUnsameNum(pokers []*PokerCard,cardIndexs []int) int{
	unsame := []int{}
	for _,index := range cardIndexs{
		for i,value := range unsame{
			if pokers[index].CardValue == value{
				break
			}else{
				if i == len(unsame)-1{
					unsame = append(unsame,pokers[index].CardValue)
				}
			}
		}
	}
	return len(unsame)
}
