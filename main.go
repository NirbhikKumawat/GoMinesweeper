package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
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
type model struct {
	game     *Game
	irow     string
	icol     string
	isRowSel bool
}

func initialModel(m, r, c int) model {
	return model{
		game:     NewGame(m, r, c),
		isRowSel: true,
	}
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab", "enter":
			m.isRowSel = !m.isRowSel
		case "backspace":
			if m.isRowSel && len(m.irow) > 0 {
				m.irow = m.irow[:len(m.irow)-1]
			} else {
				m.icol = m.icol[:len(m.icol)-1]
			}
		case "r":
			r, _ := strconv.Atoi(m.irow)
			c, _ := strconv.Atoi(m.icol)
			if r >= 0 && r < m.game.Rows && c >= 0 && c < m.game.Cols {
				m.game.revealCell(r, c)
				m.game.Board[r][c].handleRevealed(m.game, r, c)
				m.game.checkCompleted()
			}
			m.irow = ""
			m.icol = ""
			m.isRowSel = true
		case "f":
			r, _ := strconv.Atoi(m.irow)
			c, _ := strconv.Atoi(m.icol)
			if r >= 0 && r < m.game.Rows && c >= 0 && c < m.game.Cols {
				m.game.flagCell(r, c)
			}
			m.irow = ""
			m.icol = ""
			m.isRowSel = true
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			if m.isRowSel {
				m.irow += msg.String()
			} else {
				m.icol += msg.String()
			}
		}
	}
	return m, nil
}
func (m model) View() string {
	g := m.game
	s := "Minesweeper (Press 'q' to quit)\n\n"
	r := g.Rows
	c := g.Cols

	s += "  "
	for i := 0; i < c; i++ {
		s += fmt.Sprintf("%d ", i)
	}
	s += "\n"
	for i := 0; i < r; i++ {
		s += fmt.Sprintf("%d ", i)
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsFlagged {
				s += fmt.Sprintf("\033[38;5;88mF \033[0m")
			} else if !g.Board[i][j].IsRevealed {
				s += fmt.Sprintf(". ")
			} else {
				if g.Board[i][j].NearbyMines == 0 {
					s += fmt.Sprintf("\033[90m0 \033[0m")
				} else if g.Board[i][j].NearbyMines == 1 {
					s += fmt.Sprintf("\033[34m1 \033[0m")
				} else if g.Board[i][j].NearbyMines == 2 {
					s += fmt.Sprintf("\033[32m2 \033[0m")
				} else if g.Board[i][j].NearbyMines == 3 {
					s += fmt.Sprintf("\033[31m3 \033[0m")
				} else if g.Board[i][j].NearbyMines == 4 {
					s += fmt.Sprintf("\033[35m4 \033[0m")
				} else if g.Board[i][j].NearbyMines == 5 {
					s += fmt.Sprintf("\033[38;5;214m5 \033[0m")
				} else if g.Board[i][j].NearbyMines == 6 {
					s += fmt.Sprintf("\033[36m6 \033[0m")
				} else if g.Board[i][j].NearbyMines == 7 {
					s += fmt.Sprintf("\033[33m7 \033[0m")
				} else if g.Board[i][j].NearbyMines == 8 {
					s += fmt.Sprintf("\033[91m8 \033[0m")
				}
			}
		}
		s += "\n"
	}
	rowMarker, colMarker := " ", " "
	if m.isRowSel {
		rowMarker = ">"
	} else {
		colMarker = "<"
	}
	s += fmt.Sprintf("\n%s Row: %s %s Col: %s", rowMarker, m.irow, colMarker, m.icol)
	s += "\n (Type numbers, Tab to switch,'r' to reveal,'f' to flag)\n"

	if g.GameOver {
		s += "\n\033[31mGAME OVER! Press 'q' to quit.\033[0m\n"
	}
	if g.GameWon {
		s += "\n\033[31mGAME Won! Press 'q' to quit.\033[0m\n"
	}
	return s
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
				os.Stdout.Write([]byte("\033[97m"))
				os.Stdout.Write([]byte("X "))
				os.Stdout.Write([]byte("\033[0m"))
			} else {
				if g.Board[i][j].NearbyMines == 0 {
					os.Stdout.Write([]byte("\033[90m"))
					os.Stdout.Write([]byte("0 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 1 {
					os.Stdout.Write([]byte("\033[34m"))
					os.Stdout.Write([]byte("1 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 2 {
					os.Stdout.Write([]byte("\033[32m"))
					os.Stdout.Write([]byte("2 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 3 {
					os.Stdout.Write([]byte("\033[31m"))
					os.Stdout.Write([]byte("3 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 4 {
					os.Stdout.Write([]byte("\033[35m"))
					os.Stdout.Write([]byte("4 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 5 {
					os.Stdout.Write([]byte("\033[38;5;214m"))
					os.Stdout.Write([]byte("5 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 6 {
					os.Stdout.Write([]byte("\033[36m"))
					os.Stdout.Write([]byte("6 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 7 {
					os.Stdout.Write([]byte("\033[33m"))
					os.Stdout.Write([]byte("7 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if g.Board[i][j].NearbyMines == 8 {
					os.Stdout.Write([]byte("\033[91m"))
					os.Stdout.Write([]byte("8 "))
					os.Stdout.Write([]byte("\033[0m"))
				} else {
					os.Stdout.Write([]byte(strconv.Itoa(g.Board[i][j].NearbyMines)))
					os.Stdout.Write([]byte(" "))
				}

			}
		}
		os.Stdout.Write([]byte("\n"))
	}
}

/*
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
					os.Stdout.Write([]byte("\033[38;5;88m"))
					os.Stdout.Write([]byte("F "))
					os.Stdout.Write([]byte("\033[0m"))
				} else if !g.Board[i][j].IsRevealed {
					os.Stdout.Write([]byte(". "))
				} else {
					if g.Board[i][j].IsMine {
						os.Stdout.Write([]byte("\033[34m"))
						os.Stdout.Write([]byte("X "))
						os.Stdout.Write([]byte("\033[0m"))
					} else {
						if g.Board[i][j].NearbyMines == 0 {
							os.Stdout.Write([]byte("\033[90m"))
							os.Stdout.Write([]byte("0 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 1 {
							os.Stdout.Write([]byte("\033[34m"))
							os.Stdout.Write([]byte("1 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 2 {
							os.Stdout.Write([]byte("\033[32m"))
							os.Stdout.Write([]byte("2 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 3 {
							os.Stdout.Write([]byte("\033[31m"))
							os.Stdout.Write([]byte("3 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 4 {
							os.Stdout.Write([]byte("\033[35m"))
							os.Stdout.Write([]byte("4 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 5 {
							os.Stdout.Write([]byte("\033[38;5;214m"))
							os.Stdout.Write([]byte("5 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 6 {
							os.Stdout.Write([]byte("\033[36m"))
							os.Stdout.Write([]byte("6 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 7 {
							os.Stdout.Write([]byte("\033[33m"))
							os.Stdout.Write([]byte("7 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else if g.Board[i][j].NearbyMines == 8 {
							os.Stdout.Write([]byte("\033[91m"))
							os.Stdout.Write([]byte("8 "))
							os.Stdout.Write([]byte("\033[0m"))
						} else {
							os.Stdout.Write([]byte(strconv.Itoa(g.Board[i][j].NearbyMines)))
							os.Stdout.Write([]byte(" "))
						}

					}
				}
			}
			os.Stdout.Write([]byte("\n"))
		}
	}
*/
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

/*
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
*/
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

/*
	func runMinesweeper(mines, rows, cols int) {
		minesweeper := NewGame(mines, rows, cols)
		minesweeper.mainLoop()
	}
*/
func playMinesweeper(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(initialModel(pmines, prows, pcols))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
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
