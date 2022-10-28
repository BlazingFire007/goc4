package cache

import (
	"werichardson.com/connect4/src/board"
)

type Key struct {
	First  board.Bitboard
	Second board.Bitboard
}
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
