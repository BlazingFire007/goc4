package util

import "testing"

func TestConvertCol(t *testing.T) {
	if ConvertCol('a') != 0 {
		t.Errorf("ConvertCol('a') is not 0")
	}
	if ConvertCol('A') != 0 {
		t.Errorf("ConvertCol('A') is not 0")
	}
	if ConvertCol('B') != 1 {
		t.Errorf("ConvertCol('B') is not 1")
	}
	if ConvertCol('C') != 2 {
		t.Errorf("ConvertCol('C') is not 2")
	}
	if ConvertCol('D') != 3 {
		t.Errorf("ConvertCol('D') is not 3")
	}
	if ConvertCol('E') != 4 {
		t.Errorf("ConvertCol('E') is not 4")
	}
	if ConvertCol('F') != 5 {
		t.Errorf("ConvertCol('F') is not 5")
	}
	if ConvertCol('G') != 6 {
		t.Errorf("ConvertCol('G') is not 6")
	}
}

func TestConvertRow(t *testing.T) {
	if ConvertRow('1') != 5 {
		t.Errorf("ConvertRow('1') is not 5")
	}
	if ConvertRow('2') != 4 {
		t.Errorf("ConvertRow('2') is not 4")
	}
	if ConvertRow('3') != 3 {
		t.Errorf("ConvertRow('3') is not 3")
	}
	if ConvertRow('4') != 2 {
		t.Errorf("ConvertRow('4') is not 2")
	}
	if ConvertRow('5') != 1 {
		t.Errorf("ConvertRow('5') is not 1")
	}
	if ConvertRow('6') != 0 {
		t.Errorf("ConvertRow('6') is not 0")
	}
}

func TestConvertColBack(t *testing.T) {
	if ConvertColBack(0) != 'A' {
		t.Errorf("ConvertColBack(0) is not 'A'")
	}
	if ConvertColBack(1) != 'B' {
		t.Errorf("ConvertColBack(1) is not 'B'")
	}
	if ConvertColBack(2) != 'C' {
		t.Errorf("ConvertColBack(2) is not 'C'")
	}
	if ConvertColBack(3) != 'D' {
		t.Errorf("ConvertColBack(3) is not 'D'")
	}
	if ConvertColBack(4) != 'E' {
		t.Errorf("ConvertColBack(4) is not 'E'")
	}
	if ConvertColBack(5) != 'F' {
		t.Errorf("ConvertColBack(5) is not 'F'")
	}
	if ConvertColBack(6) != 'G' {
		t.Errorf("ConvertColBack(6) is not 'G'")
	}
}

func TestConvertSquare(t *testing.T) {
	if ConvertSquare("A1") != 35 {
		t.Errorf("ConvertSquare('A1') is not 35")
	}
	if ConvertSquare("G6") != 6 {
		t.Errorf("ConvertSquare('G6') is not 6")
	}
}
