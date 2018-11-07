package doudizhu

import (
	"chessSever/program/game"
	"errors"
	"chessSever/program/game/poker"
	"fmt"
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
			fmt.Println(cardsType)
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckJokerBomb(pokers);err == nil{
			fmt.Println(cardsType)
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//三张
	case 3:
		cardsType,err := CheckThreePlus(pokers)
		if err == nil{
			fmt.Println(cardsType)
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//炸弹或三带一
	case 4:
		if cardsType,err := CheckCommonBomb(pokers);err == nil{
			fmt.Println(cardsType)
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckThreePlus(pokers);err == nil{
			fmt.Println(cardsType)
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else{
			return nil,errors.New("牌型不符合规则")
		}
		//三带二或者一条龙
	case 5:
		if cardsType,err := CheckThreePlus(pokers);err == nil{
			fmt.Println(cardsType)
			return game.NewLastCards(currPlayerIndex,cardsType.cardsType,pokers,cardsType.cardMinAndMax["min"],cardsType.cardMinAndMax["max"]),nil
		}else if cardsType,err := CheckDragon(pokers);err == nil{
			fmt.Println(cardsType)
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
	//2和王不能作为连对出牌
	if pokers[len(pokers)-1].CardValue >= poker.PokerTwo{
		return nil,errors.New("不是连对")
	}

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
	pokersNum := len(pokers)
	if pokersNum < 3 || pokersNum >5{
		return nil,errors.New("不是三带牌")
	}

	poker.CommonSort(pokers)
	cardNum := poker.CheckEachCardNum(pokers)
	cardNumCount := len(cardNum)
	if pokersNum == 3{
		if cardNumCount == 1{
			return newCardsType(game.POKERS_TYPE_THREE,pokers[0].CardValue,pokers[0].CardValue),nil
		}else{
			return nil,errors.New("不是三带牌")
		}
	}else{
		if cardNumCount == 2 {

			for k,v := range cardNum{
				if v == 3{
					if(pokersNum == 4){
						return newCardsType(game.POKERS_TYPE_THREE_PLUS_ONE,k,k),nil
					}else{
						return newCardsType(game.POKERS_TYPE_THREE_PLUS_TWO,k,k),nil
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
func CheckFourPlus(pokers []*poker.PokerCard) (*subCardsType,error){

	pokersNum := len(pokers)
	if pokersNum != 6 && pokersNum != 8{
		return nil,errors.New("不是四带牌")
	}

	poker.CommonSort(pokers)

	cardNum := poker.CheckEachCardNum(pokers)
	cardNumCount := len(cardNum)
	if cardNumCount == 2{
		for k,v := range cardNum{
			if v == 4{
				if pokersNum == 6{
					return newCardsType(game.POKERS_TYPE_FOUR_PLUS_TWO,k,k),nil
				}else{
					return newCardsType(game.POKERS_TYPE_FOUR_PLUS_FOUR,k,k),nil
				}
			}
		}
		return nil,errors.New("不是四带牌")
	}else{
		return nil,errors.New("不是四带牌")
	}
}
//是否多个三带一，或三代二，或不带
func CheckMultiThreePlus(pokers []*poker.PokerCard) (*subCardsType,error){
	pokerNum := len(pokers)
	if pokerNum < 6 {
		return nil,errors.New("不是三顺")
	}

	poker.CommonSort(pokers)
	cardNum := poker.CheckEachCardNum(pokers)

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
	if mainCardNum > 1 && mainCardValues[len(mainCardValues)-1] > poker.PokerAce{
		return nil,errors.New("不是三顺")
	}

	//没有附牌
	if attachCardNum == 0{
		return newCardsType(game.POKERS_TYPE_MULITY_THREE,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum == attachCardNum{//三带一
		return newCardsType(game.POKERS_TYPE_MULITY_THREE_PLUS_ONE,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum*2 == attachCardNum{//三带二
		for _,v := range attachCardNumMap{
			if v != 2{
				return nil,errors.New("不是三顺")
			}
		}
		return newCardsType(game.POKERS_TYPE_MULITY_THREE_PLUS_TWO,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else{
		return nil,errors.New("不是三顺")
	}
}

//是否多个四带一或四代二，或不带
func CheckMultiFourPlus(pokers []*poker.PokerCard) (*subCardsType,error){

	pokerNum := len(pokers)
	if pokerNum < 8 || pokerNum%2 != 0 {
		return nil,errors.New("不是四顺")
	}

	poker.CommonSort(pokers)
	cardNum := poker.CheckEachCardNum(pokers)

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
	if mainCardNum > 1 && mainCardValues[len(mainCardValues)-1] > poker.PokerAce{
		return nil,errors.New("不是四顺")
	}

	//没有附牌
	if attachCardNum == 0{//四不带
		return newCardsType(game.POKERS_TYPE_MULITY_FOUR,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum*2 == attachCardNum{//四带二
		return newCardsType(game.POKERS_TYPE_MULITY_FOUR_PLUS_TWO,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
	}else if mainCardNum*4 == attachCardNum{//四带四
		for _,v := range attachCardNumMap{
			if v != 2{
				return nil,errors.New("不是四顺")
			}
		}
		return newCardsType(game.POKERS_TYPE_MULITY_FOUR_PLUS_FOUR,mainCardValues[0],mainCardValues[len(mainCardValues)-1]),nil
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
	//2和王不能参与顺子出牌
	if pokers[len(pokers)-1].CardValue >= poker.PokerTwo{
		return nil,errors.New("不是顺子")
	}
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


