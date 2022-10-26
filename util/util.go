package util

import "strings"

func ConvertCol(col byte) int {
	col = strings.ToUpper(string(col))[0]
	return int(col - 'A')
}

func ConvertRow(row byte) int {
	return 5 - int((row - '1'))
}

func ConvertColBack(col int) byte {
	return byte(col + 'A')
}

func ConvertSquare(square string) int {
	col, row := square[0], square[1]
	colIndex, rowIndex := ConvertCol(col), ConvertRow(row)
	return colIndex + rowIndex*7
}
