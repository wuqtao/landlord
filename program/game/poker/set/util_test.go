package set

import "testing"

func TestCheckEachCardNum(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
			&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
			&dec.Cards[8],&dec.Cards[9],
			&dec.Cards[12]}

	cardNum := CheckEachCardNum(set1)

	if cardNum[dec.Cards[0].CardValue] != 4 ||
		cardNum[dec.Cards[4].CardValue] != 3 ||
		cardNum[dec.Cards[8].CardValue] != 2 ||
		cardNum[dec.Cards[12].CardValue] != 1{
			t.Error("TestCheckEachCardNum err")
	}

}

func TestIsUnsameCardNumSame(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[9],
		&dec.Cards[12]}

	if IsUnsameCardNumSame(set1){
		t.Error("TestIsUnsameCardNumSame err")
	}

	set2 := PokerSet{&dec.Cards[0],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[9],
		&dec.Cards[12]}

	if IsUnsameCardNumSame(set2){
		t.Error("TestIsUnsameCardNumSame err")
	}

	set3 := PokerSet{&dec.Cards[0],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[9],
		&dec.Cards[12],&dec.Cards[13]}

	if !IsUnsameCardNumSame(set3){
		t.Error("TestIsUnsameCardNumSame err")
	}

	set4 := PokerSet{&dec.Cards[0],
		&dec.Cards[4],
		&dec.Cards[8],
		&dec.Cards[12],}

	if !IsUnsameCardNumSame(set4){
		t.Error("TestIsUnsameCardNumSame err")
	}


}

func TestIsCardSame(t *testing.T) {
	set1 := PokerSet{&dec.Cards[0],&dec.Cards[1],&dec.Cards[2],&dec.Cards[3],
		&dec.Cards[4],&dec.Cards[5],&dec.Cards[6],
		&dec.Cards[8],&dec.Cards[9],
		&dec.Cards[12]}


	if !IsCardSame(set1,[]int{0,1,2,3}){
		t.Error("TestIsCardSame err")
	}

	if !IsCardSame(set1,[]int{4,5,6,}){
		t.Error("TestIsCardSame err")
	}

	if IsCardSame(set1,[]int{0,1,5,6}){
		t.Error("TestIsCardSame err")
	}

	if IsCardSame(set1,[]int{0,1,5,6,7,8}){
		t.Error("TestIsCardSame err")
	}
}