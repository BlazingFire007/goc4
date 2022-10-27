package cache

import "werichardson.com/connect4/src/board"

type Key board.Bitboard
type Value int8

type Table struct {
	entries map[Key]Value
}

func NewTable() *Table {
	return &Table{
		entries: make(map[Key]Value),
	}
}

func (t *Table) Get(key Key) (Value, bool) {
	v, ok := t.entries[key]
	return v, ok
}

func (t *Table) Set(key Key, value Value) {
	t.entries[key] = value
}

func (t *Table) Reset() {
	for k := range t.entries {
		delete(t.entries, k)
	}
}
