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
	Rows          int
	Cols          int
	GameOver      bool
	GameWon       bool
	Board         [][]Cell
	TotalCells    int
	RevealedCells int
	Mines         int
}

func (g *Game) revealBoard() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if g.Board[i][j].IsMine {
				fmt.Print("X ")
			} else {
				fmt.Print(g.Board[i][j].NearbyMines)
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func (g *Game) printBoard() {
	c := g.Cols
	r := g.Rows
	fmt.Print("  ")
	for i := 0; i < c; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
	for i := 0; i < r; i++ {
		fmt.Print(i)
		fmt.Print(" ")
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsFlagged {
				fmt.Print("F ")
			} else if !g.Board[i][j].IsRevealed {
				fmt.Print(". ")
			} else {
				if g.Board[i][j].IsMine {
					fmt.Print("X ")
				} else {
					fmt.Print(g.Board[i][j].NearbyMines)
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
func (g *Game) placeMines() {
	mines := g.Mines
	minesPlaced := 0
	for minesPlaced < mines {
		i := rand.Intn(g.Rows)
		j := rand.Intn(g.Cols)
		if !g.Board[i][j].IsMine {
			g.Board[i][j].IsMine = true
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
	r := g.Rows
	c := g.Cols
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsMine {
				if i > 0 && i < r-1 && j > 0 && j < c-1 {
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
				} else if i == 0 && j == c-1 {
					g.Board[i][j-1].checkMine()
					g.Board[i+1][j].checkMine()
					g.Board[i+1][j-1].checkMine()
				} else if i == r-1 && j == 0 {
					g.Board[i-1][j].checkMine()
					g.Board[i][j+1].checkMine()
					g.Board[i-1][j+1].checkMine()
				} else if i == r-1 && j == c-1 {
					g.Board[i-1][j-1].checkMine()
					g.Board[i-1][j].checkMine()
					g.Board[i][j-1].checkMine()
				} else if i == 0 {
					g.Board[i][j-1].checkMine()
					g.Board[i][j+1].checkMine()
					g.Board[i+1][j].checkMine()
					g.Board[i+1][j-1].checkMine()
					g.Board[i+1][j+1].checkMine()
				} else if i == r-1 {
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
				} else if j == c-1 {
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
func (cell *Cell) handleRevealed(g *Game, r, c, i, j int) {
	if cell.IsRevealed || cell.IsMine || cell.IsFlagged {
		fmt.Println("Already revealed")
		return
	}
	cell.IsRevealed = true
	g.RevealedCells++
	if cell.NearbyMines == 0 {
		g.handleRevealedNeighbours(r, c, i, j)
	}
}
func (g *Game) handleRevealedNeighbours(r, c, i, j int) {
	if g.Board[r][c].NearbyMines != 0 {
		return
	}
	if r > 0 && r < i-1 && c > 0 && c < j-1 {
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1, i, j)
		g.Board[r-1][c].handleRevealed(g, r-1, c, i, j)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1, i, j)
		g.Board[r][c-1].handleRevealed(g, r, c-1, i, j)
		g.Board[r][c+1].handleRevealed(g, r, c+1, i, j)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1, i, j)
		g.Board[r+1][c].handleRevealed(g, r+1, c, i, j)
		g.Board[r+1][c+1].handleRevealed(g, r+1, c+1, i, j)
	} else if r == 0 && c == 0 {
		g.Board[0][1].handleRevealed(g, 0, 1, i, j)
		g.Board[1][0].handleRevealed(g, 1, 0, i, j)
		g.Board[1][1].handleRevealed(g, 1, 1, i, j)
	} else if r == 0 && c == j-1 {
		g.Board[r][c-1].handleRevealed(g, r, c-1, i, j)
		g.Board[r+1][c].handleRevealed(g, r+1, c, i, j)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1, i, j)
	} else if r == i-1 && c == 0 {
		g.Board[r-1][c].handleRevealed(g, r-1, c, i, j)
		g.Board[r][c+1].handleRevealed(g, r, c+1, i, j)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1, i, j)
	} else if r == i-1 && c == j-1 {
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1, i, j)
		g.Board[r-1][c].handleRevealed(g, r-1, c, i, j)
		g.Board[r][c-1].handleRevealed(g, r, c-1, i, j)
	} else if r == 0 {
		g.Board[r][c-1].handleRevealed(g, r, c-1, i, j)
		g.Board[r][c+1].handleRevealed(g, r, c+1, i, j)
		g.Board[r+1][c].handleRevealed(g, r+1, c, i, j)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1, i, j)
		g.Board[r+1][c+1].handleRevealed(g, r+1, c+1, i, j)
	} else if c == 0 {
		g.Board[r+1][c].handleRevealed(g, r+1, c, i, j)
		g.Board[r-1][c].handleRevealed(g, r-1, c, i, j)
		g.Board[r+1][c+1].handleRevealed(g, r+1, c+1, i, j)
		g.Board[r][c+1].handleRevealed(g, r, c+1, i, j)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1, i, j)
	} else if r == i-1 {
		g.Board[r-1][c].handleRevealed(g, r-1, c, i, j)
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1, i, j)
		g.Board[r][c-1].handleRevealed(g, r, c-1, i, j)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1, i, j)
		g.Board[r][c+1].handleRevealed(g, r, c+1, i, j)
	} else if c == j-1 {
		g.Board[r+1][c].handleRevealed(g, r+1, c, i, j)
		g.Board[r-1][c].handleRevealed(g, r-1, c, i, j)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1, i, j)
		g.Board[r][c-1].handleRevealed(g, r, c-1, i, j)
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1, i, j)
	}
}
func (g *Game) revealCell(r, c int) {
	if g.Board[r][c].IsFlagged {
		fmt.Println("Flagged cell should not be revealed")
		return
	}
	if g.Board[r][c].IsMine {
		g.GameOver = true
	}
}
func (g *Game) flagCell(r, c int) {
	if g.Board[r][c].IsRevealed {
		fmt.Println("Revealed Cell cannot be flagged")
		return
	}

	g.Board[r][c].IsFlagged = !g.Board[r][c].IsFlagged
}
func (g *Game) checkCompleted() bool {
	if g.RevealedCells+g.Mines == g.TotalCells {
		g.GameWon = true
		return true
	}
	return false
}
func (g *Game) mainLoop() {
	i := g.Rows
	j := g.Cols
	var r int
	var c int
	var op int
	for !g.GameOver {
		fmt.Print("Enter operation: ")
		fmt.Scan(&op)
		if op != 1 && op != 2 {
			fmt.Println("Invalid operation")
			continue
		}
		fmt.Print("Enter the row no: ")
		fmt.Scan(&r)
		if r < 0 || r > i-1 {
			fmt.Println("Invalid row no!")
			continue
		}
		fmt.Print("Enter the column no: ")
		fmt.Scan(&c)
		if c < 0 || c > j-1 {
			fmt.Println("Invalid column no!")
			continue
		}
		switch op {
		case 1:
			g.revealCell(r, c)
			g.Board[r][c].handleRevealed(g, r, c, i, j)
		case 2:
			g.flagCell(r, c)
		default:
			fmt.Println("Invalid operation!")
			continue
		}
		g.printBoard()
		if g.checkCompleted() {
			break
		}
	}
	if g.GameWon {
		fmt.Println("Game completed")
	} else {
		g.revealBoard()
		fmt.Println("Game Over")
	}
}
func NewGame(mines, r, c int) *Game {
	g := &Game{
		Rows:          r,
		Cols:          c,
		GameOver:      false,
		GameWon:       false,
		RevealedCells: 0,
		TotalCells:    r * c,
		Mines:         mines,
	}
	g.Board = make([][]Cell, r)
	for i := range g.Board {
		g.Board[i] = make([]Cell, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			g.Board[i][j] = Cell{
				IsRevealed:  false,
				IsFlagged:   false,
				NearbyMines: 0,
			}
		}
	}
	g.placeMines()
	g.countMines()
	return g
}
func main() {
	minesweeper := NewGame(12, 10, 10)
	minesweeper.printBoard()
	minesweeper.mainLoop()
}
