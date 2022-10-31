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

const (
	X = 1
	O = 0
)

// Bitboards
// 1: X
// 0: O

// Position
// 1: Occupied
// 0: Empty

type Board struct {
	Bitboards [2]Bitboard
	Turn      int
	Hash      uint64
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

func (b *Board) Lowest(col Column) Position {
	bb := b.Bitboards[0] | b.Bitboards[1]
	start := 42 - (7 - col)
	for i := start; i >= 0; i -= 7 {
		if bb&(1<<i) == 0 {
			return Position(i)
		}
	}
	return -1
}

func (b *Board) Undo(col Column) bool {
	colPos := b.Lowest(col) + 7
	b.Unset(colPos)
	b.Turn ^= 1
	b.Hash ^= zobrist[int(colPos)][b.Turn]
	return true
}

func (b *Board) Move(col Column) bool {
	lowestPosOfCol := b.Lowest(col)
	b.Set(lowestPosOfCol, b.Turn)
	b.Turn ^= 1
	b.Hash ^= zobrist[int(lowestPosOfCol)][b.Turn]
	return true
}

func (b *Board) Load(s string) {
	for char := range s {
		b.Move(Column(util.ConvertCol(s[char])))
	}
}

func (b *Board) Reset() {
	b.Bitboards[0] = 0
	b.Bitboards[1] = 0
	b.Turn = 1
}

func GetMoves(b Board) []Column {
	var moves = make([]Column, 0, 7)
	for i := 0; i < 7; i++ {
		if b.Lowest(Column(i)) >= 0 {
			moves = append(moves, Column(i))
		}
	}
	sort.Slice(moves, func(i, j int) bool {
		move1 := float64(moves[i])
		move2 := float64(moves[j])
		center := float64(3)
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
