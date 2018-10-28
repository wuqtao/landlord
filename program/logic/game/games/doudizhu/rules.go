package doudizhu

import (
	"chessSever/program/logic/game/poker"
	"errors"
	"chessSever/program/logic/game"
)

type CardsType struct {
	cardsType int
	cardMinAndMax map[string]int
}
func newCardsType(cardsType int,minValue int,maxValue int) *CardsType{
	cards := CardsType{
		cardsType,
		make(map[string]int),
	}
	cards.cardMinAndMax["min"] = minValue
	cards.cardMinAndMax["max"] = maxValue
	return &cards
}
//检查给定的扑克牌的类型
func CheckRules(currPlayerIndex int,pokers []*poker.PokerCard) (*game.LastCardsType,error){
	switch len(pokers) {
	case 0:
		return nil,errors.New("玩家出牌为空")
		//单张
	case 1:
		return game.NewLastCards(currPlayerIndex,POKERS_TYPE_SINGLE,pokers,pokers[0].CardValue,pokers[0].CardValue),nil
		//对子或者王炸
	case 2:
		if cardsType,err := CheckPair(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckJokerBomb(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//三张
	case 3:
		cardsType,err := CheckThreePlus(pokers)
		if err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//炸弹或三带一
	case 4:
		if cardsType,err := CheckCommonBomb(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckThreePlus(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//三带二或者一条龙
	case 5:
		if cardsType,err := CheckThreePlus(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckDragon(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//一条龙，或者四带二，或者四带二对
	default:
		if cardsType,err := CheckDragon(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckFourPlus(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckMultiPairs(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckMultiFourPlus(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckMultiThreePlus(pokers);err == nil{
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
	}

	return nil,nil
}

//是否是对子
func CheckPair(pokers []*poker.PokerCard) (*CardsType,error){
	if len(pokers) != 2 {
		return nil,errors.New("不是对牌")
	}

	if(pokers[0].CardValue == pokers[1].CardValue){
		return newCardsType(POKERS_TYPE_PAIR,pokers[0].CardValue,pokers[0].CardValue),nil
	}else{
		return nil,errors.New("不是对牌")
	}
}
//是否是连对
func CheckMultiPairs(pokers []*poker.PokerCard) (*CardsType,error){

	if len(pokers)%2 != 0 || len(pokers) < 6 {
		return nil,errors.New("不是连对")
	}

	poker.CommonSort(pokers)

	currValue := -1

	for i,card := range pokers{
		if i == 0 {
			currValue = card.CardValue
		}else{
			if i % 2 == 1{
				if card.CardValue != currValue{
					return nil,errors.New("不是连对")
				}
			}else{
				if card.CardValue == currValue+1{
					currValue=card.CardValue
				}else{
					return nil,errors.New("不是连对")
				}
			}
		}
	}
	return newCardsType(POKERS_TYPE_MULIT_PAIRS,pokers[0].CardValue,pokers[len(pokers)-1].CardValue),nil
}
//是否是王炸
func CheckJokerBomb(pokers []*poker.PokerCard) (*CardsType,error){

	if len(pokers) != 2{
		return nil,errors.New("不是王炸")
	}
	poker.CommonSort(pokers)
	if(pokers[0].CardValue == poker.PokerBlackJoker && pokers[1].CardValue == poker.PokerRedJoker){
		return newCardsType(POKERS_TYPE_JOKER_BOMB,pokers[0].CardValue,pokers[1].CardValue),nil
	}else{
		return nil,errors.New("不是王炸")
	}
}
//是否是炸弹
func CheckCommonBomb(pokers []*poker.PokerCard) (*CardsType,error){

	if len(pokers) != 4{
		return nil,errors.New("不是普通炸弹")
	}
	if pokers[0].CardValue == pokers[1].CardValue  && pokers[2].CardValue == pokers[3].CardValue &&
		pokers[0].CardValue == pokers[2].CardValue{
		return newCardsType(POKERS_TYPE_COMMON_BOMB,pokers[0].CardValue,pokers[0].CardValue),nil
	}else{
		return nil,errors.New("不是普通炸弹")
	}
}

//判断给定的扑克牌切片中是否存在3个一样大小的牌,以及不一样的牌的index数组,
//返回一样的三张牌数量和不同的数的索引
//调用前需要先排序
func CheckThreeSameValueCard(pokers []*poker.PokerCard) ([]int,[]int){//

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
func CheckFourSameValueCard(pokers []*poker.PokerCard) ([]int,[]int){

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
//是否是三代一或三代二或者不带
func CheckThreePlus(pokers []*poker.PokerCard) (*CardsType,error){
	if len(pokers) < 3 || len(pokers) >5{
		return nil,errors.New("不是三带牌")
	}

	poker.CommonSort(pokers)
	cardValue,indexs := CheckThreeSameValueCard(pokers)
	if len(cardValue) == 0{
		return nil,errors.New("不是三带牌")
	}else{
		if len(indexs) == 1{
			return newCardsType(POKERS_TYPE_THREE_PLUS_ONE,cardValue[0],cardValue[0]),nil
		}else if len(indexs) == 0{
			return newCardsType(POKERS_TYPE_THREE,cardValue[0],cardValue[0]),nil
		}else{
			if IsUnsameCardNumSame(pokers,indexs){
				return newCardsType(POKERS_TYPE_THREE_PLUS_TWO,cardValue[0],cardValue[0]),nil
			}else{
				return nil,errors.New("不是三带牌")
			}
		}
	}
}
//是否是四代一或者四代二
func CheckFourPlus(pokers []*poker.PokerCard) (*CardsType,error){

	if len(pokers) !=4 && len(pokers) != 6 && len(pokers) != 8{
		return nil,errors.New("不是四带牌")
	}

	poker.CommonSort(pokers)
	cardValue,indexs := CheckFourSameValueCard(pokers)
	if len(cardValue) == 0 || len(indexs) == 0 {
		return nil,errors.New("不是四带牌")
	}else{
		if len(indexs) == 2{
			return newCardsType(POKERS_TYPE_FOUR_PLUS_TWO,cardValue[0],cardValue[0]),nil
		}else if len(indexs) == 3{
			return nil,errors.New("不是四带牌")
		}else{
			if IsUnsameCardNumSame(pokers,indexs){
				return newCardsType(POKERS_TYPE_FOUR_PLUS_FOUR,cardValue[0],cardValue[0]),nil
			}else{
				return nil,errors.New("不是四带牌")
			}
		}
	}
}
//是否多个三带一，或三代二，或不带
func CheckMultiThreePlus(pokers []*poker.PokerCard) (*CardsType,error){

	if len(pokers) < 6 {
		return nil,errors.New("不是三顺")
	}

	poker.CommonSort(pokers)
	cardValues,indexs := CheckThreeSameValueCard(pokers)
	//不可能是多连
	if len(cardValues) < 2{
		return nil,errors.New("不是三顺")
	}
	//多连不带或者各带一个
	if len(indexs) == 0{
		return newCardsType(POKERS_TYPE_MULITY_THREE,cardValues[0],cardValues[len(cardValues)-1]),nil
	}else if len(cardValues) == len(indexs){
		return newCardsType(POKERS_TYPE_MULITY_THREE_PLUS_ONE,cardValues[0],cardValues[len(cardValues)-1]),nil
	}
	//不可能符合多连
	if len(indexs) != 2*len(cardValues) {
		return nil,errors.New("不是三顺")
	}
	//数量符合三代二，然后判断牌型是否符合
	if IsUnsameCardNumSame(pokers,indexs){
		return newCardsType(POKERS_TYPE_MULITY_THREE_PLUS_TWO,cardValues[0],cardValues[len(cardValues)-1]),nil
	}else{
		return nil,errors.New("不是三顺")
	}
}

//是否多个四带一或四代二，或不带
func CheckMultiFourPlus(pokers []*poker.PokerCard) (*CardsType,error){

	if len(pokers) < 8 {
		return nil,errors.New("不是四顺")
	}

	poker.CommonSort(pokers)
	cardValues,indexs := CheckFourSameValueCard(pokers)
	//四个相同牌的数量小于二的话可能是四代几，但不是多个
	if len(cardValues) < 2{
		return nil,errors.New("不是四顺")
	}
	//纯四张牌连牌
	if len(indexs) == 0{
		return newCardsType(POKERS_TYPE_MULITY_FOUR,cardValues[0],cardValues[len(cardValues)-1]),nil
	}
	//不可能符合四代二
	if len(indexs) != 2*len(cardValues) {
		return nil,errors.New("不是四顺")
	}
	//数量符合四代二，然后判断牌型是否是四代二
	if IsUnsameCardNumSame(pokers,indexs){
		if len(indexs) != 2*len(cardValues){
			return newCardsType(POKERS_TYPE_MULITY_FOUR_PLUS_TWO,cardValues[0],cardValues[len(cardValues)-1]),nil
		}else{
			return newCardsType(POKERS_TYPE_MULITY_FOUR_PLUS_FOUR,cardValues[0],cardValues[len(cardValues)-1]),nil
		}
	}else{
		return nil,errors.New("不是四顺")
	}
}
//是否是一条龙
func CheckDragon(pokers []*poker.PokerCard) (*CardsType,error){
	if len(pokers) < 5 {
		return nil,errors.New("不是顺子")
	}

	poker.CommonSort(pokers)
	//2和大小王不能参与顺子牌
	if pokers[len(pokers)-1].CardValue == poker.PokerRedJoker ||
		pokers[len(pokers)-1].CardValue == poker.PokerRedJoker ||
		pokers[len(pokers)-1].CardValue == poker.PokerTwo{
		return nil,errors.New("不是顺子")
	}
	tempValue := -1
	for i,card := range pokers{
		if i == 0 {
			tempValue = card.CardValue
		}else{
			if card.CardValue == tempValue+1{
				tempValue = card.CardValue
			}else{
				return nil,errors.New("不是顺子")
			}
		}
	}
	return newCardsType(POKERS_TYPE_DRAGON,pokers[0].CardValue,pokers[len(pokers)-1].CardValue),nil
}
//判断一组牌中，不同数字的数量是否相同
func IsUnsameCardNumSame(pokers []*poker.PokerCard,cardIndexs []int) bool{
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


