package poker

//从小到大对扑克牌排序
func CommonSort(pokers []*PokerCard) []*PokerCard{
	len := len(pokers)
	for i:=0;i<len;i++{
		for j:=0;j<len-i;j++{
			if pokers[j].CardValue > pokers[j+1].CardValue{
				pokers[j+1],pokers[j] = pokers[j],pokers[j+1]
			}
		}
	}
	return pokers
}
