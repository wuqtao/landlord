package set

import (
	"testing"
)
//检查不同的扑克副是否冲突，一副牌里面是否有扑克牌冲突
func Test_Deck(t *testing.T){
	dec1 := CreateDeck()
	dec2 := CreateDeck()

	tempCard := dec1.Cards[0]
	for i,_ := range dec1.Cards{
		if i != 0{
			if dec1.Cards[i].CardValue == tempCard.CardValue &&
				dec1.Cards[i].CardName == tempCard.CardName &&
				dec1.Cards[i].CardSuit == tempCard.CardSuit{
					t.Error("同一副牌中有重复的牌")
			}
		}
		if &dec1.Cards[i] == &dec2.Cards[i]{
			t.Error("不同扑克牌地址相同")
		}

		if dec1.Cards[i].CardValue != dec2.Cards[i].CardValue ||
			dec1.Cards[i].CardName != dec2.Cards[i].CardName ||
			dec1.Cards[i].CardSuit != dec2.Cards[i].CardSuit{
				t.Error("同一索引扑克牌值不一致")
		}
	}
}
