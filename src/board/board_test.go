package board

import (
	"testing"
)

func TestBoardInit(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	if b.Bitboards[0] != 0 {
		t.Errorf("Bitboard[0] is not 0")
	}
	if b.Bitboards[1] != 0 {
		t.Errorf("Bitboard[1] is not 0")
	}
	if b.Turn != 0 {
		t.Errorf("Turn is not 0")
	}
	if b.Hash != 0 {
		t.Errorf("Hash is not 0")
	}
}

func TestZobrist(t *testing.T) {
	zobrist := InitZobrist()
	if zobrist[0][0] == zobrist[0][1] {
		t.Errorf("Zobrist[0][0] and Zobrist[0][1] are the same")
	}
}

func TestBoardSet(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Set(0, 0)
	if b.Bitboards[0] != 1 {
		t.Errorf("Bitboard[0] is not 1")
	}
	b.Set(1, 1)
	if b.Bitboards[1] != 2 {
		t.Errorf("Bitboard[1] is not 2")
	}
}

func TestBoardUnset(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Set(0, 0)
	b.Set(1, 1)
	b.Unset(0)
	if b.Bitboards[0] != 0 {
		t.Errorf("Bitboard[0] is not 0")
	}
	if b.Bitboards[1] != 2 {
		t.Errorf("Bitboard[1] is not 2")
	}
	b.Unset(1)
	if b.Bitboards[0] != 0 {
		t.Errorf("Bitboard[0] is not 0")
	}
	if b.Bitboards[1] != 0 {
		t.Errorf("Bitboard[1] is not 0")
	}
}
func TestBoardGet(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Set(0, 0)
	if !b.Get(0, 0) {
		t.Errorf("Get(0, 0) is not true")
	}
	if b.Get(0, 1) {
		t.Errorf("Get(0, 1) is not false")
	}
}

func TestBoardLowest(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Set(0, 0)
	b.Set(7, 0)
	b.Set(14, 0)
	b.Set(21, 0)
	b.Set(28, 0)
	b.Set(35, 0)
	if b.Lowest(0) != -1 {
		t.Errorf("Lowest(0) is not -1")
	}
	if b.Lowest(1) != 36 {
		t.Errorf("Lowest(1) is not 36")
	}
}

func TestBoardMove(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Move(0)
	b.Move(1)
	b.Move(2)
	b.Move(2)
	if b.Lowest(0) != 28 {
		t.Errorf("Lowest(0) is not 28")
	}
	if b.Lowest(1) != 29 {
		t.Errorf("Lowest(1) is not 29")
	}
	if b.Lowest(2) != 23 {
		t.Errorf("Lowest(2) is not 23")
	}
}

func TestBoardUndo(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Move(0)
	b.Move(0)
	b.Move(1)
	b.Move(1)
	b.Undo(1)
	if b.Lowest(1) != 29 {
		t.Errorf("Lowest(1) is not 29")
	}
	if b.Lowest(0) != 21 {
		t.Errorf("Lowest(0) is not 21")
	}
}

func TestBoordLoad(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Load("ABCDEFG")
	if b.Lowest(0) != 28 {
		t.Errorf("Lowest(0) is not 28")
	}
	if b.Lowest(1) != 29 {
		t.Errorf("Lowest(1) is not 29")
	}
	if b.Lowest(2) != 30 {
		t.Errorf("Lowest(2) is not 30")
	}
	if b.Lowest(3) != 31 {
		t.Errorf("Lowest(3) is not 31")
	}
	if b.Lowest(4) != 32 {
		t.Errorf("Lowest(4) is not 32")
	}
	if b.Lowest(5) != 33 {
		t.Errorf("Lowest(5) is not 33")
	}
	if b.Lowest(6) != 34 {
		t.Errorf("Lowest(6) is not 34")
	}
}

func TestBoardReset(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Load("A")
	b.Reset()
	if b.Lowest(0) != 35 {
		t.Errorf("Lowest(1) is not 36")
	}
}

func TestBoardMoves(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Load("A")
	moves := GetMoves(b)
	if len(moves) != 7 {
		t.Errorf("GetMoves() is not 7")
	}
	b.Reset()
	b.Load("ABCDEFGABCDEFGABCDEFGABCABCABC")
	moves = GetMoves(b)
	if len(moves) != 4 {
		t.Errorf("GetMoves() is not 4")
	}
}
