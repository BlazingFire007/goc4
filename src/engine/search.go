package engine

import (
	"fmt"

	"werichardson.com/connect4/src/board"
	"werichardson.com/connect4/src/cache"
)

type threadResult struct {
	move   byte
	score  int
	thread chan int
}

var table = cache.NewTable()

func RootSearch(b board.Board, depth int) byte {
	var ply int = 0

	moves := board.GetMoves(b)

	var alpha int = -10000
	var beta int = 10000
	var bestMove byte
	var bestScore int = -10000
	var threads []threadResult
	for _, move := range moves {
		b.Move(move)
		nb := board.Board{Position: 0, Bitboards: [2]board.Bitboard{0, 0}, Turn: true}
		nb.Load(b.History)
		tr := threadResult{move: move, score: 0, thread: make(chan int)}
		go nmcaller(nb, depth-1, alpha, beta, ply+1, tr.thread)
		threads = append(threads, tr)
		b.Undo(move)
	}
	for _, tr := range threads {
		score := <-tr.thread
		fmt.Printf("Move: %c, Score: %d\n", tr.move, score)
		if score > bestScore {
			bestScore = score
			bestMove = tr.move
		}
		if bestScore > alpha {
			alpha = bestScore
		}
		if alpha >= beta {
			return bestMove
		}
	}
	return bestMove
}

func nmcaller(b board.Board, depth, alpha, beta, ply int, thread chan int) {
	thread <- -negamax(b, depth, alpha, beta, ply)
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
	for _, move := range moves {
		b.Move(move)
		value, exists := table.Get(cache.Key(b.Position))
		if exists {
			b.Undo(move)
			return int(value)
		}
		score := -negamax(b, depth-1, -beta, -alpha, ply+1)
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
		table.Set(cache.Key(b.Position), cache.Value(score))
	}
	return bestScore
}
