package ui

import (
	"fmt"
	"minesweeper/internal/engine"
	"strconv"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	game     *engine.Game
	progress progress.Model
	irow     string
	icol     string
	isRowSel bool
}

func InitialModel(m, r, c int) Model {
	return Model{
		game:     engine.NewGame(m, r, c),
		isRowSel: true,
		progress: progress.New(progress.WithDefaultGradient()),
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.game.RevealCell(r, c)
				m.game.Board[r][c].HandleRevealed(m.game, r, c)
				m.game.CheckCompleted()
			}
			m.irow = ""
			m.icol = ""
			m.isRowSel = true
		case "f":
			r, _ := strconv.Atoi(m.irow)
			c, _ := strconv.Atoi(m.icol)
			if r >= 0 && r < m.game.Rows && c >= 0 && c < m.game.Cols {
				m.game.FlagCell(r, c)
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
func (m Model) View() string {
	g := m.game
	s := "Minesweeper (Press 'q' to quit)\n"
	r := g.Rows
	c := g.Cols
	ratio := float64(g.RevealedCells) / float64(g.TotalCells-g.Mines)
	s += "\n" + m.progress.ViewAs(ratio) + "\n\n"
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
		s += "  "
		for i := 0; i < c; i++ {
			s += fmt.Sprintf("%d ", i)
		}
		s += "\n"
		for i := 0; i < r; i++ {
			s += fmt.Sprintf("%d ", i)
			for j := 0; j < c; j++ {
				if g.Board[i][j].IsMine {
					s += fmt.Sprintf("\033[38;5;88mX \033[0m")
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
		s += "\n\033[31mGAME OVER! Press 'q' to quit.\033[0m\n"

	}
	if g.GameWon {
		s += "\n\033[31mGAME Won! Press 'q' to quit.\033[0m\n"
	}
	return s
}
