package set

import (
	"testing"
)

func TestBubbleSortIntMin2Max(t *testing.T) {
	ints := []int{1,3,5,4,6,2}
	BubbleSortIntMin2Max(ints)
	if ints[0] != 1 ||
		ints[1] != 2 ||
		ints[2] != 3 ||
		ints[3] != 4 ||
		ints[4] != 5 ||
		ints[5] != 6{


			t.Error("TestBubbleSortIntMin2Max err")
	}

}

func TestBubbleSortCardsMax2Min(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[10],&dec.Cards[5]}
	SortDesc(set)

	if set[0].CardValue != dec.Cards[10].CardValue ||
		set[1].CardValue != dec.Cards[5].CardValue ||
		set[2].CardValue != dec.Cards[0].CardValue{
		t.Error("TestBubbleSortCardsMax2Min err")
	}
}

func TestBubbleSortCardsMin2Max(t *testing.T) {
	set := PokerSet{&dec.Cards[0],&dec.Cards[10],&dec.Cards[5]}
	SortAsc(set)

	if set[0].CardValue != dec.Cards[0].CardValue ||
		set[1].CardValue != dec.Cards[5].CardValue ||
		set[2].CardValue != dec.Cards[10].CardValue{
		t.Error("TestBubbleSortCardsMin2Max err")
	}
}
