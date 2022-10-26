package board

import (
	"fmt"

	"werichardson.com/connect4/util"
)

type Bitboard uint64

// Bitboards
// 1: X
// 0: O

// Position
// 1: Occupied
// 0: Empty

type Board struct {
	Position  Bitboard
	Bitboards [2]Bitboard
	Turn      bool
	History   string
}

func (b *Board) Set(col, player int) {
	b.Position |= 1 << col
	b.Bitboards[player] |= 1 << col
}

func (b *Board) Unset(col int) {
	b.Position &= ^(1 << col)
	b.Bitboards[0] &= ^(1 << col)
	b.Bitboards[1] &= ^(1 << col)
}

func (b *Board) Get(col, player int) bool {
	return b.Bitboards[player]&(1<<col) != 0
}

func (b *Board) Lowest(col int) int {
	bb := b.Bitboards[0] | b.Bitboards[1]
	var pos int = 0
	for i := 0; i < 6; i++ {
		if bb&(1<<(col+i*7)) == 0 {
			pos = col + i*7
		}
	}
	if pos == 0 {
		if b.Get(col, 0) || b.Get(col, 1) {
			scol := util.ConvertColBack(col)
			pos = util.ConvertSquare(string(scol) + "6")
			return pos - 7
		}
	}
	return pos
}

func (b *Board) Undo(col byte) bool {
	colPos := util.ConvertCol(col)
	b.Unset(colPos)
	b.Turn = !b.Turn
	b.History = b.History[:len(b.History)-1]
	return true
}

func (b *Board) Move(col byte) bool {
	colPos := util.ConvertCol(col)
	var player int
	if b.Turn {
		player = 1
	} else {
		player = 0
	}
	canSet := b.Lowest(colPos)
	if canSet < 0 {
		return false
	}
	b.Set(canSet, player)
	b.Turn = !b.Turn
	b.History += string(col)
	return true
}

func (b *Board) Load(s string) {
	for char := range s {
		b.Move(s[char])
	}
}

func GetMoves(b Board) []byte {
	var moves []byte
	for i := 0; i < 7; i++ {
		if b.Lowest(i) >= 0 {
			moves = append(moves, util.ConvertColBack(i))
		}
	}
	return moves
}

func Print(b Board) {
	for i := 0; i < 42; i++ {
		if i%7 == 0 {
			if i != 0 {
				fmt.Printf("|")
			}
			fmt.Printf("\n|%d|: ", 6-(i/6+1)+1)
		}
		if b.Get(i, 1) {
			fmt.Printf("|X")
		} else if b.Get(i, 0) {
			fmt.Printf("|O")
		} else {
			fmt.Printf("| ")
		}
	}
	fmt.Printf("|\n")
	fmt.Printf("     ---------------\n")
	fmt.Printf("     |A|B|C|D|E|F|G|\n\n")
}
