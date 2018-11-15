package set

const(
	POKERS_SET_TYPE_SINGLE  = iota //"单牌"
	POKERS_SET_TYPE_DRAGON         //"单顺子(5个及以上)"

	POKERS_SET_TYPE_PAIR         //"对牌"
	POKERS_SET_TYPE_MULIT_PAIRS  //"双顺子(3个及以上)"

	POKERS_SET_TYPE_THREE                  //三张不带牌
	POKERS_SET_TYPE_THREE_PLUS_ONE         //三带一
	POKERS_SET_TYPE_THREE_PLUS_TWO         //三带二
	POKERS_SET_TYPE_MULITY_THREE           //3顺子不带牌
	POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE  //3顺子带1个牌
	POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO  //3顺子带2个牌

	POKERS_SET_TYPE_FOUR_PLUS_TWO          //四带2张
	POKERS_SET_TYPE_FOUR_PLUS_FOUR         //四带2对
	POKERS_SET_TYPE_MULITY_FOUR            //四顺子不带牌
	POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO   //四顺子带2张
	POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR  //四顺子带4张

	POKERS_SET_TYPE_COMMON_BOMB  // "炸弹"
	POKERS_SET_TYPE_JOKER_BOMB   // "王炸"
)

func GetPokerSetTypeName(pattern int) string {
	switch pattern {
	case POKERS_SET_TYPE_SINGLE:
		return "单牌"
	case POKERS_SET_TYPE_DRAGON:
		return "单顺子"
	case POKERS_SET_TYPE_PAIR:
		return "对牌"
	case POKERS_SET_TYPE_MULIT_PAIRS:
		return "双顺子"
	case POKERS_SET_TYPE_THREE:
		return "三张牌"
	case POKERS_SET_TYPE_THREE_PLUS_ONE:
		return "三带一"
	case POKERS_SET_TYPE_THREE_PLUS_TWO:
		return "三带二"
	case POKERS_SET_TYPE_MULITY_THREE:
		return "三顺子不带牌"
	case POKERS_SET_TYPE_MULITY_THREE_PLUS_ONE:
		return "三顺子带一张"
	case POKERS_SET_TYPE_MULITY_THREE_PLUS_TWO:
		return "三顺子带两张"
	case POKERS_SET_TYPE_FOUR_PLUS_TWO:
		return "四带二"
	case POKERS_SET_TYPE_FOUR_PLUS_FOUR:
		return "四带四"
	case POKERS_SET_TYPE_MULITY_FOUR:
		return "四顺子不带牌"
	case POKERS_SET_TYPE_MULITY_FOUR_PLUS_TWO:
		return "四顺子带两张"
	case POKERS_SET_TYPE_MULITY_FOUR_PLUS_FOUR:
		return "四顺子带四张"
	case POKERS_SET_TYPE_COMMON_BOMB:
		return "炸弹"
	case POKERS_SET_TYPE_JOKER_BOMB:
		return "火箭"
	default:
		return "未定义的牌型"
	}
}

type SetTypeInfo struct{
	SetType int
	CardValueMinAndMax map[string]int
}

func NewSetInfo(setType int,minValue int,maxValue int) *SetTypeInfo {
	setInfo := SetTypeInfo{
		SetType:setType,
		CardValueMinAndMax:make(map[string]int),
	}
	setInfo.CardValueMinAndMax["min"] = minValue
	setInfo.CardValueMinAndMax["max"] = maxValue
	return &setInfo
}

func (setInfo *SetTypeInfo)GetRangeWidth() int{
	return setInfo.CardValueMinAndMax["max"] - setInfo.CardValueMinAndMax["min"]+1
}
func (setInfo *SetTypeInfo)GetSetTypeName() string{
	return GetPokerSetTypeName(setInfo.SetType)
}
func (setInfo *SetTypeInfo)GetMinValue() int{
	return setInfo.CardValueMinAndMax["min"]
}
func (setInfo *SetTypeInfo)GetMaxValue() int{
	return setInfo.CardValueMinAndMax["max"]
}
