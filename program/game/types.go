package game

import "github.com/wqtapp/pokergame"

func IsDoudizhuTypeBiger(type1 int,type2 int) bool{
	if type1 == pokergame.LANDLORD_SET_TYPE_JOKER_BOMB && type2 != pokergame.LANDLORD_SET_TYPE_JOKER_BOMB {
		return true
	}

	if type1 == pokergame.LANDLORD_SET_TYPE_COMMON_BOMB && type2 != pokergame.LANDLORD_SET_TYPE_JOKER_BOMB && type2 != pokergame.LANDLORD_SET_TYPE_COMMON_BOMB {
		return true
	}

	return false
}