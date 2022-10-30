package engine

import (
	"werichardson.com/connect4/src/board"
)

func Eval(b board.Board) int {
	player := 0
	if b.Turn {
		player = 1
	}
	opp := 1 ^ player

	pboard := b.Bitboards[player]
	oboard := b.Bitboards[opp]

	playerRemain := board.WinsRemaining(pboard, oboard)
	oppRemain := board.WinsRemaining(oboard, pboard)

	return playerRemain - oppRemain
}

func Check_winner(b board.Board) int {
	player, opp := -1, -1
	if b.Turn {
		player, opp = 1, 0
	} else {
		player, opp = 0, 1
	}
	pboard := b.Bitboards[player]
	oboard := b.Bitboards[opp]

	pwin := board.CheckAlign(pboard)
	owin := board.CheckAlign(oboard)

	if pwin {
		return player
	} else if owin {
		return opp
	}
	return -1

}
