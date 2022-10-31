package cache

const NoEntry int8 = 0
const Alpha int8 = 1
const Beta int8 = 2
const Exact int8 = 3

type Entry struct {
	Value     int8
	Hash      uint64
	Depth     int8
	EntryType int8
}

type Table struct {
	Entries []Entry
	Length  uint64
}

func NewTable(length uint64) *Table {
	return &Table{Length: length, Entries: make([]Entry, length)}
}
