package main

import (
	"fmt"
	"math/rand"
)

type Cell struct {
	IsRevealed  bool
	IsFlagged   bool
	IsMine      bool
	NearbyMines int
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
				fmt.Print(g.Board[i][j].NearbyMines, " ")
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
func (cell *Cell) checkMine() {
	if cell.IsMine {
	} else {
		cell.NearbyMines++
	}
}
func (g *Game) countMines() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if g.Board[i][j].IsMine {
				if i > 0 && i < 9 && j > 0 && j < 9 {
					g.Board[i-1][j-1].checkMine()
					g.Board[i-1][j].checkMine()
					g.Board[i-1][j+1].checkMine()
					g.Board[i][j-1].checkMine()
					g.Board[i][j+1].checkMine()
					g.Board[i+1][j-1].checkMine()
					g.Board[i+1][j].checkMine()
					g.Board[i+1][j+1].checkMine()
				} else if i == 0 && j == 0 {
					g.Board[0][1].checkMine()
					g.Board[1][0].checkMine()
					g.Board[1][1].checkMine()
				} else if i == 0 && j == 9 {
					g.Board[i][j-1].checkMine()
					g.Board[i+1][j].checkMine()
					g.Board[i+1][j-1].checkMine()
				} else if i == 9 && j == 0 {
					g.Board[i-1][j].checkMine()
					g.Board[i][j+1].checkMine()
					g.Board[i-1][j+1].checkMine()
				} else if i == 9 && j == 9 {
					g.Board[i-1][j-1].checkMine()
					g.Board[i-1][j].checkMine()
					g.Board[i][j-1].checkMine()
				} else if i == 0 {
					g.Board[i][j-1].checkMine()
					g.Board[i][j+1].checkMine()
					g.Board[i+1][j].checkMine()
					g.Board[i+1][j-1].checkMine()
					g.Board[i+1][j+1].checkMine()
				} else if i == 9 {
					g.Board[i-1][j].checkMine()
					g.Board[i-1][j-1].checkMine()
					g.Board[i-1][j+1].checkMine()
					g.Board[i][j-1].checkMine()
					g.Board[i][j+1].checkMine()
				} else if j == 0 {
					g.Board[i+1][j].checkMine()
					g.Board[i-1][j].checkMine()
					g.Board[i+1][j+1].checkMine()
					g.Board[i][j+1].checkMine()
					g.Board[i-1][j+1].checkMine()
				} else if j == 9 {
					g.Board[i+1][j].checkMine()
					g.Board[i-1][j].checkMine()
					g.Board[i+1][j-1].checkMine()
					g.Board[i][j-1].checkMine()
					g.Board[i-1][j-1].checkMine()
				}
			}
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
				IsRevealed:  false,
				IsFlagged:   false,
				NearbyMines: 0,
			}
		}
	}
	g.placeMines(12)
	g.countMines()
	return g
}
func main() {
	minesweeper := NewGame()
	minesweeper.printBoard(10, 10)
}
