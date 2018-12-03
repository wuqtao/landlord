package games

import (
	"landlord/program/game/games/doudizhu"
	"landlord/program/game"
)

/**
*该包用于解决game和doudizhu包循环依赖问题
 */
func NewGame(gameID int,baseScore int) game.IGame {
	switch gameID {
	case game.GAME_TYPE_OF_DOUDOZHU:
		return doudizhu.GetDoudizhu(baseScore)
	case game.GAME_TYPE_OF_SHENGJI:
		return nil
	case game.GAME_TYPE_OF_BAOHUANG:
		return nil
	case game.GAME_TYPE_OF_ZHAJINHUA:
		return nil
	default:
		return nil
	}
}