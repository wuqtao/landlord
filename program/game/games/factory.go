package games

import (
	"chessSever/program/game/games/doudizhu"
	"chessSever/program/game"
)

/**
*该包用于解决game和doudizhu包循环依赖问题
 */
func NewGame(gameID int,baseScore int) game.IGame {
	switch gameID {
	case game.GAME_ID_OF_DOUDOZHU:
		return doudizhu.GetDoudizhu(baseScore)
	case game.GAME_ID_OF_SHENGJI:
		return nil
	case game.GAME_ID_OF_BAOHUANG:
		return nil
	case game.GAME_ID_OF_ZHAJINHUA:
		return nil
	default:
		return nil
	}
}