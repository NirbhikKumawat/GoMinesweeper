package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	prows  int
	pcols  int
	pmines int
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
	c := g.Cols
	r := g.Rows
	os.Stdout.Write([]byte("  "))
	for i := 0; i < c; i++ {
		os.Stdout.Write([]byte(strconv.Itoa(i)))
		os.Stdout.Write([]byte(" "))
	}
	os.Stdout.Write([]byte("\n"))
	for i := 0; i < r; i++ {
		os.Stdout.Write([]byte(strconv.Itoa(i)))
		os.Stdout.Write([]byte(" "))
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsMine {
				os.Stdout.Write([]byte("X "))
			} else {
				os.Stdout.Write([]byte(strconv.Itoa(g.Board[i][j].NearbyMines)))
				os.Stdout.Write([]byte(" "))
			}
		}
		os.Stdout.Write([]byte("\n"))
	}
}
func (g *Game) printBoard() {
	c := g.Cols
	r := g.Rows
	os.Stdout.Write([]byte("  "))
	for i := 0; i < c; i++ {
		os.Stdout.Write([]byte(strconv.Itoa(i)))
		os.Stdout.Write([]byte(" "))
	}
	os.Stdout.Write([]byte("\n"))
	for i := 0; i < r; i++ {
		os.Stdout.Write([]byte(strconv.Itoa(i)))
		os.Stdout.Write([]byte(" "))
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsFlagged {
				os.Stdout.Write([]byte("F "))
			} else if !g.Board[i][j].IsRevealed {
				os.Stdout.Write([]byte(". "))
			} else {
				if g.Board[i][j].IsMine {
					os.Stdout.Write([]byte("X "))
				} else {
					os.Stdout.Write([]byte(strconv.Itoa(g.Board[i][j].NearbyMines)))
					os.Stdout.Write([]byte(" "))
				}
			}
		}
		os.Stdout.Write([]byte("\n"))
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
func (cell *Cell) handleRevealed(g *Game, r, c int) {
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
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1)
		g.Board[r-1][c].handleRevealed(g, r-1, c)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1)
		g.Board[r][c-1].handleRevealed(g, r, c-1)
		g.Board[r][c+1].handleRevealed(g, r, c+1)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1)
		g.Board[r+1][c].handleRevealed(g, r+1, c)
		g.Board[r+1][c+1].handleRevealed(g, r+1, c+1)
	} else if r == 0 && c == 0 {
		g.Board[0][1].handleRevealed(g, 0, 1)
		g.Board[1][0].handleRevealed(g, 1, 0)
		g.Board[1][1].handleRevealed(g, 1, 1)
	} else if r == 0 && c == j-1 {
		g.Board[r][c-1].handleRevealed(g, r, c-1)
		g.Board[r+1][c].handleRevealed(g, r+1, c)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1)
	} else if r == i-1 && c == 0 {
		g.Board[r-1][c].handleRevealed(g, r-1, c)
		g.Board[r][c+1].handleRevealed(g, r, c+1)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1)
	} else if r == i-1 && c == j-1 {
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1)
		g.Board[r-1][c].handleRevealed(g, r-1, c)
		g.Board[r][c-1].handleRevealed(g, r, c-1)
	} else if r == 0 {
		g.Board[r][c-1].handleRevealed(g, r, c-1)
		g.Board[r][c+1].handleRevealed(g, r, c+1)
		g.Board[r+1][c].handleRevealed(g, r+1, c)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1)
		g.Board[r+1][c+1].handleRevealed(g, r+1, c+1)
	} else if c == 0 {
		g.Board[r+1][c].handleRevealed(g, r+1, c)
		g.Board[r-1][c].handleRevealed(g, r-1, c)
		g.Board[r+1][c+1].handleRevealed(g, r+1, c+1)
		g.Board[r][c+1].handleRevealed(g, r, c+1)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1)
	} else if r == i-1 {
		g.Board[r-1][c].handleRevealed(g, r-1, c)
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1)
		g.Board[r][c-1].handleRevealed(g, r, c-1)
		g.Board[r-1][c+1].handleRevealed(g, r-1, c+1)
		g.Board[r][c+1].handleRevealed(g, r, c+1)
	} else if c == j-1 {
		g.Board[r+1][c].handleRevealed(g, r+1, c)
		g.Board[r-1][c].handleRevealed(g, r-1, c)
		g.Board[r+1][c-1].handleRevealed(g, r+1, c-1)
		g.Board[r][c-1].handleRevealed(g, r, c-1)
		g.Board[r-1][c-1].handleRevealed(g, r-1, c-1)
	}
}
func (g *Game) revealCell(r, c int) {
	if g.Board[r][c].IsFlagged {
		os.Stdout.Write([]byte("Flagged cell should not be revealed"))
		return
	}
	if g.Board[r][c].IsMine {
		g.GameOver = true
	}
}
func (g *Game) flagCell(r, c int) {
	if g.Board[r][c].IsRevealed {
		os.Stdout.Write([]byte("Revealed Cell cannot be flagged"))
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
		g.printBoard()
		os.Stdout.Write([]byte("\n"))
		os.Stdout.Write([]byte("Enter operation: "))
		fmt.Scan(&op)
		if op != 1 && op != 2 {
			os.Stdout.Write([]byte("Invalid Operation"))
			continue
		}
		os.Stdout.Write([]byte("Enter the row no: "))
		fmt.Scan(&r)
		if r < 0 || r > i-1 {
			os.Stdout.Write([]byte("Invalid row no"))
			continue
		}
		os.Stdout.Write([]byte("Enter the column no: "))
		fmt.Scan(&c)
		if c < 0 || c > j-1 {
			os.Stdout.Write([]byte("Invalid column no!"))
			continue
		}
		switch op {
		case 1:
			g.revealCell(r, c)
			g.Board[r][c].handleRevealed(g, r, c)
		case 2:
			g.flagCell(r, c)
		default:
			os.Stdout.Write([]byte("Invalid Operation"))
			os.Stdout.Write([]byte("\n"))
			continue
		}
		if g.checkCompleted() {
			break
		}
	}
	if g.GameWon {
		os.Stdout.Write([]byte("Game Completed"))
	} else {
		g.revealBoard()
		os.Stdout.Write([]byte("Game Over"))
	}
	os.Stdout.Write([]byte("\n"))
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
func runMinesweeper(mines, rows, cols int) {
	minesweeper := NewGame(mines, rows, cols)
	minesweeper.mainLoop()
}
func playMinesweeper(cmd *cobra.Command, args []string) {
	runMinesweeper(pmines, prows, pcols)
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   "minesweeper",
		Short: "play minesweeper",
		Long:  "minesweeper in your terminal",
		Run:   playMinesweeper,
	}
	rootCmd.Flags().IntVarP(&pmines, "mines", "m", 12, "no of mines on the board")
	rootCmd.Flags().IntVarP(&prows, "rows", "r", 10, "no of rows on the board")
	rootCmd.Flags().IntVarP(&pcols, "cols", "c", 10, "no of columns on the board")

	if err := rootCmd.Execute(); err != nil {
		os.Stdout.Write([]byte(err.Error()))
		os.Exit(1)
	}
}
