package engine

import (
	"werichardson.com/connect4/src/board"
)

func Eval(b board.Board, ply int) int {
	player, opp := -1, -1
	if b.Turn {
		player, opp = 1, 0
	} else {
		player, opp = 0, 1
	}
	score := 0
	pboard := b.Bitboards[player]
	oboard := b.Bitboards[opp]

	pwin := board.CheckWin(pboard)
	owin := board.CheckWin(oboard)
	winner := -1
	if pwin {
		winner = player
	} else if owin {
		winner = opp
	}
	if winner == opp {
		score -= (100 + ply)
	}
	if winner == player {
		score += (100 - ply)
	}
	return score
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

	pwin := board.CheckWin(pboard)
	owin := board.CheckWin(oboard)

	if pwin {
		return player
	} else if owin {
		return opp
	}
	return -1

}
