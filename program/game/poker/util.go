package poker

//检查一组牌中，各牌的数量,返回map[cardValue]num
func CheckEachCardNum(pokers []*PokerCard) map[int]int{
	cardMap := make(map[int]int)

	if len(pokers) == 0 {
		return nil
	}

	for _,card := range pokers {
		_,ok := cardMap[card.CardValue]
		if ok {
			cardMap[card.CardValue]++
		}else{
			cardMap[card.CardValue] = 1
		}
	}

	return cardMap
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
			temp = pokers[v].CardValue
		}else{
			if temp != pokers[v].CardValue{
				return false
			}
		}
	}
	return true
}

