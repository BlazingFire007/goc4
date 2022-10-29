package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"werichardson.com/connect4/src/board"
	"werichardson.com/connect4/src/engine"
)

type Options struct {
	first bool
	depth int
}

func main() {
	b := board.Board{Bitboards: [2]board.Bitboard{0, 0}, Turn: true}
	if len(os.Args) > 1 {
		depth, _ := strconv.Atoi(os.Args[1])
		b.Load(os.Args[2])
		board.Print(b)
		cmove := engine.Root(b, float64(depth))
		fmt.Println(string(cmove))
		os.Exit(0)
	}
	options := Options{first: true, depth: 12}
	fmt.Println("Welcome to Connect 4!")
	fmt.Println("Enter a move in the form of a letter (A-G) to place a piece in that column.")
	fmt.Println("The first player to get 4 pieces in a row wins!")
	fmt.Println()
	gofirstInput := ask("Would you like to go first? (Y/N): ")
	gofirstInput = strings.ToUpper(gofirstInput)
	if gofirstInput == "Y" {
		options.first = true
	} else if gofirstInput == "N" {
		options.first = false
	}
	fmt.Print("Enter a search time. The computer will use ABOUT this many seconds. Recommended: (5-20): ")
	fmt.Scanf("%d", &options.depth)
	gameLoop(b, options)
}

func gameLoop(b board.Board, options Options) {
	if !options.first {
		cmove := byte('d')
		b.Move(cmove)
		fmt.Printf("Computer move: %c\n", cmove)
	}
	for {
		board.Print(b)
		move := getMoveInput()
		tryMove := b.Move(move)
		if !tryMove {
			fmt.Println("Invalid move. Try again.")
			continue
		}
		checkGameOver(b, options)
		cmove := engine.Root(b, float64(options.depth))
		b.Move(cmove)
		fmt.Printf("Computer move: %c\n", cmove)
		checkGameOver(b, options)
	}
}

func getMoveInput() board.SquareCol {
	moveInput := ask("Enter a move: ")
	moveInput = strings.ToUpper(moveInput)
	return moveInput[0]
}

func ask(question string) string {
	var input string
	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}

func checkGameOver(b board.Board, options Options) {
	var winner int = engine.Check_winner(b)
	if board.CheckDraw(b) {
		winner = 2
	}
	if winner == -1 {
		return
	}
	board.Print(b)
	var player int
	if options.first {
		player = 1
	} else {
		player = 0
	}
	if winner == player {
		fmt.Println("You win!")
	} else if winner != player {
		fmt.Println("You lose!")
	} else {
		fmt.Println("Draw!")
	}
	os.Exit(0)
}
