package engine

import (
	"fmt"
	"math"
	"time"

	"werichardson.com/connect4/src/board"
	"werichardson.com/connect4/src/cache"
	"werichardson.com/connect4/src/util"
)

var table = cache.NewTable(100000000)
var nodes uint64 = 0

func Root(b board.Board, seconds float64) board.Column {
	const maxDepth = 43
	var bestScore int = -10000
	var bestMove board.Column
	start := time.Now()
	for depth := 11; depth <= maxDepth; depth++ {
		if time.Since(start).Seconds() > seconds {
			break
		}
		move, score := RootSearch(b, depth, start, seconds)
		fmt.Printf("Depth: %d, Move: %s, Score: %d\n", depth, string(util.ConvertColBack(int(move))), score)
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}
	fmt.Println("Nodes: ", nodes)
	fmt.Printf("Time: %fs\n", math.Round(time.Since(start).Seconds()*100)/100)
	return bestMove
}

func RootSearch(b board.Board, depth int, start time.Time, seconds float64) (board.Column, int) {
	var ply int = 0

	moves := board.GetMoves(b)

	var alpha int = -1000 - depth
	var beta int = -alpha
	var bestMove board.Column
	var bestScore int = -1000 - depth
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

	player := 0
	if b.Turn {
		player = 1
	}

	pwin := board.CheckAlign(b.Bitboards[player])
	owin := board.CheckAlign(b.Bitboards[1-player])
	if pwin {
		return 1000 - ply
	}
	if owin {
		return -1000 + ply
	}
	if depth == 0 {
		return Eval(b)
	}

	var bestScore int = -1001
	moves := board.GetMoves(b)
	if len(moves) == 0 {
		return 0
	}
	var score int
	for _, move := range moves {
		b.Move(move)
		// check if move is in cache and retrive score if it is
		entry := table.Entries[b.Hash%table.Length]
		if entry.EntryType == cache.Exact && entry.Depth >= depth && entry.Hash == b.Hash {
			score = entry.Value
		} else {
			score = -negamax(b, depth-1, -beta, -alpha, ply+1)
			table.Entries[b.Hash%table.Length] = cache.Entry{Value: score, Hash: b.Hash, Depth: depth, EntryType: cache.Exact}
		}

		// 			score = -negamax(b, depth-1, -beta, -alpha, ply+1)

		b.Undo(move)
		if score > 0 {
			return score
		}
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
