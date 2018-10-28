package poker

//判断给定的扑克牌切片中是否存在3个一样大小的牌,以及不一样的牌的index数组,
//返回一样的三张牌数量和不同的数的索引
//调用前需要先排序
func CheckThreeSameValueCard(pokers []*PokerCard) ([]int,[]int){

	if len(pokers) < 3{
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
					indexs = append(indexs,i)
					currValue = card.CardValue
					num = 1
				}
			}
		}
	}

	if num >= 3{
		return cardValue,indexs
	}else{
		return cardValue,indexs
	}
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
			if num == 4 {
				cardValue = append(cardValue,currValue)
			}
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
					currValue = card.CardValue
					num = 1
				}
			}
		}
	}

	if num >= 4{
		return cardValue,indexs
	}else{
		return cardValue,indexs
	}
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
			return false
		}
	}
	temp++ //消除，定义未引用错误
	return true
}

