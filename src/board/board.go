package board

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/fatih/color"
	"werichardson.com/connect4/src/util"
)

type Bitboard uint64

type SquareCol = byte
type SquareRow = byte
type Square = [2]byte

type Position int
type Column int
type Row int

// Bitboards
// 1: X
// 0: O

// Position
// 1: Occupied
// 0: Empty

type Board struct {
	Bitboards [2]Bitboard
	Turn      bool
	History   string
	Hash      uint64
}

func (b *Board) Set(pos Position, player int) {
	b.Bitboards[player] |= 1 << pos
}

func (b *Board) Unset(pos Position) {
	b.Bitboards[0] &= ^(1 << pos)
	b.Bitboards[1] &= ^(1 << pos)
}

func (b *Board) Get(pos Position, player int) bool {
	return b.Bitboards[player]&(1<<pos) != 0
}

func (b *Board) Lowest(col Position) Position {
	bb := b.Bitboards[0] | b.Bitboards[1]
	var pos int = 0
	for i := 0; i < 6; i++ {
		if bb&(1<<(int(col)+i*7)) == 0 {
			pos = int(col) + i*7
		}
	}
	if pos == 0 {
		if b.Get(col, 0) || b.Get(col, 1) {
			square_col := util.ConvertColBack(int(col))
			pos = util.ConvertSquare(string(square_col) + "6")
			return Position(pos - 7)
		}
	}
	return Position(pos)
}

func (b *Board) Undo(colSquare SquareCol) bool {
	col := Position(util.ConvertCol(colSquare))
	colPos := b.Lowest(col) + 7
	b.Unset(colPos)
	b.Turn = !b.Turn
	b.History = b.History[:len(b.History)-1]
	return true
}

func (b *Board) Move(col SquareCol) bool {
	colPos := Position(util.ConvertCol(col))
	var player int
	if b.Turn {
		player = 1
	} else {
		player = 0
	}
	lowestPosOfCol := b.Lowest(colPos)
	if lowestPosOfCol < 0 {
		return false
	}
	b.Set(lowestPosOfCol, player)
	b.Turn = !b.Turn
	b.History += string(col)
	b.Hash = ZobristHash(*b)
	return true
}

func (b *Board) Load(s string) {
	for char := range s {
		b.Move(s[char])
	}
}

func (b *Board) Reset() {
	b.Bitboards[0] = 0
	b.Bitboards[1] = 0
	b.Turn = true
	b.History = ""
}

func GetMoves(b Board) []SquareCol {
	var moves []SquareCol
	for i := 0; i < 7; i++ {
		if b.Lowest(Position(i)) >= 0 {
			moves = append(moves, util.ConvertColBack(i))
		}
	}
	sort.Slice(moves, func(i, j int) bool {
		move1 := float64(moves[i])
		move2 := float64(moves[j])
		center := float64(68)
		return math.Abs(center-move1) < math.Abs(center-move2)
	})
	return moves
}

func Print(b Board) {
	cp := color.New(color.FgHiMagenta).PrintfFunc()
	co := color.New(color.FgHiYellow).PrintfFunc()
	for i := 0; i < 42; i++ {
		if i%7 == 0 {
			if i != 0 {
				fmt.Printf("|")
			}
			fmt.Printf("\n|%d|: ", 6-(i/6+1)+1)
		}
		if b.Get(Position(i), 1) {
			fmt.Printf("|")
			cp("X")

		} else if b.Get(Position(i), 0) {
			fmt.Printf("|")
			co("O")
		} else {
			fmt.Printf("| ")
		}
	}
	fmt.Printf("|\n")
	fmt.Printf("     ---------------\n")
	fmt.Printf("     |A|B|C|D|E|F|G|\n\n")
}

// init a zobrist hash table
func InitZobrist() [42][2]uint64 {
	var zobrist [42][2]uint64
	for i := 0; i < 42; i++ {
		for j := 0; j < 2; j++ {
			zobrist[i][j] = uint64(rand.Int63())
		}
	}
	return zobrist
}

var zobrist = InitZobrist()

// zobrist hash
func ZobristHash(b Board) uint64 {
	var hash uint64
	for i := 0; i < 42; i++ {
		if b.Bitboards[0]&Bitboard(1<<i) != 0 {
			hash ^= zobrist[i][0]
		} else if b.Bitboards[1]&Bitboard(1<<i) != 0 {
			hash ^= zobrist[i][1]
		}
	}
	return hash
}
