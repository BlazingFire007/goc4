package engine

import (
	"fmt"
	"math"
	"time"

	"werichardson.com/connect4/src/board"
	"werichardson.com/connect4/src/cache"
	"werichardson.com/connect4/src/util"
)

var table = cache.NewTable(20000000)
var nodes uint64 = 0

const WIN_SCORE int8 = 100

func Root(b board.Board, seconds float64) board.Column {
	const maxDepth int8 = 43
	var bestMove board.Column
	start := time.Now()
	var depth int8
	for depth = 11; depth <= maxDepth; depth++ {
		if time.Since(start).Seconds() > seconds {
			break
		}
		move, score, fullDepth := RootSearch(b, depth, start, seconds)
		fmt.Printf("Depth: %d, Move: %s, Score: %d\n", depth, string(util.ConvertColBack(int(move))), score)
		if fullDepth {
			bestMove = move
		}
		if score >= WIN_SCORE-42 {
			break
		}
	}
	fmt.Println("Nodes: ", nodes)
	fmt.Printf("Time: %fs\n", math.Round(time.Since(start).Seconds()*100)/100)
	return bestMove
}

func RootSearch(b board.Board, depth int8, start time.Time, seconds float64) (board.Column, int8, bool) {
	var ply int8 = 0

	moves := board.GetMoves(b)

	var alpha int8 = -(WIN_SCORE - depth)
	var beta int8 = -alpha
	var bestMove board.Column
	var bestScore int8 = alpha
	for _, move := range moves {
		if time.Since(start).Seconds() > seconds {
			return bestMove, bestScore, false
		}
		b.Move(move)
		score := -negamax(b, depth-1, -beta, -alpha, ply+1)
		b.Undo(move)
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
	}
	return bestMove, bestScore, true
}

func negamax(b board.Board, depth int8, alpha, beta, ply int8) int8 {
	nodes++

	pwin := board.CheckAlign(b.Bitboards[b.Turn])

	if pwin {
		return WIN_SCORE - ply
	}

	if depth == 0 {
		return Eval(b)
	}

	var bestScore int8 = -(WIN_SCORE - depth)
	moves := board.GetMoves(b)
	if len(moves) == 0 {
		return 0
	}
	var score int8
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
