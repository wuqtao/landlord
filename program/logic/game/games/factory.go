package games

import (
	"chessSever/program/logic/game/games/doudizhu"
	"chessSever/program/logic/game"
)

/**
*该包用于解决game和doudizhu包循环依赖问题
 */
func GetGame(name string) game.IGame {
	switch name {
	case game.DouDiZhu:
		return doudizhu.GetDoudizhu()
	case game.ShengJi:
		return nil
	case game.BaoHuang:
		return nil
	case game.Zhajinhua:
		return nil
	default:
		return nil
	}
}