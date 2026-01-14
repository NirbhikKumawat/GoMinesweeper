package ui

import (
	"fmt"
	"minesweeper/internal/engine"
	"strings"
	"time"
	"unicode"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	game      *engine.Game
	progress  progress.Model
	irow      string
	icol      string
	isRowSel  bool
	stopwatch stopwatch.Model
}

func InitialModel(m, r, c int) Model {
	return Model{
		stopwatch: stopwatch.NewWithInterval(time.Second),
		game:      engine.NewGame(m, r, c),
		isRowSel:  true,
		progress:  progress.New(progress.WithDefaultGradient()),
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.isRowSel = !m.isRowSel
		case "backspace":
			if m.isRowSel && len(m.irow) > 0 {
				m.irow = m.irow[:len(m.irow)-1]
			} else if len(m.icol) > 0 {
				m.icol = m.icol[:len(m.icol)-1]
			}
		case "enter":
			if m.game.FirstTurn {
				cmd = m.stopwatch.Start()
			}
			r := cellIndex(m.irow)
			c := cellIndex(m.icol)
			if r >= 0 && r < m.game.Rows && c >= 0 && c < m.game.Cols {
				m.game.RevealCell(r, c)
				m.game.Board[r][c].HandleRevealed(m.game, r, c)
				m.game.CheckCompleted()
			}
			m.irow = ""
			m.icol = ""
			m.isRowSel = true
			if m.game.GameWon || m.game.GameOver {
				return m, m.stopwatch.Stop()
			}
		case ".":
			r := cellIndex(m.irow)
			c := cellIndex(m.icol)
			if r >= 0 && r < m.game.Rows && c >= 0 && c < m.game.Cols {
				m.game.FlagCell(r, c)
			}
			m.irow = ""
			m.icol = ""
			m.isRowSel = true
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
			if m.isRowSel {
				m.irow += msg.String()
			} else {
				m.icol += msg.String()
			}
		case "ctrl+r":
			return InitialModel(m.game.Mines, m.game.Rows, m.game.Cols), m.stopwatch.Reset()
		}
	}
	var swcmd tea.Cmd
	m.stopwatch, swcmd = m.stopwatch.Update(msg)
	return m, tea.Batch(cmd, swcmd)
}
func (m Model) View() string {
	var b strings.Builder
	g := m.game
	b.WriteString("Minesweeper (Press 'q' to quit)\n")
	b.WriteString(m.stopwatch.View() + "\n")
	ratio := float64(g.RevealedCells) / float64(g.TotalCells-g.Mines)
	b.WriteString("\n" + m.progress.ViewAs(ratio) + "\n\n" + "  ")

	if !g.GameOver {
		b.WriteString(m.Board())
		rowMarker, colMarker := " ", " "
		if m.isRowSel {
			rowMarker = ">"
		} else {
			colMarker = "<"
		}
		b.WriteString(fmt.Sprintf("\n%s Row: %s %s Col: %s", rowMarker, m.irow, colMarker, m.icol))
		b.WriteString("\n (Type numbers, Tab to switch,'r' to reveal,'f' to flag)\n")
	} else if g.GameOver {
		b.WriteString("\n")
		b.WriteString("  ")
		b.WriteString(m.Board())
		b.WriteString("\n\033[31mGAME OVER! Press 'q' to quit.\033[0m\n")
	}
	if g.GameWon {
		b.WriteString("\n\033[31mGAME Won! Press 'q' to quit.\033[0m\n")
		b.WriteString("Solved in " + m.stopwatch.View() + "\n")
	}
	return b.String()
}
func (m Model) Board() string {
	g := m.game
	r := g.Rows
	c := g.Cols
	var b strings.Builder
	for i := 0; i < c; i++ {
		if i < 26 {
			b.WriteString(fmt.Sprintf("%c ", i+'A'))
		} else {
			b.WriteString(fmt.Sprintf("%c ", i-26+'a'))
		}
	}
	b.WriteString("\n")
	for i := 0; i < r; i++ {
		if i < 26 {
			b.WriteString(fmt.Sprintf("%c ", i+'A'))
		} else {
			b.WriteString(fmt.Sprintf("%c ", i-26+'a'))
		}
		for j := 0; j < c; j++ {
			if g.Board[i][j].IsFlagged {
				b.WriteString(fmt.Sprintf("\033[38;5;88mF \033[0m"))
			} else if !g.Board[i][j].IsRevealed && !g.GameOver {
				b.WriteString(". ")
			} else {
				if g.Board[i][j].NearbyMines == 0 {
					b.WriteString(fmt.Sprintf("\033[90m0 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 1 {
					b.WriteString(fmt.Sprintf("\033[34m1 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 2 {
					b.WriteString(fmt.Sprintf("\033[32m2 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 3 {
					b.WriteString(fmt.Sprintf("\033[31m3 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 4 {
					b.WriteString(fmt.Sprintf("\033[35m4 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 5 {
					b.WriteString(fmt.Sprintf("\033[38;5;214m5 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 6 {
					b.WriteString(fmt.Sprintf("\033[36m6 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 7 {
					b.WriteString(fmt.Sprintf("\033[33m7 \033[0m"))
				} else if g.Board[i][j].NearbyMines == 8 {
					b.WriteString(fmt.Sprintf("\033[91m8 \033[0m"))
				}
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
func cellIndex(s string) int {
	if len(s) == 0 {
		return 0
	}
	r := []rune(s)[0]

	if unicode.IsUpper(r) {
		return int(r - 'A')
	}
	return int(r-'a') + 26
}
