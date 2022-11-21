package board

import (
	"testing"
)

func TestCheckDraw(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	if CheckDraw(b) {
		t.Errorf("CheckDraw() is true")
	}
	b.Load("ABCABCGGABCABCGGABCABCGGDEFDEFDEFDEFDEFDEF")
	if !CheckDraw(b) {
		t.Errorf("CheckDraw() is false")
	}
	b.Reset()
	b.Load("DGDGDGD")
	if CheckDraw(b) {
		t.Errorf("CheckDraw is not false due to win condition")
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
}

func TestMasks(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	for _, mask := range Win_masks {
		b.Bitboards[0] = Bitboard(mask)
		if !CheckAlign(b.Bitboards[0]) {
			t.Errorf("Mask does not represent win condition")
		}
	}
}

func TestRemaining(t *testing.T) {
	b := Board{Bitboards: [2]Bitboard{0, 0}, Turn: 0, Hash: 0}
	b.Load("ABCABCABCABCED")
	if WinsRemaining(b.Bitboards[0], b.Bitboards[1]) != 38 {
		t.Errorf("Wins_remaining() is not 0")
	}
}
