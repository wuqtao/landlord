package game

import "chessSever/program/game/poker/set"

func IsDoudizhuTypeBiger(type1 int,type2 int) bool{
	if type1 == set.POKERS_SET_TYPE_JOKER_BOMB && type2 != set.POKERS_SET_TYPE_JOKER_BOMB {
		return true
	}

	if type1 == set.POKERS_SET_TYPE_COMMON_BOMB && type2 != set.POKERS_SET_TYPE_JOKER_BOMB && type2 != set.POKERS_SET_TYPE_COMMON_BOMB {
		return true
	}

	return false
}