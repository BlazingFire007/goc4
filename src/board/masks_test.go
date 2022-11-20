package board

import "testing"

func TestCheckDraw(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	if CheckDraw(b) {
		t.Errorf("CheckDraw() is true")
	}
	b.Load("ABCABCGGABCABCGGABCABCGGDEFDEFDEFDEFDEFDEF")
	if !CheckDraw(b) {
		t.Errorf("CheckDraw() is false")
	}
}

func TestCheckAlign(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	if CheckAlign(b.Bitboards[0]) {
		t.Errorf("CheckAlign() is true")
	}
	b.Load("ABCABCABCABCED")
	if !CheckAlign(b.Bitboards[1]) {
		t.Errorf("CheckAlign() is false")
	}
	Print(b)
}
