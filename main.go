package main

import "fmt"

type Cell struct {
	IsRevealed bool
	IsFlagged  bool
	IsMine     bool
}
type Game struct {
	GameOver bool
	GameWon  bool
	Board    [10][10]Cell
}

func (g *Game) printBoard(r, c int) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsMine {
				fmt.Print("X")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
func NewGame() *Game {
	g := &Game{
		GameOver: false,
		GameWon:  false,
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			g.Board[i][j] = Cell{
				IsRevealed: false,
				IsFlagged:  false,
				IsMine:     false,
			}
		}
	}
	return g
}
func main() {
	minesweeper := NewGame()
	minesweeper.printBoard(10, 10)
}
