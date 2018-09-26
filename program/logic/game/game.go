package game

import (
	"chessSever/program/logic/game/games/doudizhu"
	"chessSever/program/logic/game/games"
)

/**
*该包用于解决game和doudizhu包循环依赖问题
 */
func GetGame(name string) games.IGame {
	switch name {
	case games.DouDiZhu:
		return doudizhu.GetDoudizhu()
	case games.ShengJi:
		return nil
	case games.BaoHuang:
		return nil
	case games.Zhajinhua:
		return nil
	default:
		return nil
	}
}