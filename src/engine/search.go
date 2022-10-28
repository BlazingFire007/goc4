package engine

import (
	"fmt"
	"time"

	"werichardson.com/connect4/src/board"
	"werichardson.com/connect4/src/cache"
)

var table = cache.NewTable()
var nodes uint64 = 0

func Root(b board.Board, seconds float64) byte {
	const maxDepth = 43
	var bestScore int = -1000
	var bestMove byte
	start := time.Now()
	for depth := 11; depth <= maxDepth; depth++ {
		if time.Since(start).Seconds() > seconds {
			break
		}
		move, score := RootSearch(b, depth, start, seconds)
		fmt.Printf("Depth: %d, Move: %s, Score: %d\n", depth, string(move), score)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
		if score > 100 {
			break
		}
	}
	fmt.Println("Nodes: ", nodes)
	return bestMove
}

func RootSearch(b board.Board, depth int, start time.Time, seconds float64) (byte, int) {
	var ply int = 0

	moves := board.GetMoves(b)

	var alpha int = -100 - depth
	var beta int = -alpha
	var bestMove byte
	var bestScore int = -100 - depth
	for _, move := range moves {
		if time.Since(start).Seconds() > seconds {
			break
		}
		b.Move(move)
		score := -negamax(b, depth-1, -beta, -alpha, ply+1)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
		if bestScore > 100 {
			return bestMove, bestScore
		}
		if bestScore > alpha {
			alpha = bestScore
		}
		if alpha >= beta {
			break
		}
		b.Undo(move)
	}
	return bestMove, bestScore
}

func negamax(b board.Board, depth, alpha, beta, ply int) int {
	nodes++
	if depth == 0 {
		return Eval(b, ply)
	}
	if Check_winner(b) != -1 {
		return Eval(b, ply)
	}

	var bestScore int = -1000
	moves := board.GetMoves(b)
	var score int
	for _, move := range moves {
		b.Move(move)
		key := cache.Key{First: b.Bitboards[0], Second: b.Bitboards[1]}
		val, exists := table.Get(key)
		if exists && val.Depth >= depth {
			score = -val.Score
		} else {
			score = -negamax(b, depth-1, -beta, -alpha, ply+1)
			table.Set(key, cache.Value{Depth: depth, Score: -score})
		}
		b.Undo(move)
		if score > bestScore {
			bestScore = score
		}
		if score > 100 {
			return score
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
