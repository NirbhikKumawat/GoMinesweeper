package engine

import (
	"math/rand"
	"os"
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
func (cell *Cell) HandleRevealed(g *Game, r, c int) {
	if cell.IsRevealed || cell.IsMine || cell.IsFlagged {
		return
	}
	cell.IsRevealed = true
	g.RevealedCells++
	if cell.NearbyMines == 0 {
		g.handleRevealedNeighbours(r, c)
	}
}
func (g *Game) handleRevealedNeighbours(r, c int) {
	i := g.Rows
	j := g.Cols
	if g.Board[r][c].NearbyMines != 0 {
		return
	}
	if r > 0 && r < i-1 && c > 0 && c < j-1 {
		g.Board[r-1][c-1].HandleRevealed(g, r-1, c-1)
		g.Board[r-1][c].HandleRevealed(g, r-1, c)
		g.Board[r-1][c+1].HandleRevealed(g, r-1, c+1)
		g.Board[r][c-1].HandleRevealed(g, r, c-1)
		g.Board[r][c+1].HandleRevealed(g, r, c+1)
		g.Board[r+1][c-1].HandleRevealed(g, r+1, c-1)
		g.Board[r+1][c].HandleRevealed(g, r+1, c)
		g.Board[r+1][c+1].HandleRevealed(g, r+1, c+1)
	} else if r == 0 && c == 0 {
		g.Board[0][1].HandleRevealed(g, 0, 1)
		g.Board[1][0].HandleRevealed(g, 1, 0)
		g.Board[1][1].HandleRevealed(g, 1, 1)
	} else if r == 0 && c == j-1 {
		g.Board[r][c-1].HandleRevealed(g, r, c-1)
		g.Board[r+1][c].HandleRevealed(g, r+1, c)
		g.Board[r+1][c-1].HandleRevealed(g, r+1, c-1)
	} else if r == i-1 && c == 0 {
		g.Board[r-1][c].HandleRevealed(g, r-1, c)
		g.Board[r][c+1].HandleRevealed(g, r, c+1)
		g.Board[r-1][c+1].HandleRevealed(g, r-1, c+1)
	} else if r == i-1 && c == j-1 {
		g.Board[r-1][c-1].HandleRevealed(g, r-1, c-1)
		g.Board[r-1][c].HandleRevealed(g, r-1, c)
		g.Board[r][c-1].HandleRevealed(g, r, c-1)
	} else if r == 0 {
		g.Board[r][c-1].HandleRevealed(g, r, c-1)
		g.Board[r][c+1].HandleRevealed(g, r, c+1)
		g.Board[r+1][c].HandleRevealed(g, r+1, c)
		g.Board[r+1][c-1].HandleRevealed(g, r+1, c-1)
		g.Board[r+1][c+1].HandleRevealed(g, r+1, c+1)
	} else if c == 0 {
		g.Board[r+1][c].HandleRevealed(g, r+1, c)
		g.Board[r-1][c].HandleRevealed(g, r-1, c)
		g.Board[r+1][c+1].HandleRevealed(g, r+1, c+1)
		g.Board[r][c+1].HandleRevealed(g, r, c+1)
		g.Board[r-1][c+1].HandleRevealed(g, r-1, c+1)
	} else if r == i-1 {
		g.Board[r-1][c].HandleRevealed(g, r-1, c)
		g.Board[r-1][c-1].HandleRevealed(g, r-1, c-1)
		g.Board[r][c-1].HandleRevealed(g, r, c-1)
		g.Board[r-1][c+1].HandleRevealed(g, r-1, c+1)
		g.Board[r][c+1].HandleRevealed(g, r, c+1)
	} else if c == j-1 {
		g.Board[r+1][c].HandleRevealed(g, r+1, c)
		g.Board[r-1][c].HandleRevealed(g, r-1, c)
		g.Board[r+1][c-1].HandleRevealed(g, r+1, c-1)
		g.Board[r][c-1].HandleRevealed(g, r, c-1)
		g.Board[r-1][c-1].HandleRevealed(g, r-1, c-1)
	}
}
func (g *Game) RevealCell(r, c int) {
	if g.Board[r][c].IsFlagged {
		os.Stdout.Write([]byte("Flagged cell should not be revealed"))
		return
	}
	if g.Board[r][c].IsMine {
		g.GameOver = true
	}
}
func (g *Game) FlagCell(r, c int) {
	if g.Board[r][c].IsRevealed {
		os.Stdout.Write([]byte("Revealed Cell cannot be flagged"))
		return
	}

	g.Board[r][c].IsFlagged = !g.Board[r][c].IsFlagged
}
func (g *Game) CheckCompleted() bool {
	if g.RevealedCells+g.Mines == g.TotalCells {
		g.GameWon = true
		return true
	}
	return false
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
