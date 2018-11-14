package set

import (
	"errors"
	"chessSever/program/game/poker/card"
)

/**
	定义扑克集合
 */
type PokerSet []*card.PokerCard
//创建新的扑克集
func NewPokerSet() PokerSet{
	return PokerSet{}
}
//向扑克集中添加扑克
func (set PokerSet)AddPokers(cards PokerSet) PokerSet{
	for _,card := range cards{
		set = append(set,card)
	}
	//append元素可能需要重新分配空间，导致原切片引用发生变化，此处需要返回新的切片引用
	return set
}
//检查给定的索引是否存在
func (set PokerSet)checkIndex(indexs []int) error{
	setLen := set.GetLength()
	for _,index := range indexs{
		if index >= setLen{
			return errors.New("给定的索引超过扑克集的长度")
		}
	}
	return nil
}
//从扑克集中删除制定索引的扑克,删除后扑克集元素变少
func (set PokerSet)DelPokersByIndex(indexs []int) (PokerSet,error){

	err := set.checkIndex(indexs)
	if err != nil{
		return nil,err
	}

	for _,index := range indexs{
		set[index] = nil
	}
	newSet := NewPokerSet()

	for _,card := range set{
		if card != nil{
			newSet = append(newSet,card)
		}
	}
	return newSet,nil
}
func (set PokerSet)DelPokers(pokers PokerSet) (PokerSet,error){
	indexs,err := set.GetPokerIndexs(pokers)
	if err != nil{
		return nil,err
	}
	newSet,err := set.DelPokersByIndex(indexs)
	return newSet,nil
}
//根据给定索引，从扑克集中获取指定扑克
func (set PokerSet)GetPokerByIndex(index int) (*card.PokerCard,error){
	err := set.checkIndex([]int{index})
	if err != nil{
		return nil,err
	}
	return set[index],nil
}
//根据给定索引，从扑克集中获取子扑克集
func (set PokerSet)GetPokersByIndexs(indexs []int) (PokerSet,error){

	err := set.checkIndex(indexs)
	if err != nil{
		return nil,err
	}

	newSet := NewPokerSet()

	for _,index := range indexs{
		newSet = append(newSet,set[index])
	}
	return newSet,nil
}
//将指定索引的扑克牌替换
func (set PokerSet)ReplacePoker(index int,card *card.PokerCard) error{
	err := set.checkIndex([]int{index})
	if err != nil{
		return err
	}
	set[index] = card
	return nil
}
//将扑克集中的各个扑克牌用于某个任务
func (set PokerSet) DoOnEachPokerCard(do func(index int,card *card.PokerCard)){
	for i,card := range set{
		do(i,card)
	}
}
//获取指定扑克牌在扑克集中的index
func (set PokerSet)GetPokerIndex(card *card.PokerCard) (int,error){
	for i,c := range set{
		if c == card{
			return i,nil
		}
	}
	return -1,errors.New("查找的扑克牌不在该扑克集中")
}
//获取指定扑克集中各扑克牌在扑克集中的index
func (set PokerSet)GetPokerIndexs(pokers PokerSet) ([]int,error){
	indexs := []int{}
	for _,card := range pokers{
		index,err := set.GetPokerIndex(card)
		if err != nil{
			return nil,err
		}
		indexs = append(indexs,index)
	}
	return indexs,nil
}

func (set PokerSet) GetLength() int{
	return len(set)
}

func (set PokerSet) GetSetTypeInfo() (*SetTypeInfo,error) {
	switch set.GetLength() {
	case 0:
		return nil,errors.New("玩家出牌为空")
		//单张
	case 1:
		return NewSetInfo(POKERS_SET_TYPE_SINGLE,set[0].GetValue(),set[0].GetValue()),nil
		//对子或者王炸
	case 2:
		if set.IsPair(){
			return NewSetInfo(POKERS_SET_TYPE_PAIR,set[0].GetValue(),set[0].GetValue()),nil
		}

		if set.IsJokerBomb(){
			return NewSetInfo(POKERS_SET_TYPE_JOKER_BOMB,set[0].GetValue(),set[1].GetValue()),nil
		}

		return nil,errors.New("牌型不符合规则")
		//三张
	case 3:
		return set.CheckThreePlus()
		//炸弹或三带一
	case 4:
		if set.IsCommonBomb(){
			return NewSetInfo(POKERS_SET_TYPE_COMMON_BOMB,set[0].GetValue(),set[0].GetValue()),nil
		}
		return set.CheckThreePlus()
		//三带二或者一条龙
	case 5:
		if set.IsDragon(){
			return NewSetInfo(POKERS_SET_TYPE_DRAGON,set[0].GetValue(),set[set.GetLength()-1].GetValue()),nil
		}
		return set.CheckThreePlus()
		//一条龙，或者四带二，或者四带二对
	default:
		if set.IsDragon(){
			return NewSetInfo(POKERS_SET_TYPE_DRAGON,set[0].GetValue(),set[set.GetLength()-1].GetValue()),nil
		}

		if set.IsMultiPair(){
			return NewSetInfo(POKERS_SET_TYPE_MULIT_PAIRS,set[0].GetValue(),set[set.GetLength()-1].GetValue()),nil
		}

		if cardsType,err := set.CheckFourPlus();err == nil{
			return cardsType,err
		}else if cardsType,err := set.CheckMultiFourPlus();err == nil{
			return cardsType,err
		}else{
			return set.CheckMultiThreePlus()
		}
	}
}

func (set PokerSet)IsPair() bool{
	if set.GetLength() != 2 {
		return false
	}

	if(set[0].CardValue == set[1].CardValue){
		return true
	}
	return false
}

func (set PokerSet)IsMultiPair() bool{
	if set.GetLength()%2 != 0 || len(set) < 6 {
		return false
	}

	set.SortAsc()
	//2和王不能作为连对出牌
	if set[set.GetLength()-1].CardValue >= card.POKER_VALUE_TWO{
		return false
	}

	currValue := -1

	for i,card := range set{
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

func (set PokerSet)IsJokerBomb() bool{
	if set.GetLength() != 2{
		return false
	}
	set.SortAsc()
	if(set[0].CardValue == card.POKER_VALUE_BLACK_JOKER && set[1].CardValue == card.POKER_VALUE_RED_JOKER){
		return true
	}else{

	}
	return false
}

func (set PokerSet)IsCommonBomb() bool{
	if set.GetLength() != 4{
		return false
	}
	if set[0].CardValue == set[1].CardValue  && set[2].CardValue == set[3].CardValue &&
		set[0].CardValue == set[2].CardValue{
		return true
	}else{
		return false
	}
}

func (set PokerSet)IsDragon() bool{
	if len(set) < 5 {
		return false
	}

	set.SortAsc()
	//2和王不能参与顺子出牌
	if set[set.GetLength()-1].CardValue >= card.POKER_VALUE_TWO{
		return false
	}

	tempValue := -1
	for i,card := range set{
		if i == 0 {
			tempValue = card.CardValue
		}else{
			if card.CardValue == tempValue+1{
				tempValue = card.CardValue
			}else{
				return false
			}
		}
	}
	return true
}

func (set PokerSet)CheckThreePlus() (*SetTypeInfo,error){
	pokersNum := set.GetLength()
	if pokersNum < 3 || pokersNum >5{
		return nil,errors.New("不是三带牌")
	}

	set.SortAsc()
	cardNum := CheckEachCardNum(set)
	cardNumCount := len(cardNum)
	if pokersNum == 3{
		if cardNumCount == 1{
			return NewSetInfo(POKERS_SET_TYPE_THREE,set[0].GetValue(),set[0].GetValue()),nil
		}else{
			return nil,errors.New("不是三带牌")
		}
	}else{
		if cardNumCount == 2 {

			for k,v := range cardNum{
				if v == 3{
					if(pokersNum == 4){
						return NewSetInfo(POKERS_SET_TYPE_THREE_PLUS_ONE,k,k),nil
					}else{
						return NewSetInfo(POKERS_SET_TYPE_THREE_PLUS_TWO,k,k),nil
					}
				}
			}
			return nil,errors.New("不是三带牌")
		}else{
			return nil,errors.New("不是三带牌")
		}
	}
}

//是否是四代一或者四代二
func (set PokerSet)CheckFourPlus() (*SetTypeInfo,error){

	pokersNum := set.GetLength()
	if pokersNum != 6 && pokersNum != 8{
		return nil,errors.New("不是四带牌")
	}

	set.SortAsc()

	cardNum := CheckEachCardNum(set)
	cardNumCount := len(cardNum)
	if cardNumCount == 2{
		for k,v := range cardNum{
			if v == 4{
				if pokersNum == 6{
					return NewSetInfo(POKERS_SET_TYPE_FOUR_PLUS_TWO,k,k),nil
				}else{
					return NewSetInfo(POKERS_SET_TYPE_FOUR_PLUS_FOUR,k,k),nil
				}
			}
		}
		return nil,errors.New("不是四带牌")
	}else{
		return nil,errors.New("不是四带牌")
	}
}

//是否多个三带一，或三代二，或不带
func (set PokerSet)CheckMultiThreePlus() (*SetTypeInfo,error){
	pokerNum := set.GetLength()
	if pokerNum < 6 {
		return nil,errors.New("不是三顺")
	}

	set.SortAsc()
	cardNum := CheckEachCardNum(set)

	mainCardValue := -1      //暂存主牌的value，用于比较是否连续
	mainCardNum := 0        //主牌的数量
	mainCardValues := []int{}  //存放主牌的值
	attachCardNum := 0    //附牌的数量
	attachCardNumMap := make(map[int]int)  //附牌的value和num的map

	for k,v := range cardNum{
		if v == 3{
			mainCardValues = append(mainCardValues,k)
			if mainCardValue == -1{
				mainCardValue = k
				mainCardNum++
			}else{
				if k == mainCardValue+1{
					mainCardValue = k
					mainCardNum++
				}else{//主牌连不起来，不能作为三顺子
					return nil,errors.New("不是三顺")
				}
			}
		}else{
			attachCardNumMap[k] = v
			attachCardNum += v
		}
	}
	//2和王不能参与连顺
	if mainCardNum > 1 && mainCardValues[len(mainCardValues)-1] > card.POKER_VALUE_ACE{
		return nil,errors.New("不是三顺")
	}

	//没有附牌
	if attachCardNum == 0{
		return NewSetInfo(POKERS_SET_TYPE_MULITY_THREE,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum == attachCardNum{//三带一
		return NewSetInfo(POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum*2 == attachCardNum{//三带二
		for _,v := range attachCardNumMap{
			if v != 2{
				return nil,errors.New("不是三顺")
			}
		}
		return NewSetInfo(POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else{
		return nil,errors.New("不是三顺")
	}
}

//是否多个四带一或四代二，或不带
func (set PokerSet)CheckMultiFourPlus() (*SetTypeInfo,error){

	pokerNum := set.GetLength()
	if pokerNum < 8 || pokerNum%2 != 0 {
		return nil,errors.New("不是四顺")
	}

	set.SortAsc()
	cardNum := CheckEachCardNum(set)

	mainCardValue := -1      //暂存主牌的value，用于比较是否连续
	mainCardNum := 0        //主牌的数量
	mainCardValues := []int{}  //存放主牌的值
	attachCardNum := 0      //附牌的数量
	attachCardNumMap := make(map[int]int)  //附牌的value和num的map

	for k,v := range cardNum{
		if v == 4{
			mainCardValues = append(mainCardValues,k)
			if mainCardValue == -1 {
				mainCardValue = k
				mainCardNum++
			}else{
				if k == mainCardValue+1{
					mainCardValue = k
					mainCardNum++
				}else{//主牌连不起来，不能作为三顺子
					return nil,errors.New("不是四顺")
				}
			}
		}else{
			attachCardNumMap[k] = v
			attachCardNum += v
		}
	}

	//2和王不能参与连顺
	if mainCardNum > 1 && mainCardValues[len(mainCardValues)-1] > card.POKER_VALUE_ACE{
		return nil,errors.New("不是四顺")
	}

	//没有附牌
	if attachCardNum == 0{//四不带
		return NewSetInfo(POKERS_SET_TYPE_MULITY_FOUR,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum*2 == attachCardNum{//四带二
		return NewSetInfo(POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum*4 == attachCardNum{//四带四
		for _,v := range attachCardNumMap{
			if v != 2{
				return nil,errors.New("不是四顺")
			}
		}
		return NewSetInfo(POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else{
		return nil,errors.New("不是四顺")
	}
}

func (set PokerSet)SortAsc(){
	SortAsc(set)
}

func (set PokerSet)SortDesc(){
	SortDesc(set)
}
//检测是否有相同元素，有则冲突
func (set PokerSet)HasSamePoker(s PokerSet) bool{
	for _,card1 := range set{
		for _,card2 := range s{
			if card1 == card2{
				return true
			}
		}
	}
	return false
}
