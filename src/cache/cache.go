package cache

const NoEntry int = 0
const Alpha int = 1
const Beta int = 2
const Exact int = 3

type Entry struct {
	Value     int
	Hash      uint64
	Depth     int
	EntryType int
}

type Table struct {
	Entries []Entry
	Length  uint64
}

func NewTable(length uint64) *Table {
	return &Table{Length: length, Entries: make([]Entry, length)}
}
