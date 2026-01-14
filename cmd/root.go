package cmd

import (
	"fmt"
	"minesweeper/internal/ui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var Prows int
var Pcols int
var Pmines int

var RootCmd = &cobra.Command{
	Use:   "minesweeper",
	Short: "play minesweeper",
	Long:  "minesweeper in your terminal",
	Run:   RootRun,
}

func RootRun(_ *cobra.Command, _ []string) {
	if Prows > 26 {
		fmt.Println("Rows should be atmost 26")
		os.Exit(1)
	}
	if Pcols > 52 {
		fmt.Println("Columns should be atmost 51")
		os.Exit(1)
	}
	if Pmines >= Prows*Pcols {
		fmt.Println("Too many mines")
		os.Exit(1)
	}
	p := tea.NewProgram(ui.InitialModel(Pmines, Prows, Pcols), tea.WithAltScreen())
	if _, err := p.Run(); err == nil {
		os.Exit(1)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
func init() {
	RootCmd.Flags().IntVarP(&Pmines, "mines", "m", 12, "no of mines on the board")
	RootCmd.Flags().IntVarP(&Prows, "rows", "r", 10, "no of rows on the board")
	RootCmd.Flags().IntVarP(&Pcols, "cols", "c", 10, "no of columns on the board")
}
