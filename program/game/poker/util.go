package poker

//判断给定的扑克牌切片中是否存在3个一样大小的牌,以及不一样的牌的index数组,
//返回一样的三张牌数量和不同的数的索引
//调用前需要先排序
func CheckThreeSameValueCard(pokers []*PokerCard) ([]int,[]int){

	if len(pokers) < 3{
		return nil,nil
	}
	cardValue := []int{}  //主牌的value值切片，同一值只存在一个
	indexs := []int{}     //非主牌的索引
	num := 0
	currValue := -1
	for i,card := range pokers{
		if i == 0 {
			currValue = card.CardValue
			num++
		}else{
			if card.CardValue == currValue{
				num++
				if(num == 3){
					cardValue = append(cardValue,currValue)
				}
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
					currValue = card.CardValue
					num = 1
				}
				//最后一张的话没有后续的牌进行参照了，直接加入
				if i == len(pokers)-1{
					indexs = append(indexs,i)
				}
			}
		}
	}

	return cardValue,indexs
}

//判断给定的扑克牌切片中是否存在4个一样大小的牌,以及不一样的牌的index数组
func CheckFourSameValueCard(pokers []*PokerCard) ([]int,[]int){

	if len(pokers) < 4{
		return nil,nil
	}
	cardValue := []int{}
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
				if num == 4 {
					cardValue = append(cardValue,currValue)
				}
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
					currValue = card.CardValue
					num = 1
				}
				//最后一张的话没有后续的牌进行参照了，直接加入
				if i == len(pokers)-1{
					indexs = append(indexs,i)
				}
			}
		}
	}

	return cardValue,indexs
}

//判断一组牌中，不同数字的数量是否相同
func IsUnsameCardNumSame(pokers []*PokerCard,cardIndexs []int) bool{
	numMap := make(map[int]int)
	for _,index := range cardIndexs{
		_,ok := numMap[pokers[index].CardValue]
		if ok {
			numMap[pokers[index].CardValue]++
		}else{
			numMap[pokers[index].CardValue] = 1
		}
	}
	temp := 0
	index := 1
	for _,v:= range numMap{
		if index == 1{
			temp = v
			index++
		}else{
			if temp != v{
				return false
			}
		}
	}
	temp++ //消除，定义未引用错误
	return true
}
//判断一组牌中，给定索引的牌是否一样大小
func IsCardSame(pokers []*PokerCard,cardIndexs []int) bool{
	temp := -1
	for i,v:= range cardIndexs{
		if i == 0{
			temp = v
		}else{
			if temp != v{
				return false
			}
		}
	}
	return true
}

