package poker

import "iris/core/errors"

/**
	定义扑克集合
 */
type PokerSet []*PokerCard
//创建新的扑克集
func NewPokerSet() PokerSet{
	return PokerSet{}
}
//向扑克集中添加扑克
func (set PokerSet)AddPokers(cards PokerSet){
	for _,card := range cards{
		set = append(set,card)
	}
}
//检查给定的索引是否存在
func (set PokerSet)checkIndex(indexs []int) error{
	setLen := set.GetLenth()
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
func (set PokerSet)GetPokerByIndex(index int) (*PokerCard,error){
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
func (set PokerSet)ReplacePoker(index int,card *PokerCard) error{
	err := set.checkIndex([]int{index})
	if err != nil{
		return err
	}
	set[index] = card
	return nil
}
//将扑克集中的各个扑克牌用于某个任务
func (set PokerSet) DoOnEachPokerPoker(do func(index int,card *PokerCard)){
	for i,card := range set{
		do(i,card)
	}
}
//获取指定扑克牌在扑克集中的index
func (set PokerSet)GetPokerIndex(card *PokerCard) (int,error){
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

func (set PokerSet)GetLenth() int{
	return len(set)
}



