package analyzer

import (
	"chessSever/program/game/poker/set"
	"chessSever/program/game/poker/card"
	"sync"
	"github.com/ant0ine/go-json-rest/rest"
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

	useableSets := []set.PokerSet{}
	//上一次出牌不是炸弹，则直接将炸弹加入可出的排中
	if setType.SetType != set.POKERS_SET_TYPE_COMMON_BOMB && setType.SetType != set.POKERS_SET_TYPE_JOKER_BOMB{
		useableSets = append(useableSets,ana.GetJokerBomb())
		for _,tempSet := range ana.getSingleValueSet(4,-1){
			useableSets = append(useableSets,tempSet)
		}
	}

	switch setType.SetType{
		case set.POKERS_SET_TYPE_SINGLE:
			for _,tempset := range ana.getSingleValueSet(1,setType.CardValueMinAndMax["min"]){
				useableSets = append(useableSets,tempset)
			}
			return useableSets
		case set.POKERS_SET_TYPE_DRAGON:
			for _,tempset := range ana.getMultiValueSet(1,setType.CardValueMinAndMax["min"],setType.CardValueMinAndMax["max"]){
				useableSets = append(useableSets,tempset)
			}
			return useableSets
		case set.POKERS_SET_TYPE_PAIR:
			for _,tempset := range ana.getSingleValueSet(2,setType.CardValueMinAndMax["min"]){
				useableSets = append(useableSets,tempset)
			}
			return useableSets
		case set.POKERS_SET_TYPE_MULIT_PAIRS:
			for _,tempset := range ana.getMultiValueSet(2,setType.CardValueMinAndMax["min"],setType.CardValueMinAndMax["max"]){
				useableSets = append(useableSets,tempset)
			}
			return useableSets
		case set.POKERS_SET_TYPE_THREE:
			for _,tempset := range ana.getSingleValueSet(3,setType.CardValueMinAndMax["min"]){
				useableSets = append(useableSets,tempset)
			}
			return useableSets
		case set.POKERS_SET_TYPE_THREE_PLUS_ONE:
			tempSets := ana.getSingleValueSet(3,setType.CardValueMinAndMax["min"])
			for _,tempset := range tempSets{
				tempsetPlus := ana.getPlusSet(1,setType.CardValueMinAndMax["max"]-setType.CardValueMinAndMax["min"]+1,tempset)
				if tempsetPlus.GetLength() >0 {
					tempset = tempset.AddPokers(tempsetPlus)
					useableSets = append(useableSets,tempset)
				}
			}
			return useableSets
		case set.POKERS_SET_TYPE_THREE_PLUS_TWO:
			tempSets := ana.getSingleValueSet(3,setType.CardValueMinAndMax["min"])
			for _,tempset := range tempSets{
				tempsetPlus := ana.getPlusSet(2,setType.CardValueMinAndMax["max"]-setType.CardValueMinAndMax["min"]+1,tempset)
				if tempsetPlus.GetLength() >0 {
					tempset = tempset.AddPokers(tempsetPlus)
					useableSets = append(useableSets,tempset)
				}
			}
			return useableSets
		case set.POKERS_SET_TYPE_MULITY_THREE:
			return ana.getMultiValueSet(3,setType.CardValueMinAndMax["min"],setType.CardValueMinAndMax["max"])
		case set.POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE:
			return "三顺子带一张"
		case set.POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO:
			return "三顺子带两张"
		case set.POKERS_SET_TYPE_FOUR_PLUS_TWO:
			tempSets := ana.getSingleValueSet(4,setType.CardValueMinAndMax["min"])
			for _,tempset := range tempSets{
				tempsetPlus := ana.getPlusSet(2,setType.CardValueMinAndMax["max"]-setType.CardValueMinAndMax["min"]+1,tempset)
				if tempsetPlus.GetLength() >0 {
					tempset = tempset.AddPokers(tempsetPlus)
					useableSets = append(useableSets,tempset)
				}
			}
			return useableSets
		case set.POKERS_SET_TYPE_FOUR_PLUS_FOUR:
			return "四带四"
		case set.POKERS_SET_TYPE_MULITY_FOUR:
			return ana.getMultiValueSet(4,setType.CardValueMinAndMax["min"],setType.CardValueMinAndMax["max"])
		case set.POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO:
			return "四顺子带两张"
		case set.POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR:
			return "四顺子带四张"
		case set.POKERS_SET_TYPE_COMMON_BOMB:
			return "炸弹"
		case set.POKERS_SET_TYPE_JOKER_BOMB:
			return []set.PokerSet{}
		default:
			return []set.PokerSet{}
	}
}
//获取单值牌组成的扑克集的切片
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
//获取多种不同值组成的扑克集的切片
func (ana PokerAnalyzer) getMultiValueSet(count int,minValue int,maxValue int) []set.PokerSet{
	sets := []set.PokerSet{}
	se := set.NewPokerSet()
	for i:=minValue+1;i<=card.POKER_VALUE_TWO-(maxValue-minValue+1);i++{
		for j:=i;j<i+maxValue-minValue+1;j++{
			if ana.Dic[i].GetLength() >= count {
				se.AddPokers(ana.Dic[i][:count])
			}
		}
		if se.GetLength() == maxValue-minValue+1{
			sets = append(sets,se)
			se =  set.NewPokerSet()
		}else{
			se =  set.NewPokerSet()
		}
	}
	return sets
}
//不拆牌为第一原则，可能会带出去大牌
func (ana PokerAnalyzer) getPlusSet(num int,count int,exceptSet set.PokerSet) set.PokerSet{
	resSet := set.NewPokerSet()
	//第一原则不拆牌原则
	for i:=card.POKER_VALUE_THREE;i<= card.POKER_VALUE_RED_JOKER;i++{
		set,_ := ana.Dic[i]
		if set.GetLength() >= num{
			if !set[:num].HasSamePoker(exceptSet) {
				resSet = resSet.AddPokers(set[:num])
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
		set, _ := ana.Dic[i]
		if set.GetLength() > 0 {
			resSet = resSet.AddPokers(set)
		}
	}
	if resSet.GetLength() > 1{
		return resSet
	}else{
		return set.NewPokerSet()
	}
}


