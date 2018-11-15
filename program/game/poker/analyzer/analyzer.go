package analyzer

import (
	"chessSever/program/game/poker/set"
	"chessSever/program/game/poker/card"
	"sync"
)

//定义玩家的扑克牌分析器map的索引为poker的value,value为改值得扑克牌在玩家牌中的索引
type PokerAnalyzer struct{
	sync.RWMutex
	Dic map[int]set.PokerSet
}
//创建新的扑克牌分析器
func NewPokerAnalyzer() PokerAnalyzer{
	analyzer := PokerAnalyzer{
		Dic:make(map[int]set.PokerSet),
	}
	return analyzer
}
//根据给定的扑克集初始化分析器
func (ana PokerAnalyzer) InitAnalyzer(){
	ana.Dic[card.POKER_VALUE_THREE] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_FOUR] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_FIVE] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_SIX] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_SEVEN] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_EIGHT] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_NINE] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_TEN] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_JACK] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_QUEEN] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_KING] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_ACE] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_TWO] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_BLACK_JOKER] = set.PokerSet{}
	ana.Dic[card.POKER_VALUE_RED_JOKER] = set.PokerSet{}
}
//根据给定的扑克集更新记牌器,出牌时调用
func (ana PokerAnalyzer) RemovePokerSet(pokers set.PokerSet){
	pokers.DoOnEachPokerCard(func(index int,card *card.PokerCard){
		ana.Dic[card.CardValue],_ = ana.Dic[card.CardValue].DelPokers(set.PokerSet{card})
	})
}

func (ana PokerAnalyzer) AddPokerSet(pokers set.PokerSet){
	pokers.DoOnEachPokerCard(func(index int,card *card.PokerCard){
		ana.Dic[card.CardValue] = ana.Dic[card.CardValue].AddPokers(set.PokerSet{card})
	})
}

func (ana PokerAnalyzer) GetMinPlayableCards() set.PokerSet{
	for i:=card.POKER_VALUE_THREE;i<=card.POKER_VALUE_RED_JOKER;i++{
		set,_ := ana.Dic[i]
		if set.GetLength() > 0{
			return set
		}
	}
	return set.PokerSet{}
}
//根据最后一次出牌的牌型信息，返回可出的扑克集
func (ana PokerAnalyzer) GetUseableCards(setType *set.SetTypeInfo) []set.PokerSet{

	var useableSets []set.PokerSet

	switch setType.SetType{
		case set.POKERS_SET_TYPE_SINGLE:
			useableSets = ana.getSingleValueSet(1,setType.GetMinValue())
		case set.POKERS_SET_TYPE_DRAGON:
			useableSets = ana.getMultiValueSet(1,setType.GetMinValue(),setType.GetMaxValue())
		case set.POKERS_SET_TYPE_PAIR:
			useableSets = ana.getSingleValueSet(2,setType.GetMinValue())
		case set.POKERS_SET_TYPE_MULIT_PAIRS:
			useableSets = ana.getMultiValueSet(2,setType.GetMinValue(),setType.GetMaxValue())
		case set.POKERS_SET_TYPE_THREE:
			useableSets = ana.getSingleValueSet(3,setType.GetMinValue())
		case set.POKERS_SET_TYPE_THREE_PLUS_ONE:
			useableSets = ana.getSingleValueSet(3,setType.GetMinValue())
			for i,tempset := range useableSets{
				tempsetPlus := ana.getPlusSet(1,1,tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{//没有牌可以带，将之前的主牌移除可出牌集合
					useableSets[i] = nil
				}
			}
		case set.POKERS_SET_TYPE_THREE_PLUS_TWO:
			useableSets = ana.getSingleValueSet(3,setType.GetMinValue())
			for i,tempset := range useableSets{
				tempsetPlus := ana.getPlusSet(2,1,tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{//没有牌可以带，将之前的主牌移除可出牌集合
					useableSets[i] = nil
				}
			}
		case set.POKERS_SET_TYPE_MULITY_THREE:
			useableSets = ana.getMultiValueSet(3,setType.GetMinValue(),setType.GetMaxValue())
		case set.POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE:
			useableSets = ana.getMultiValueSet(3,setType.GetMinValue(),setType.GetMaxValue())
			for i,tempset := range useableSets{
				tempsetPlus := ana.getPlusSet(1,setType.GetRangeWidth(),tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{//没有牌可以带，将之前的主牌移除可出牌集合
					useableSets[i] = nil
				}
			}
		case set.POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO:
			useableSets = ana.getMultiValueSet(3,setType.GetMinValue(),setType.GetMaxValue())
			for i,tempset := range useableSets{
				tempsetPlus := ana.getPlusSet(2,setType.GetRangeWidth(),tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{//没有牌可以带，将之前的主牌移除可出牌集合
					useableSets[i] = nil
				}
			}
		case set.POKERS_SET_TYPE_FOUR_PLUS_TWO:
			useableSets = ana.getSingleValueSet(4,setType.GetMinValue())
			for i,tempset := range useableSets{
				//带两个单牌
				tempsetPlus := ana.getPlusSet(1,2,tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{
					//带一对牌，看做两个单牌
					tempsetPlus := ana.getPlusSet(2,1,tempset)
					if tempsetPlus.GetLength() >0 {
						useableSets[i] = tempset.AddPokers(tempsetPlus)
					}else{//没有牌可以带，将之前的主牌移除可出牌集合
						useableSets[i] = nil
					}
				}
			}
		case set.POKERS_SET_TYPE_FOUR_PLUS_FOUR:
			useableSets = ana.getSingleValueSet(4,setType.GetMinValue())
			for i,tempset := range useableSets{
				tempsetPlus := ana.getPlusSet(2,2,tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{//没有牌可以带，将之前的主牌移除可出牌集合
					useableSets[i] = nil
				}
			}
		case set.POKERS_SET_TYPE_MULITY_FOUR:
			useableSets = ana.getMultiValueSet(4,setType.GetMinValue(),setType.GetMaxValue())
		case set.POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO:
			useableSets = ana.getMultiValueSet(4,setType.GetMinValue(),setType.GetMaxValue())
			for i,tempset := range useableSets{
				//带两个单牌
				tempsetPlus := ana.getPlusSet(1,2*setType.GetRangeWidth(),tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{
					//带一对牌，看做两个单牌
					tempsetPlus := ana.getPlusSet(2,setType.GetRangeWidth(),tempset)
					if tempsetPlus.GetLength() >0 {
						useableSets[i] = tempset.AddPokers(tempsetPlus)
					}else{//没有牌可以带，将之前的主牌移除可出牌集合
						useableSets[i] = nil
					}
				}
			}
		case set.POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR:
			useableSets = ana.getMultiValueSet(4,setType.GetMinValue(),setType.GetMaxValue())
			for i,tempset := range useableSets{
				tempsetPlus := ana.getPlusSet(2,2*setType.GetRangeWidth(),tempset)
				if tempsetPlus.GetLength() >0 {
					useableSets[i] = tempset.AddPokers(tempsetPlus)
				}else{//没有牌可以带，将之前的主牌移除可出牌集合
					useableSets[i] = nil
				}
			}
		case set.POKERS_SET_TYPE_COMMON_BOMB:
			useableSets = ana.getSingleValueSet(4,setType.GetMinValue())
		case set.POKERS_SET_TYPE_JOKER_BOMB:
			useableSets = []set.PokerSet{}
		default:
			useableSets = []set.PokerSet{}
	}
	//上一次出牌不是炸弹，则直接将炸弹加入可出的排中
	if setType.SetType != set.POKERS_SET_TYPE_COMMON_BOMB && setType.SetType != set.POKERS_SET_TYPE_JOKER_BOMB{
		//王炸
		jokerBombSet := ana.GetJokerBomb()
		if jokerBombSet.GetLength() > 0{
			useableSets = append(useableSets,jokerBombSet)
		}
		//普通炸弹
		for _,tempSet := range ana.getSingleValueSet(4,-1){
			if tempSet.GetLength() > 0{
				useableSets = append(useableSets,tempSet)
			}
		}
	}
	return useableSets
}
//获取单值牌组成的扑克集的切片，单排对牌三牌四排等等
//count表示单值牌的张数
//minValue表示上家出牌的最小的牌的大小
func (ana PokerAnalyzer) getSingleValueSet(count int,minValue int) []set.PokerSet{
	sets := []set.PokerSet{}
	se := set.NewPokerSet()
	for i:=minValue+1;i<=card.POKER_VALUE_RED_JOKER;i++{
		if ana.Dic[i].GetLength() >= count {
			se = se.AddPokers(ana.Dic[i][:count])
			sets = append(sets,se)
			se =  set.NewPokerSet()
		}
	}
	return sets
}
//获取多种不同值组成的扑克集的切片,2连3连4连5连等
func (ana PokerAnalyzer) getMultiValueSet(count int,minValue int,maxValue int) []set.PokerSet{
	sets := []set.PokerSet{}
	se := set.NewPokerSet()
	valueRange := maxValue-minValue+1
	for i:=minValue+1;i<=card.POKER_VALUE_TWO-valueRange;i++{
		for j:=i;j<i+valueRange;j++{
			if ana.Dic[j].GetLength() >= count {
				se = se.AddPokers(ana.Dic[j][:count])
			}
		}
		//该范围内连续的牌的张数符合要求
		if se.GetLength() == valueRange*count{
			sets = append(sets,se)
			se =  set.NewPokerSet()
		}else{
			se =  set.NewPokerSet()
		}
	}
	return sets
}
//获取附牌，比如三带一中的一，四带二中二，只获取一种可能即可
//不拆牌为第一原则，可能会带出去大牌
//num张数count系列数exceptset不能包含在内的扑克集
func (ana PokerAnalyzer) getPlusSet(num int,count int,exceptSet set.PokerSet) set.PokerSet{
	resSet := set.NewPokerSet()
	//第一原则不拆牌原则
	for i:=card.POKER_VALUE_THREE;i<= card.POKER_VALUE_RED_JOKER;i++{
		if ana.Dic[i].GetLength() >= num{
			if !ana.Dic[i][:num].HasSamePoker(exceptSet) {
				resSet = resSet.AddPokers(ana.Dic[i][:num])
			}
		}
		if resSet.GetLength() == num*count{
			return resSet
		}
	}
	return set.PokerSet{}
}
func (ana PokerAnalyzer) GetJokerBomb() set.PokerSet{
	resSet := set.NewPokerSet()
	for i:=card.POKER_VALUE_BLACK_JOKER;i<= card.POKER_VALUE_RED_JOKER;i++ {
		if ana.Dic[i].GetLength() > 0 {
			resSet = resSet.AddPokers(ana.Dic[i])
		}
	}
	if resSet.GetLength() > 1{
		return resSet
	}else{
		return set.NewPokerSet()
	}
}


