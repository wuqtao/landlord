package set

import (
	"testing"
	"fmt"
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
	for i,_ := range c{
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
	for i,_ := range c{
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
							t.Error(funcName+" typeInfo "+strconv.Itoa(i))
						}
					}else{
						t.Error(funcName+" typeInfo "+strconv.Itoa(i))
					}
				}else{
					t.Error(funcName+" typeInfo "+strconv.Itoa(i))
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
func echoSets(set PokerSet){
	for _,card := range set{
		fmt.Print(card)
	}
	fmt.Println("")
}
func TestPokerSet_AddPokers(t *testing.T) {
	set1 := PokerSet{}
	for i,_ := range dec.Cards{
		if set1.GetLength() != i{
			t.Error("addPokers长度不匹配")
		}
		set1 = set1.AddPokers(PokerSet{&dec.Cards[i]})
	}
}

func TestPokerSet_CheckFourPlus(t *testing.T) {
	checks := []Check{
				Check{[]string{"3"},false,-1},
				Check{[]string{"3","3"},false,-1},
				Check{[]string{"3","3","3"},false,-1},
				Check{[]string{"3","3","3","3"},false,-1},
				Check{[]string{"3","3","3","3","4"},false,-1},
				Check{[]string{"3","3","3","3","4","4"},true,POKERS_SET_TYPE_FOUR_PLUS_TWO},
				Check{[]string{"3","3","3","3","4","5"},true,POKERS_SET_TYPE_FOUR_PLUS_TWO},
				Check{[]string{"3","3","3","3","4","4","5"},false,-1},
				Check{[]string{"3","3","3","3","4","4","5","5"},true,POKERS_SET_TYPE_FOUR_PLUS_FOUR},
				Check{[]string{"3","3","3","3","5","5","5","5"},true,POKERS_SET_TYPE_FOUR_PLUS_FOUR},
				Check{[]string{"3","3","3","3","4","4","5","6"},false,-1},
				Check{[]string{"3","3","3","3","4","4","5","5","6"},false,-1},
	}

	checkBoolWithType(t,checks,"CheckFourPlus")
}

func TestPokerSet_CheckMultiFourPlus(t *testing.T) {
	checks := []Check{
		Check{[]string{"3"},false,-1},
		Check{[]string{"3","3"},false,-1},
		Check{[]string{"3","3","3"},false,-1},
		Check{[]string{"3","3","3","3"},false,-1},
		Check{[]string{"3","3","3","3","4"},false,-1},

		Check{[]string{"3","3","3","3","4","4","4","4"},true,POKERS_SET_TYPE_MULITY_FOUR},
		Check{[]string{"3","3","3","3","4","4","4","4","5","5","5","5"},true,POKERS_SET_TYPE_MULITY_FOUR},
		Check{[]string{"3","3","3","3","4","4","4","4","5"},false,-1},
		Check{[]string{"3","3","3","3","4","4","4","4","5","5"},false,-1},
		Check{[]string{"3","3","3","3","4","4","4","4","5","5","5"},false,-1},


		Check{[]string{"3","3","3","3","4","4","4","4","5","5","6","6"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		Check{[]string{"3","3","3","3","4","4","4","4","5","5","6","7"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		Check{[]string{"3","3","3","3","4","4","4","4","5","6","7","8"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		Check{[]string{"3","3","3","3","4","4","4","4","6","6","6","6"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO},
		Check{[]string{"3","3","3","3","4","4","4","4","5","6","7","8","9"},false,-1},

		Check{[]string{"3","3","3","3","4","4","4","4","6","6","7","7","8","8","9","9"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		Check{[]string{"3","3","3","3","4","4","4","4","6","6","7","7","7","7","9","9"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		Check{[]string{"3","3","3","3","4","4","4","4","6","6","6","6","8","8","9","9"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},
		Check{[]string{"3","3","3","3","4","4","4","4","6","6","6","6","8","8","8","8"},true,POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR},

	}

	checkBoolWithType(t,checks,"CheckMultiFourPlus")
}

func TestPokerSet_CheckMultiThreePlus(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[10]}
	typeInfo1,err1 := set1.CheckMultiThreePlus()
	if err1 != nil{
		echoSets(set1)
		t.Error(err1.Error())
		t.Error("CheckMultiThreePlus err1")
	}else{
		if typeInfo1.SetType != POKERS_SET_TYPE_MULITY_THREE{
			echoSets(set1)
			t.Error("CheckMultiThreePlus typeInfo1")
		}
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],
		&dec.Cards[16],&dec.Cards[17],&dec.Cards[18],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[10],}

	_,err2 := set2.CheckMultiThreePlus()
	if err2 == nil{
		echoSets(set2)
		t.Error("CheckMultiThreePlus err2")
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[12]}
	typeInfo3,err3 := set3.CheckMultiThreePlus()
	if err3 != nil{
		echoSets(set3)
		t.Error("CheckMultiThreePlus err3")
	}else{
		if typeInfo3.SetType != POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE{
			echoSets(set3)
			t.Error("CheckMultiThreePlus typeInfo3")
		}
	}

	set7 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[16],&dec.Cards[17]}
	typeInfo7,err7 := set7.CheckMultiThreePlus()
	if err7 != nil{
		echoSets(set7)
		t.Error("CheckMultiThreePlus err7")
	}else{
		if typeInfo7.SetType != POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO{
			echoSets(set7)
			t.Error("CheckMultiThreePlus typeInfo7")
		}
	}

	set4 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[16],&dec.Cards[20]}
	_,err4 := set4.CheckMultiThreePlus()
	if err4 == nil{
		echoSets(set4)
		t.Error("CheckMultiThreePlus err4")
	}

}

func TestPokerSet_CheckThreePlus(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	_,err1 := set1.CheckThreePlus()
	if err1 == nil{
		t.Error("三带拍检测错误err1")
	}

	set9 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2]}
	typeInfo9,err9 := set9.CheckThreePlus()
	if err9 != nil{
		echoSets(set9)
		t.Error("三带拍检测错误err9")
	}else{
		if typeInfo9.SetType != POKERS_SET_TYPE_THREE{
			echoSets(set9)
			t.Error("三带拍检测错误typeInfo9")
		}
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4]}
	typeInfo2,err2 := set2.CheckThreePlus()
	if err2 != nil{
		echoSets(set2)
		t.Error("三带拍检测错误err2")
	}else{
		if typeInfo2.SetType != POKERS_SET_TYPE_THREE_PLUS_ONE{
			echoSets(set2)
			t.Error("三带拍检测错误typeInfo2")
		}
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4],&dec.Cards[5]}
	typeInfo3,err3 := set3.CheckThreePlus()
	if err3 != nil{
		echoSets(set3)
		t.Error("三带拍检测错误err3")
	}else{
		if typeInfo3.SetType != POKERS_SET_TYPE_THREE_PLUS_TWO{
			echoSets(set3)
			t.Error("三带拍检测错误err3")
		}
	}

	set4 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5]}
	_,err := set4.CheckThreePlus()
	if err == nil{
		echoSets(set4)
		t.Error("三带拍检测错误err4")
	}


	set5 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4],&dec.Cards[9]}
	_,err5 := set5.CheckThreePlus()
	if err5 == nil{
		echoSets(set5)
		t.Error("三带拍检测错误err5")
	}

	set6 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8]}
	_,err6 := set6.CheckThreePlus()
	if err6 == nil{
		echoSets(set6)
		t.Error("三带拍检测错误err6")
	}
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

func TestPokerSet_GetSetTypeInfo(t *testing.T) {

	set1 := PokerSet{&dec.Cards[0]}
	type1,err1 := set1.GetSetTypeInfo()
	if err1 != nil{
		t.Error("GetSetTypeInfo err1")
	}else{
		if type1.SetType != POKERS_SET_TYPE_SINGLE{
			t.Error("GetSetTypeInfo type1")
		}
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	type2,err2 := set2.GetSetTypeInfo()
	if err2 != nil{
		t.Error("GetSetTypeInfo err2")
	}else{
		if type2.SetType != POKERS_SET_TYPE_PAIR{
			t.Error("GetSetTypeInfo type2")
		}
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2]}
	type3,err3 := set3.GetSetTypeInfo()
	if err3 != nil{
		t.Error("GetSetTypeInfo err3")
	}else{
		if type3.SetType != POKERS_SET_TYPE_THREE{
			t.Error("GetSetTypeInfo type3")
		}
	}

	set4 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4]}
	type4,err4 := set4.GetSetTypeInfo()
	if err4 != nil{
		t.Error("GetSetTypeInfo err4")
	}else{
		if type4.SetType != POKERS_SET_TYPE_THREE_PLUS_ONE{
			t.Error("GetSetTypeInfo type4")
		}
	}

	set5 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4],&dec.Cards[5]}
	type5,err5 := set5.GetSetTypeInfo()
	if err5 != nil{
		t.Error("GetSetTypeInfo err5")
	}else{
		if type5.SetType != POKERS_SET_TYPE_THREE_PLUS_TWO{
			t.Error("GetSetTypeInfo type5")
		}
	}

	set6 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5]}
	type6,err6 := set6.GetSetTypeInfo()
	if err6 != nil{
		t.Error("GetSetTypeInfo err6")
	}else{
		if type6.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
			t.Error("GetSetTypeInfo type6")
		}
	}

	set7 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3]}
	type7,err7 := set7.GetSetTypeInfo()
	if err7 != nil{
		t.Error("GetSetTypeInfo err7")
	}else{
		if type7.SetType != POKERS_SET_TYPE_COMMON_BOMB{
			t.Error("GetSetTypeInfo type7")
		}
	}

	set8 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[8]}
	type8,err8 := set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err8")
	}else{
		if type8.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
			t.Error("GetSetTypeInfo type8")
		}
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8],&dec.Cards[9]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err9")
	}else{
		if type8.SetType != POKERS_SET_TYPE_FOUR_PLUS_FOUR{
			t.Error("GetSetTypeInfo type9")
		}
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[4],&dec.Cards[8],&dec.Cards[12],&dec.Cards[16],&dec.Cards[20]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err10")
	}else{
		if type8.SetType != POKERS_SET_TYPE_DRAGON{
			t.Error("GetSetTypeInfo type10")
		}
	}
	//包含2在内不算一条龙
	set8 = PokerSet{&dec.Cards[32],&dec.Cards[36],&dec.Cards[40],&dec.Cards[44],&dec.Cards[48]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 == nil{
		t.Error("GetSetTypeInfo err11")
	}

	set8 = PokerSet{&dec.Cards[28],&dec.Cards[32],&dec.Cards[36],&dec.Cards[40],&dec.Cards[44]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err11")
	}else{
		if type8.SetType != POKERS_SET_TYPE_DRAGON{
			t.Error("GetSetTypeInfo type11")
		}
	}

	set8 = PokerSet{&dec.Cards[36],&dec.Cards[37],&dec.Cards[40],&dec.Cards[41],&dec.Cards[44],&dec.Cards[45]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err12")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULIT_PAIRS{
			t.Error("GetSetTypeInfo type12")
		}
	}

	set8 = PokerSet{&dec.Cards[36],&dec.Cards[37],&dec.Cards[40],&dec.Cards[41],&dec.Cards[44],&dec.Cards[45],&dec.Cards[48],&dec.Cards[49]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 == nil{
		t.Error("GetSetTypeInfo err11")
	}

	set8 = PokerSet{&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err13")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_THREE{
			t.Error("GetSetTypeInfo type13")
		}
	}

	set8 = PokerSet{&dec.Cards[44],&dec.Cards[45],&dec.Cards[46],&dec.Cards[48],&dec.Cards[49],&dec.Cards[50]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 == nil{
		t.Error("GetSetTypeInfo err14")
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err15")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE{
			t.Error("GetSetTypeInfo type15")
		}
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[4],&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err16")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE{
			t.Error("GetSetTypeInfo type16")
		}
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[4],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46],&dec.Cards[48],&dec.Cards[49],&dec.Cards[50]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 == nil{
		t.Error("GetSetTypeInfo err17")
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
			&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err18")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO{
			t.Error("GetSetTypeInfo type18")
		}
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[4],&dec.Cards[5],
					&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err19")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO{
			t.Error("GetSetTypeInfo type19")
		}
	}


	set8 = PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
					&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[43],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46],&dec.Cards[47]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err20")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO{
			t.Error("GetSetTypeInfo type20")
		}
	}

	set8 = PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],&dec.Cards[8],&dec.Cards[9],
		&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[43],&dec.Cards[44],&dec.Cards[45],&dec.Cards[46],&dec.Cards[47]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err21")
	}else{
		if type8.SetType != POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR{
			t.Error("GetSetTypeInfo type21")
		}
	}

	set8 = PokerSet{&dec.Cards[40],&dec.Cards[41],&dec.Cards[42],&dec.Cards[43]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err22")
	}else{
		if type8.SetType != POKERS_SET_TYPE_COMMON_BOMB{
			t.Error("GetSetTypeInfo type22")
		}
	}

	set8 = PokerSet{&dec.Cards[52],&dec.Cards[53]}
	type8,err8 = set8.GetSetTypeInfo()
	if err8 != nil{
		t.Error("GetSetTypeInfo err23")
	}else{
		if type8.SetType != POKERS_SET_TYPE_JOKER_BOMB{
			t.Error("GetSetTypeInfo type23")
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
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3]}
	if !set1.IsCommonBomb(){
		t.Error("111")
	}
	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[4]}
	if set2.IsCommonBomb(){
		t.Error("111")
	}
}

func TestPokerSet_IsDragon(t *testing.T) {
	set := PokerSet{&dec.Cards[24],&dec.Cards[28],&dec.Cards[32],&dec.Cards[36],&dec.Cards[40],&dec.Cards[44]}
	if !set.IsDragon(){
		t.Error("TestPokerSet_IsDragon err1")
	}

	set = PokerSet{&dec.Cards[24],&dec.Cards[28],&dec.Cards[32],&dec.Cards[36],&dec.Cards[40],&dec.Cards[44],&dec.Cards[48]}
	if set.IsDragon(){
		t.Error("TestPokerSet_IsDragon err2")
	}

	set = PokerSet{&dec.Cards[24],&dec.Cards[28],&dec.Cards[36],&dec.Cards[40],&dec.Cards[44]}
	if set.IsDragon(){
		t.Error("TestPokerSet_IsDragon err2")
	}
}

func TestPokerSet_IsJokerBomb(t *testing.T) {
	set := PokerSet{&dec.Cards[24],&dec.Cards[28]}
	if set.IsJokerBomb(){
		t.Error("TestPokerSet_IsJokerBomb err1")
	}

	set = PokerSet{&dec.Cards[52],&dec.Cards[53]}
	if !set.IsJokerBomb(){
		t.Error("TestPokerSet_IsJokerBomb err1")
	}

	set = PokerSet{&dec.Cards[52],&dec.Cards[52]}
	if set.IsJokerBomb(){
		t.Error("TestPokerSet_IsJokerBomb err1")
	}
}

func TestPokerSet_IsMultiPair(t *testing.T) {
	set := PokerSet{&dec.Cards[36],&dec.Cards[37],&dec.Cards[40],&dec.Cards[41],&dec.Cards[44],&dec.Cards[45]}
	if !set.IsMultiPair(){
		t.Error("TestPokerSet_IsDragon err1")
	}

	set = PokerSet{&dec.Cards[36],&dec.Cards[37],&dec.Cards[40],&dec.Cards[41],&dec.Cards[44],&dec.Cards[45],&dec.Cards[48],&dec.Cards[49]}
	if set.IsMultiPair(){
		t.Error("TestPokerSet_IsDragon err2")
	}

	set = PokerSet{&dec.Cards[24],&dec.Cards[28],&dec.Cards[36],&dec.Cards[40],&dec.Cards[44],&dec.Cards[45]}
	if set.IsMultiPair(){
		t.Error("TestPokerSet_IsDragon err2")
	}
}

func TestPokerSet_IsPair(t *testing.T) {
	set := PokerSet{&dec.Cards[36],&dec.Cards[37]}
	if !set.IsPair(){
		t.Error("1")
	}

	set = PokerSet{&dec.Cards[36],&dec.Cards[40]}
	if set.IsPair(){
		t.Error("2")
	}
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