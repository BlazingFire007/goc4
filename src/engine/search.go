package engine

import (
	"fmt"

	"werichardson.com/connect4/src/board"
	"werichardson.com/connect4/src/cache"
)

var table = cache.NewTable()

func RootSearch(b board.Board, depth int) byte {
	var ply int = 0

	moves := board.GetMoves(b)

	var alpha int = -100 - depth
	var beta int = 100 + depth
	var bestMove byte
	var bestScore int = -100 - depth
	for _, move := range moves {
		b.Move(move)
		nb := board.Board{Position: 0, Bitboards: [2]board.Bitboard{0, 0}, Turn: true}
		nb.Load(b.History)
		score := -negamax(nb, depth-1, -beta, -alpha, ply+1)
		fmt.Printf("Move: %d, Score: %d\n", move, score)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
		if bestScore > alpha {
			alpha = bestScore
		}
		if alpha >= beta {
			return bestMove
		}
		b.Undo(move)
	}
	return bestMove
}

func negamax(b board.Board, depth, alpha, beta, ply int) int {
	if depth == 0 {
		return Eval(b, ply)
	}
	if Check_winner(b) != -1 {
		return Eval(b, ply)
	}

	var bestScore int = -10000
	moves := board.GetMoves(b)
	var score int
	for _, move := range moves {
		b.Move(move)
		key := cache.Key{First: b.Bitboards[0], Second: b.Bitboards[1]}
		val, exists := table.Get(key)
		if !exists {
			nb := board.Board{Position: 0, Bitboards: [2]board.Bitboard{0, 0}, Turn: true}
			nb.Load(b.History)
			score = -negamax(nb, depth-1, -beta, -alpha, ply+1)
			table.Set(key, cache.Value(score))
		} else {
			score = int(val)
		}
		b.Undo(move)
		if score > bestScore {
			bestScore = score
		}
		if bestScore > alpha {
			alpha = bestScore
		}
		if alpha >= beta {
			return bestScore
		}
	}
	return bestScore
}
