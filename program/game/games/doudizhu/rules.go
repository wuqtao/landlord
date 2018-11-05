package doudizhu

import (
	"chessSever/program/game"
	"errors"
	"chessSever/program/game/poker"
)
//临时类型，便于函数返回类型
type subCardsType struct {
	cardsType int
	cardMinAndMax map[string]int
}
func newCardsType(cardsType int,minValue int,maxValue int) *subCardsType {
	cards := subCardsType{
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
		return game.NewLastCards(currPlayerIndex,game.POKERS_TYPE_SINGLE,pokers,pokers[0].CardValue,pokers[0].CardValue),nil
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
func CheckPair(pokers []*poker.PokerCard) (*subCardsType,error){
	if len(pokers) != 2 {
		return nil,errors.New("不是对牌")
	}

	if(pokers[0].CardValue == pokers[1].CardValue){
		return newCardsType(game.POKERS_TYPE_PAIR,pokers[0].CardValue,pokers[0].CardValue),nil
	}else{
		return nil,errors.New("不是对牌")
	}
}
//是否是连对
func CheckMultiPairs(pokers []*poker.PokerCard) (*subCardsType,error){

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
	return newCardsType(game.POKERS_TYPE_MULIT_PAIRS,pokers[0].CardValue,pokers[len(pokers)-1].CardValue),nil
}
//是否是王炸
func CheckJokerBomb(pokers []*poker.PokerCard) (*subCardsType,error){

	if len(pokers) != 2{
		return nil,errors.New("不是王炸")
	}
	poker.CommonSort(pokers)
	if(pokers[0].CardValue == poker.PokerBlackJoker && pokers[1].CardValue == poker.PokerRedJoker){
		return newCardsType(game.POKERS_TYPE_JOKER_BOMB,pokers[0].CardValue,pokers[1].CardValue),nil
	}else{
		return nil,errors.New("不是王炸")
	}
}
//是否是炸弹
func CheckCommonBomb(pokers []*poker.PokerCard) (*subCardsType,error){

	if len(pokers) != 4{
		return nil,errors.New("不是普通炸弹")
	}
	if pokers[0].CardValue == pokers[1].CardValue  && pokers[2].CardValue == pokers[3].CardValue &&
		pokers[0].CardValue == pokers[2].CardValue{
		return newCardsType(game.POKERS_TYPE_COMMON_BOMB,pokers[0].CardValue,pokers[0].CardValue),nil
	}else{
		return nil,errors.New("不是普通炸弹")
	}
}

//是否是三代一或三代二或者不带
func CheckThreePlus(pokers []*poker.PokerCard) (*subCardsType,error){
	if len(pokers) < 3 || len(pokers) >5{
		return nil,errors.New("不是三带牌")
	}

	poker.CommonSort(pokers)
	cardValue,indexs := poker.CheckThreeSameValueCard(pokers)
	if len(cardValue) == 0{
		return nil,errors.New("不是三带牌")
	}else{
		if len(indexs) == 1{
			return newCardsType(game.POKERS_TYPE_THREE_PLUS_ONE,cardValue[0],cardValue[0]),nil
		}else if len(indexs) == 0{
			return newCardsType(game.POKERS_TYPE_THREE,cardValue[0],cardValue[0]),nil
		}else{
			if poker.IsUnsameCardNumSame(pokers,indexs){
				return newCardsType(game.POKERS_TYPE_THREE_PLUS_TWO,cardValue[0],cardValue[0]),nil
			}else{
				return nil,errors.New("不是三带牌")
			}
		}
	}
}
//是否是四代一或者四代二
func CheckFourPlus(pokers []*poker.PokerCard) (*subCardsType,error){

	if len(pokers) !=4 && len(pokers) != 6 && len(pokers) != 8{
		return nil,errors.New("不是四带牌")
	}

	poker.CommonSort(pokers)
	cardValue,indexs := poker.CheckFourSameValueCard(pokers)
	if len(cardValue) == 0 || len(indexs) == 0 {
		return nil,errors.New("不是四带牌")
	}else{
		if len(indexs) == 2{
			return newCardsType(game.POKERS_TYPE_FOUR_PLUS_TWO,cardValue[0],cardValue[0]),nil
		}else if len(indexs) == 3{
			return nil,errors.New("不是四带牌")
		}else{
			if poker.IsUnsameCardNumSame(pokers,indexs){
				return newCardsType(game.POKERS_TYPE_FOUR_PLUS_FOUR,cardValue[0],cardValue[0]),nil
			}else{
				return nil,errors.New("不是四带牌")
			}
		}
	}
}
//是否多个三带一，或三代二，或不带
func CheckMultiThreePlus(pokers []*poker.PokerCard) (*subCardsType,error){

	if len(pokers) < 6 {
		return nil,errors.New("不是三顺")
	}

	poker.CommonSort(pokers)
	cardValues,indexs := poker.CheckThreeSameValueCard(pokers)
	//不可能是多连
	if len(cardValues) < 2{
		return nil,errors.New("不是三顺")
	}
	//多连不带或者各带一个
	if len(indexs) == 0{
		return newCardsType(game.POKERS_TYPE_MULITY_THREE,cardValues[0],cardValues[len(cardValues)-1]),nil
	}else if len(cardValues) == len(indexs){
		return newCardsType(game.POKERS_TYPE_MULITY_THREE_PLUS_ONE,cardValues[0],cardValues[len(cardValues)-1]),nil
	}
	//不可能符合多连
	if len(indexs) != 2*len(cardValues) {
		return nil,errors.New("不是三顺")
	}
	//数量符合三代二，然后判断牌型是否符合
	if poker.IsUnsameCardNumSame(pokers,indexs){
		return newCardsType(game.POKERS_TYPE_MULITY_THREE_PLUS_TWO,cardValues[0],cardValues[len(cardValues)-1]),nil
	}else{
		return nil,errors.New("不是三顺")
	}
}

//是否多个四带一或四代二，或不带
func CheckMultiFourPlus(pokers []*poker.PokerCard) (*subCardsType,error){

	if len(pokers) < 8 {
		return nil,errors.New("不是四顺")
	}

	poker.CommonSort(pokers)
	cardValues,indexs := poker.CheckFourSameValueCard(pokers)
	//四个相同牌的数量小于二的话可能是四代几，但不是多个
	if len(cardValues) < 2{
		return nil,errors.New("不是四顺")
	}
	//纯四张牌连牌
	if len(indexs) == 0{
		return newCardsType(game.POKERS_TYPE_MULITY_FOUR,cardValues[0],cardValues[len(cardValues)-1]),nil
	}
	//不可能符合四代二
	if len(indexs) != 2*len(cardValues) {
		return nil,errors.New("不是四顺")
	}
	//数量符合四代二，然后判断牌型是否是四代二
	if poker.IsUnsameCardNumSame(pokers,indexs){
		if len(indexs) != 2*len(cardValues){
			return newCardsType(game.POKERS_TYPE_MULITY_FOUR_PLUS_TWO,cardValues[0],cardValues[len(cardValues)-1]),nil
		}else{
			return newCardsType(game.POKERS_TYPE_MULITY_FOUR_PLUS_FOUR,cardValues[0],cardValues[len(cardValues)-1]),nil
		}
	}else{
		return nil,errors.New("不是四顺")
	}
}
//是否是一条龙
func CheckDragon(pokers []*poker.PokerCard) (*subCardsType,error){
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
	return newCardsType(game.POKERS_TYPE_DRAGON,pokers[0].CardValue,pokers[len(pokers)-1].CardValue),nil
}


