package set

import (
	"testing"
	"fmt"
)

var dec PokerDeck

func init(){
	dec = CreateDeck()
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
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1]}
	_,err1 := set1.CheckFourPlus()
	if err1 == nil{
		echoSets(set1)
		t.Error("四带拍检测错误err1")
	}

	set9 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2]}
	_,err9 := set9.CheckFourPlus()
	if err9 == nil{
		echoSets(set9)
		t.Error("四带拍检测错误err9")
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3]}
	_,err2 := set2.CheckFourPlus()
	if err2 == nil{
		echoSets(set2)
		t.Error("四带拍检测错误err2")
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4]}
	_,err3 := set3.CheckFourPlus()
	if err3 == nil{
		echoSets(set3)
		t.Error("四带拍检测错误err3")
	}

	set4 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5]}
	typeInfo4,err := set4.CheckFourPlus()
	if err != nil{
		echoSets(set4)
		t.Error("四带拍检测错误err4")
	}else{
		if typeInfo4.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
			echoSets(set4)
			t.Error("四带拍检测错误typeInfo4")
		}
	}


	set5 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[9]}
	typeInfo5,err5 := set5.CheckFourPlus()
	if err5 != nil{
		echoSets(set5)
		t.Error("四带拍检测错误err5")
	}else{
		if typeInfo5.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
			echoSets(set5)
			t.Error("四带拍检测错误typeInfo5")
		}
	}
	
	set6 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8],&dec.Cards[9]}
	typeInfo6,err6 := set6.CheckFourPlus()
	if err6 != nil{
		echoSets(set6)
		t.Error("四带拍检测错误err6")
	}else{
		if typeInfo6.SetType != POKERS_SET_TYPE_FOUR_PLUS_FOUR{
			echoSets(set6)
			t.Error("四带拍检测错误typeInfo6")
		}
	}

	set7 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8]}
	_,err7 := set7.CheckFourPlus()
	if err7 == nil{
		echoSets(set7)
		t.Error("四带拍检测错误err7")
	}

	set8 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8],&dec.Cards[12]}
	_,err8 := set8.CheckFourPlus()
	if err8 == nil{
		echoSets(set8)
		t.Error("四带拍检测错误err8")
	}
}

func TestPokerSet_CheckMultiFourPlus(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
					&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
					&dec.Cards[8],&dec.Cards[9],&dec.Cards[10],&dec.Cards[11]}
	typeInfo1,err1 := set1.CheckMultiFourPlus()
	if err1 != nil{
		echoSets(set1)
		t.Error(err1.Error())
		t.Error("CheckMultiFourPlus err1")
	}else{
		if typeInfo1.SetType != POKERS_SET_TYPE_MULITY_FOUR{
			echoSets(set1)
			t.Error("CheckMultiFourPlus typeInfo1")
		}
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[16],&dec.Cards[17],&dec.Cards[18],&dec.Cards[19],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[10],&dec.Cards[11]}

	_,err2 := set2.CheckMultiFourPlus()
	if err2 == nil{
		echoSets(set2)
		t.Error("CheckMultiFourPlus err2")
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
		&dec.Cards[8],&dec.Cards[12],&dec.Cards[16],&dec.Cards[20]}
	typeInfo3,err3 := set3.CheckMultiFourPlus()
	if err3 != nil{
		echoSets(set3)
		t.Error("CheckMultiFourPlus err3")
	}else{
		if typeInfo3.SetType != POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO{
			echoSets(set3)
			t.Error("CheckMultiFourPlus typeInfo3")
		}
	}

	set7 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[16],&dec.Cards[20]}
	typeInfo7,err7 := set7.CheckMultiFourPlus()
	if err7 != nil{
		echoSets(set7)
		t.Error("CheckMultiFourPlus err7")
	}else{
		if typeInfo7.SetType != POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO{
			echoSets(set7)
			t.Error("CheckMultiFourPlus typeInfo7")
		}
	}

	set4 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[16],&dec.Cards[17]}
	typeInfo4,err4 := set4.CheckMultiFourPlus()
	if err4 != nil{
		echoSets(set4)
		t.Error("CheckMultiFourPlus err4")
	}else{
		if typeInfo4.SetType != POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO{
			echoSets(set4)
			t.Error("CheckMultiFourPlus typeInfo4")
		}
	}

	set5 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
		&dec.Cards[8],&dec.Cards[9],&dec.Cards[16],&dec.Cards[17],
		&dec.Cards[20],&dec.Cards[21],&dec.Cards[24],&dec.Cards[25]}
	typeInfo5,err5 := set5.CheckMultiFourPlus()
	if err5 != nil{
		echoSets(set5)
		t.Error("CheckMultiFourPlus err5")
	}else{
		if typeInfo5.SetType != POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR{
			echoSets(set5)
			t.Error("CheckMultiFourPlus typeInfo5")
		}
	}

	set8 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],&dec.Cards[7],
		&dec.Cards[8],&dec.Cards[18],&dec.Cards[16],&dec.Cards[17],
		&dec.Cards[20],&dec.Cards[21],&dec.Cards[24],&dec.Cards[25]}
	_,err8 := set8.CheckMultiFourPlus()
	if err8 == nil{
		echoSets(set8)
		t.Error("CheckMultiFourPlus err8")
	}
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