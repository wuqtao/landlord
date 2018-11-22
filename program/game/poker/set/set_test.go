package set

import (
	"testing"
	"fmt"
)

var dec PokerDeck

func init(){
	dec = CreateDeck()
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
		fmt.Println(set1)
		t.Error("四带拍检测错误err1")
	}

	set9 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2]}
	_,err9 := set9.CheckFourPlus()
	if err9 == nil{
		fmt.Println(set9)
		t.Error("四带拍检测错误")
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3]}
	_,err2 := set2.CheckFourPlus()
	if err2 == nil{
		fmt.Println(set2)
		t.Error("四带拍检测错误err2")
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4]}
	_,err3 := set3.CheckFourPlus()
	if err3 == nil{
		fmt.Println(set3)
		t.Error("四带拍检测错误err3")
	}

	set4 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5]}
	typeInfo4,err := set4.CheckFourPlus()
	if err != nil{
		fmt.Println(set4)
		t.Error("四带拍检测错误err4")
	}
	if typeInfo4.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
		fmt.Println(set4)
		t.Error("四带拍检测错误typeInfo4")
	}

	set5 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[9]}
	typeInfo5,err5 := set5.CheckFourPlus()
	if err5 != nil{
		fmt.Println(set5)
		t.Error("四带拍检测错误err5")
	}
	if typeInfo5.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
		fmt.Println(set5)
		t.Error("四带拍检测错误typeInfo5")
	}

	set6 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8],&dec.Cards[9]}
	typeInfo6,err6 := set6.CheckFourPlus()
	if err6 != nil{
		fmt.Println(set6)
		t.Error("四带拍检测错误err6")
	}
	if typeInfo6.SetType != POKERS_SET_TYPE_FOUR_PLUS_TWO{
		fmt.Println(set6)
		t.Error("四带拍检测错误typeInfo6")
	}

	set7 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8]}
	_,err7 := set7.CheckFourPlus()
	if err7 == nil{
		fmt.Println(set7)
		t.Error("四带拍检测错误err7")
	}

	set8 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],&dec.Cards[4],&dec.Cards[5],&dec.Cards[8],&dec.Cards[12]}
	_,err8 := set8.CheckFourPlus()
	if err8 == nil{
		fmt.Println(set8)
		t.Error("四带拍检测错误err8")
	}
}

func TestPokerSet_CheckMultiFourPlus(t *testing.T) {

}

func TestPokerSet_CheckMultiThreePlus(t *testing.T) {

}

func TestPokerSet_CheckThreePlus(t *testing.T) {

}

func TestPokerSet_DelPokers(t *testing.T) {

}

func TestPokerSet_DelPokersByIndex(t *testing.T) {

}

func TestPokerSet_GetLength(t *testing.T) {

}

func TestPokerSet_GetPokerByIndex(t *testing.T) {

}

func TestPokerSet_GetPokerIndex(t *testing.T) {

}

func TestPokerSet_GetPokerIndexs(t *testing.T) {

}

func TestPokerSet_GetPokersByIndexs(t *testing.T) {

}

func TestPokerSet_GetSetTypeInfo(t *testing.T) {

}

func TestPokerSet_HasSamePoker(t *testing.T) {

}

func TestPokerSet_IsCommonBomb(t *testing.T) {

}

func TestPokerSet_IsDragon(t *testing.T) {

}

func TestPokerSet_IsJokerBomb(t *testing.T) {

}

func TestPokerSet_IsMultiPair(t *testing.T) {

}

func TestPokerSet_IsPair(t *testing.T) {

}

func TestPokerSet_ReplacePoker(t *testing.T) {

}

func TestPokerSet_SortAsc(t *testing.T) {

}

func TestPokerSet_SortDesc(t *testing.T) {

}