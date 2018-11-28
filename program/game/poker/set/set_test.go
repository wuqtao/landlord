package set

import (
	"testing"
	"reflect"
	"strconv"
)

var dec PokerDeck

func init(){
	dec = CreateDeck()
}

type Check struct{
	setStr []string   //牌型字符串
	isTrue bool       //是否检测的类型
	setType int       //扑克集类型
}

func checkBool(t *testing.T,c []Check,funcName string){
	for i := range c{
		set := getPokerset(c[i].setStr)
		setInfo := reflect.ValueOf(set)
		method := setInfo.MethodByName(funcName)
		if method.String() == "<invalid Value>"{
			t.Error("no this func "+funcName)
		}else{
			res := method.Call(nil)
			if res[0].Interface().(bool) != c[i].isTrue{
				t.Error(funcName+strconv.Itoa(i))
			}
		}
	}
}

func checkBoolWithType(t *testing.T,c []Check,funcName string){
	for i := range c{
		set := getPokerset(c[i].setStr)
		setInfo := reflect.ValueOf(set)
		method := setInfo.MethodByName(funcName)
		if method.String() == "<invalid Value>"{
			t.Error("no this func "+funcName)
		}else{
			res := method.Call(nil)
			tempBool := false
			if res[1].Interface() == nil{
				tempBool = true
			}

			if tempBool != c[i].isTrue {
				t.Error(funcName+" err"+strconv.Itoa(i))
			}else{

				if !c[i].isTrue{
					continue
				}

				if res[0].Interface() != nil{
					setTypeInfo,ok := res[0].Interface().(*SetTypeInfo)
					if ok{
						if setTypeInfo == nil || setTypeInfo.SetType != c[i].setType{
							if setTypeInfo == nil{
								t.Error(funcName+" typeInfo nil"+strconv.Itoa(i))
							}else{
								t.Error(funcName+" typeInfo "+strconv.Itoa(setTypeInfo.SetType)+" "+strconv.Itoa(i))
							}
						}
					}else{
						t.Error(funcName+" typeInfo "+strconv.Itoa(setTypeInfo.SetType)+" "+strconv.Itoa(i))
					}
				}else{
					t.Error(funcName+" typeInfo result nil"+strconv.Itoa(i))
				}
			}
		}
	}
}

//根据传入的字符数组，生成相应的扑克集，便于写测试用例
//此处不区分花色
func getPokerset(setString []string) PokerSet{
	set := PokerSet{}
	for _,name := range setString{
		for i,card := range dec.Cards{
			if card.GetCardName() == name{
				set = append(set,&dec.Cards[i])
				break
			}
		}
	}
	return set
}
//先测试生成扑克集的正确性
func TestPokerSet_GetPokerset(t *testing.T){
	test := []string{"3","4","5"}
	set := getPokerset(test)
	if set[0].CardName != test[0]{
		t.Error("pokerset creator err")
	}

	if set[1].CardName != test[1]{
		t.Error("pokerset creator err")
	}

	if set[2].CardName != test[2]{
		t.Error("pokerset creator err")
	}

}
func TestPokerSet_AddPokers(t *testing.T) {
	set1 := PokerSet{}
	for i := range dec.Cards{
		if set1.GetLength() != i{
			t.Error("addPokers长度不匹配")
		}
		set1 = set1.AddPokers(PokerSet{&dec.Cards[i]})
	}
}

func TestPokerSet_CheckFourPlus(t *testing.T) {
	checks := []Check{
				{[]string{"3"},false,-1},
				{[]string{"3","3"},false,-1},
				{[]string{"3","3","3"},false,-1},
				{[]string{"3","3","3","3"},false,-1},
				{[]string{"3","3","3","3","4"},false,-1},
				{[]string{"3","3","3","3","4","4"},true,POKERS_SET_TYPE_FOUR_PLUS_TWO},
				{[]string{"3","3","3","3","4","5"},true,POKERS_SET_TYPE_FOUR_PLUS_TWO},
				{[]string{"3","3","3","3","4","4","5"},false,-1},
				{[]string{"3","3","3","3","4","4","5","5"},true,POKERS_SET_TYPE_FOUR_PLUS_FOUR},
				{[]string{"3","3","3","3","5","5","5","5"},true,POKERS_SET_TYPE_FOUR_PLUS_FOUR},
				{[]string{"3","3","3","3","4","4","4","4"},false,-1},
				{[]string{"3","3","3","3","4","4","5","6"},false,-1},
				{[]string{"3","3","3","3","4","4","5","5","6"},false,-1},
	}

	checkBoolWithType(t,checks,"CheckFourPlus")
}

func TestPokerSet_CheckMultiFourPlus(t *testing.T) {
	checks := []Check{
		{[]string{"3"},false,-1},
		{[]string{"3","3"},false,-1},
		{[]string{"3","3","3"},false,-1},
		{[]string{"3","3","3","3"},false,-1},
		{[]string{"3","3","3","3","4"},false,-1},

		{[]string{"3","3","3","3","4","4","4","4"},true,POKERS_SET_TYPE_MULITY_FOUR},
		{[]string{"3","3","3","3","4","4","4","4","5","5","5","5"},true,POKERS_SET_TYPE_MULITY_FOUR},
		{[]string{"3","3","3","3","4","4","4","4","5"},false,-1},
		{[]string{"3","3","3","3","4","4","4","4","5","5"},false,-1},
		{[]string{"3","3","3","3","4","4","4","4","5","5","5"},false,-1},


		{[]string{"3","3","3","3","4","4","4","4","5","5","6","6"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		{[]string{"3","3","3","3","4","4","4","4","5","5","6","7"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		{[]string{"3","3","3","3","4","4","4","4","5","6","7","8"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		{[]string{"3","3","3","3","4","4","4","4","6","6","6","6"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		{[]string{"3","3","3","3","4","4","4","4","5","6","7","8","9"},false,-1},

		{[]string{"3","3","3","3","4","4","4","4","6","6","7","7","8","8","9","9"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		{[]string{"3","3","3","3","4","4","4","4","6","6","7","7","7","7","9","9"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		{[]string{"3","3","3","3","4","4","4","4","6","6","6","6","8","8","9","9"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		{[]string{"3","3","3","3","4","4","4","4","6","6","6","6","8","8","8","8"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		{[]string{"A","A","A","A","2","2","2","2"},false,-1},
		{[]string{"K","K","K","K","A","A","A","A"},true,POKERS_SET_TYPE_MULITY_FOUR},
		{[]string{"A","A","A","A","2","2","2","2","3","3","4","4","5","5","6","6"},false,-1},
		{[]string{"A","A","A","A","2","2","2","2","3","4","5","6"},false,-1},
	}

	checkBoolWithType(t,checks,"CheckMultiFourPlus")
}

func TestPokerSet_CheckMultiThreePlus(t *testing.T) {
	checks := []Check{
		{[]string{"3"},false,-1},
		{[]string{"3","3"},false,-1},
		{[]string{"3","3","3"},false,-1},
		{[]string{"3","3","3","3"},false,-1},
		{[]string{"3","3","3","3","4"},false,-1},
		{[]string{"3","3","3","3","4","4"},false,-1},
		{[]string{"3","3","3","3","4","5"},false,-1},
		{[]string{"3","3","3","4","4","4"},true,POKERS_SET_TYPE_MULITY_THREE},
		{[]string{"3","3","3","4","4","4","5","5"},true,POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE},
		{[]string{"3","3","3","4","4","4","5","6"},true,POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE},
		{[]string{"3","3","3","4","4","4","4","6"},true,POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE},
		{[]string{"3","3","3","5","5","5","6","6","6","6"},true,POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO},
		{[]string{"3","3","3","5","5","5","6","6","7","7"},true,POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO},
		{[]string{"3","3","3","5","5","5","6","6","7","8"},false,-1},
		{[]string{"3","3","3","5","5","5","6","9","7","8"},false,-1},
		{[]string{"3","3","3","4","4","4","6","6","6"},false,-1},
		{[]string{"A","A","A","2","2","2"},false,-1},
		{[]string{"A","A","A","2","2","2","3","4"},false,-1},
		{[]string{"A","A","A","2","2","2","3","3","4","4"},false,-1},
		{[]string{"k","k","k","A","A","A"},true,POKERS_SET_TYPE_MULITY_THREE},
	}

	checkBoolWithType(t,checks,"CheckMultiThreePlus")
}

func TestPokerSet_CheckThreePlus(t *testing.T) {
	checks := []Check{
		{[]string{"3"},false,-1},
		{[]string{"3","3"},false,-1},
		{[]string{"3","3","3"},true,POKERS_SET_TYPE_THREE},
		{[]string{"3","3","3","4"},true,POKERS_SET_TYPE_THREE_PLUS_ONE},
		{[]string{"3","3","3","4","4"},true,POKERS_SET_TYPE_THREE_PLUS_TWO},
		{[]string{"3","3","4","4","5"},false,-1},
	}

	checkBoolWithType(t,checks,"CheckMultiThreePlus")
}

func TestPokerSet_DelPokers(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	set,err := set.DelPokers(PokerSet{&dec.Cards[0]})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if set.GetLength() != 1  || set[0].CardValue != dec.Cards[1].CardValue{
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_DelPokersByIndex(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	set,err := set.DelPokersByIndex([]int{0})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if set.GetLength() != 1  || set[0].CardValue != dec.Cards[1].CardValue{
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetLength(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	if len(set) != set.GetLength(){
		t.Error("TestPokerSet_GetLength err")
	}
}

func TestPokerSet_GetPokerByIndex(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	card,err := set.GetPokerByIndex(1)
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if card != &dec.Cards[1]{
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetPokerIndex(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	index,err := set.GetPokerIndex(&dec.Cards[1])
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if index != 1 {
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetPokerIndexs(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	indexs,err := set.GetPokerIndexs(PokerSet{&dec.Cards[0],&dec.Cards[1]})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if indexs[0] != 0 || indexs[1] != 1 {
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}

func TestPokerSet_GetPokersByIndexs(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	cards,err := set.GetPokersByIndexs([]int{0,1})
	if err != nil{
		t.Error("TestPokerSet_DelPokers err")
	}else{
		if cards[0] != &dec.Cards[0] || cards[1] != &dec.Cards[1] {
			t.Error("TestPokerSet_DelPokers err")
		}
	}
}



func TestPokerSet_HasSamePoker(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	if set1.HasSamePoker(PokerSet{&dec.Cards[3],&dec.Cards[2]}){
		t.Error("TestPokerSet_HasSamePoker err1")
	}

	if !set1.HasSamePoker(PokerSet{&dec.Cards[0],&dec.Cards[1]}){
		t.Error("TestPokerSet_HasSamePoker err1")
	}
}

func TestPokerSet_IsCommonBomb(t *testing.T) {
	checks := []Check{
		{[]string{"3","3","3","3"},true,-1},
		{[]string{"2","2","2","2"},true,-1},
		{[]string{"3","3","3","4"},false,-1},
	}

	checkBool(t,checks,"IsCommonBomb")
}

func TestPokerSet_IsDragon(t *testing.T) {
	checks := []Check{
		{[]string{"3","4","5","6"},false,-1},
		{[]string{"3","4","5","6","7"},true,-1},
		{[]string{"3","4","5","6","7","8"},true,-1},
		{[]string{"3","4","5","6","7","8","9"},true,-1},
		{[]string{"10","J","Q","K","A"},true,-1},
		{[]string{"10","J","Q","K","A","2"},false,-1},
		{[]string{"10","J","Q","K","A","2","BlackJoker"},false,-1},
	}

	checkBool(t,checks,"IsDragon")
}

func TestPokerSet_IsJokerBomb(t *testing.T) {
	checks := []Check{
		{[]string{"3","4","5","6"},false,-1},
		{[]string{"BlackJoker","RedJoker"},true,-1},
		{[]string{"BlackJoker","BlackJoker"},false,-1},
		{[]string{"RedJoker","RedJoker"},false,-1},
	}

	checkBool(t,checks,"IsJokerBomb")
}

func TestPokerSet_IsMultiPair(t *testing.T) {
	checks := []Check{
		{[]string{"3","4","5","6"},false,-1},
		{[]string{"3","3","4","4"},false,-1},
		{[]string{"3","3","4","4","5","5"},true,-1},
		{[]string{"3","3","4","4","5","5","6","6"},true,-1},
		{[]string{"3","3","4","4","5","5","7","7"},false,-1},
	}

	checkBool(t,checks,"IsMultiPair")
}

func TestPokerSet_IsPair(t *testing.T) {
	checks := []Check{
		{[]string{"3","4"},false,-1},
		{[]string{"3","3"},true,-1},
		{[]string{"BlackJoker","RedJoker"},false,-1},
	}

	checkBool(t,checks,"IsPair")
}

func TestPokerSet_ReplacePoker(t *testing.T) {
	set := PokerSet{&dec.Cards[36],&dec.Cards[37]}
	err := set.ReplacePoker(0,&dec.Cards[0])
	if err != nil{
		t.Error("1")
	}else{
		if set[0] != &dec.Cards[0]{
			t.Error("2")
		}
	}
}

func TestPokerSet_GetSetTypeInfo(t *testing.T) {
	checks := []Check{
		{[]string{"3"}, true, POKERS_SET_TYPE_SINGLE},
		{[]string{"3", "3"}, true, POKERS_SET_TYPE_PAIR},
		{[]string{"3", "3", "4", "4", "5", "5"}, true, POKERS_SET_TYPE_MULIT_PAIRS},
		{[]string{"3", "3", "3"}, true, POKERS_SET_TYPE_THREE},
		{[]string{"3", "3", "3", "4"}, true, POKERS_SET_TYPE_THREE_PLUS_ONE},
		{[]string{"3", "3", "3", "4", "4"}, true, POKERS_SET_TYPE_THREE_PLUS_TWO},
		{[]string{"3", "3", "3", "4", "4", "4"}, true, POKERS_SET_TYPE_MULITY_THREE},
		{[]string{"3", "3", "3", "4", "4", "4", "5", "6"}, true, POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE},
		{[]string{"3", "3", "3", "4", "4", "4", "5", "6", "5", "6"}, true, POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO},

		{[]string{"3", "3", "3", "3"}, true, POKERS_SET_TYPE_COMMON_BOMB},
		{[]string{"3", "3", "3", "3", "4", "5"}, true, POKERS_SET_TYPE_FOUR_PLUS_TWO},
		{[]string{"3", "3", "3", "3", "4", "4", "5", "5"}, true, POKERS_SET_TYPE_FOUR_PLUS_FOUR},
		{[]string{"3", "3", "3", "3", "4", "4", "4", "4"}, true, POKERS_SET_TYPE_MULITY_FOUR},
		{[]string{"3", "3", "3", "3", "4", "4", "4", "4", "5", "6", "7", "8"}, true, POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		{[]string{"3", "3", "3", "3", "4", "4", "4", "4", "5", "5", "6", "6", "7", "7", "8", "8"}, true, POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},

		{[]string{"3", "4", "5", "6", "7", "8"}, true, POKERS_SET_TYPE_DRAGON},
		{[]string{"BlackJoker", "RedJoker"}, true, POKERS_SET_TYPE_JOKER_BOMB},
	}

	checkBoolWithType(t, checks, "GetSetTypeInfo")
}