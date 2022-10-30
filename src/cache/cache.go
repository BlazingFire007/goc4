package cache

import (
	"math/rand"

	"werichardson.com/connect4/src/board"
)

type Key uint64
type Value struct {
	Depth int
	Score int
}

type Table struct {
	entries map[Key]Value
}

func NewTable() *Table {
	return &Table{entries: make(map[Key]Value)}
}

func (t *Table) Get(key Key) (Value, bool) {
	val, ok := t.entries[key]
	if !ok {
		return Value{0, 0}, false
	}
	return val, true
}

func (t *Table) Set(key Key, val Value) {
	t.entries[key] = val
}

func (t *Table) Reset() {
	t.entries = make(map[Key]Value)
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
func ZobristHash(b board.Board) uint64 {
	var hash uint64
	for i := 0; i < 42; i++ {
		if b.Bitboards[0]&board.Bitboard(1<<i) != 0 {
			hash ^= zobrist[i][0]
		} else if b.Bitboards[1]&board.Bitboard(1<<i) != 0 {
			hash ^= zobrist[i][1]
		}
	}
	return hash
}
