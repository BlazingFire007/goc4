package main

import (
	"fmt"

	"werichardson.com/connect4/board"
	"werichardson.com/connect4/engine"
)

func main() {
	b := board.Board{Position: 0, Bitboards: [2]board.Bitboard{0, 0}, Turn: true}
	b.Load("DCDBD")
	board.Print(b)
	fmt.Println(b.Turn)
	engine.RootSearch(b, 9)
}
