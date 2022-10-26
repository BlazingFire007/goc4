package engine

import (
	"werichardson.com/connect4/board"
)

func Eval(b board.Board, ply int) int {
	player, opp := -1, -1
	if b.Turn {
		player, opp = 1, 0
	} else {
		player, opp = 0, 1
	}
	score := 0
	winner := check_winner(b)
	if winner == opp {
		score -= 100
	} else if winner == player {
		score += 100
	}
	score += check_center(b)
	score += ply
	return score
}

func check_center(b board.Board) int {
	player, opp := -1, -1
	if b.Turn {
		player, opp = 1, 0
	} else {
		player, opp = 0, 1
	}
	pboard := b.Bitboards[player]
	oboard := b.Bitboards[opp]

	pcenter := board.CheckCenter(pboard)
	ocenter := board.CheckCenter(oboard)
	var score int = 0
	if pcenter {
		score += 20
	}
	if ocenter {
		score -= 20
	}
	return score
}

func check_winner(b board.Board) int {
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
