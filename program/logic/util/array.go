package util

import (
	"errors"
	"math/rand"
	"time"
	"chessSever/program/logic/game/poker"
)

func Random(array []poker.PokerCard, length int) (error) {

	rand.Seed(time.Now().Unix())

	if len(array) <= 0 {
		return errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(array) <= length {
		return errors.New("the size of the parameter length illegal")
	}

	for i := len(array) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		array[i], array[num] = array[num], array[i]
	}

	return nil
}


