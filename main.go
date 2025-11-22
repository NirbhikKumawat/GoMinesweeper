package main

import (
	"fmt"
	"math/rand"
)

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
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
func (g *Game) placeMines(mines int) {
	minesPlaced := 0
	for minesPlaced < mines {
		r := rand.Intn(10)
		c := rand.Intn(10)
		if !g.Board[r][c].IsMine {
			g.Board[r][c].IsMine = true
			minesPlaced++
		}
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
			}
		}
	}
	g.placeMines(10)
	return g
}
func main() {
	minesweeper := NewGame()
	minesweeper.printBoard(10, 10)
}
