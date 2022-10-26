package engine

import (
	"fmt"
	"math/rand"
	"time"

	"werichardson.com/connect4/board"
)

type threadResult struct {
	move   byte
	score  int
	thread chan int
}

func RootSearch(b board.Board, depth int) byte {
	rand.Seed(time.Now().UnixNano())
	var ply int = 0

	moves := board.GetMoves(b)
	// rand.Shuffle(len(moves), func(i, j int) { moves[i], moves[j] = moves[j], moves[i] })

	alpha, beta := -10000, 10000
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

	var bestScore int
	moves := board.GetMoves(b)
	var score int = -10000
	for _, move := range moves {
		b.Move(move)
		score = -negamax(b, depth-1, -beta, -alpha, ply+1)
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
