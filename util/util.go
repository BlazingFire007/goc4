package util

import "strings"

func ConvertCol(col_square byte) int {
	col_square = strings.ToUpper(string(col_square))[0]
	return int(col_square - 'A')
}

func ConvertRow(row_square byte) int {
	return 5 - int((row_square - '1'))
}

func ConvertColBack(col_index int) byte {
	return byte(col_index + 'A')
}

func ConvertSquare(square string) int {
	col, row := square[0], square[1]
	colIndex, rowIndex := ConvertCol(col), ConvertRow(row)
	return colIndex + rowIndex*7
}
