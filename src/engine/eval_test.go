package engine

import (
	"testing"

	"github.com/eli-rich/goc4/src/board"
)

func TestCheckWinner(t *testing.T) {
	b := board.Board{Bitboards: [2]board.Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Load("DGDGDGD")
	if Check_winner(b) != 0 {
		t.Errorf("CheckWinner() is not 0")
	}
	b.Reset()
	b.Load("CGDGDGDG")
	if Check_winner(b) != 0 {
		t.Errorf("CheckWinner() is not 0")
	}
	b.Reset()
	b.Load("ABCABCGG")
	if Check_winner(b) != -1 {
		t.Errorf("CheckWinner() is not -1")
	}
}

func TestEval(t *testing.T) {
	b := board.Board{Bitboards: [2]board.Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Load("ABCABCGG")
	val := Eval(b)
	if val != -3 {
		t.Errorf("Eval() is not -3")
	}
}
