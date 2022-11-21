package cache

import "testing"

func TestNewTable(t *testing.T) {
	table := NewTable(100)
	if table.Length != 100 {
		t.Errorf("Table length is not 100")
	}
}
