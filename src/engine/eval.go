package engine

import (
	"github.com/eli-rich/goc4/src/board"
)

func Eval(b board.Board) int8 {

	pboard := b.Bitboards[b.Turn]
	oboard := b.Bitboards[b.Turn^1]

	playerRemain := board.WinsRemaining(pboard, oboard)
	oppRemain := board.WinsRemaining(oboard, pboard)

	return playerRemain - oppRemain
}

func Check_winner(b board.Board) int8 {
	oboard := b.Bitboards[b.Turn^1]
	owin := board.CheckAlign(oboard)
	if owin {
		return b.Turn ^ 1
	}
	return -1

}
