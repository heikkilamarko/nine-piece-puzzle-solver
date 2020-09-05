package main

import (
	"fmt"
)

type piece struct {
	top    int
	right  int
	bottom int
	left   int
}

func (p *piece) isEmpty() bool {
	return p.top == 0
}

func (p *piece) rotate() {
	top := p.top
	p.top = p.right
	p.right = p.bottom
	p.bottom = p.left
	p.left = top
}

func getBoardCoords(i int) (r int, c int) {
	return i / 3, i % 3
}

func isValidSolution() bool {

	for r := 0; r < 3; r++ {
		for c := 0; c < 2; c++ {
			if !board[r][c].isEmpty() && !board[r][c+1].isEmpty() && board[r][c].right != -board[r][c+1].left {
				return false
			}
		}
	}

	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			if !board[r][c].isEmpty() && !board[r+1][c].isEmpty() && board[r][c].bottom != -board[r+1][c].top {
				return false
			}
		}
	}

	return true
}

func printSolution() {
	fmt.Println("Solution:")
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			fmt.Printf("\t  %2d  ", board[r][c].top)
		}
		fmt.Println()
		for c := 0; c < 3; c++ {
			fmt.Printf("\t%2d  %2d", board[r][c].left, board[r][c].right)
		}
		fmt.Println()
		for c := 0; c < 3; c++ {
			fmt.Printf("\t  %2d  ", board[r][c].bottom)
		}
		fmt.Println()
	}
	fmt.Println()
}

func nextPiece(pieces *[9]piece, pieceNumber int) {
	if pieceNumber == 9 {
		if isValidSolution() {
			printSolution()
		}
	} else {
		r, c := getBoardCoords(pieceNumber)
		boardLocation := &board[r][c]

		rotations := 4
		if pieceNumber == 0 {
			rotations = 1
		}

		for i := 0; i < 9; i++ {
			if !boardLocationStatus[i] {
				boardLocationStatus[i] = true
				for j := 0; j < rotations; j++ {
					*boardLocation = pieces[i]
					if isValidSolution() {
						nextPiece(pieces, pieceNumber+1)
					}
					pieces[i].rotate()
				}
				boardLocationStatus[i] = false
				*boardLocation = emptyPiece
			}
		}
	}
}

const (
	white      = 1
	black      = 2
	greenSmall = 3
	greenBig   = 4
)

var (
	emptyPiece          piece
	board               [3][3]piece
	boardLocationStatus [9]bool
)

func main() {

	pieces := [9]piece{
		{white, black, -white, greenSmall},
		{white, black, greenSmall, -greenBig},
		{-white, black, -greenSmall, greenBig},
		{-greenBig, white, black, -greenSmall},
		{greenSmall, -black, -white, greenBig},
		{greenBig, white, greenSmall, -greenBig},
		{-black, black, greenBig, -greenSmall},
		{-white, greenSmall, greenBig, -black},
		{white, -greenBig, greenSmall, black}}

	nextPiece(&pieces, 0)

}
