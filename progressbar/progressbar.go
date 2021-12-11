/*
* Visualizes the current progress.
* Progress is getting calculated in module tracking and send to progressbar over channel.
* The progressbar is updated in leaping fashion.
* On Finish a tick poping up on the upper right side of the progressbar.
 */

package progressbar

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//----------------------- Constansts ----------------------- //

const (
	padding  = 2
	maxWidth = 80
)

//----------------------- Variables ----------------------- //

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
)

//----------------------- Functions ----------------------- //

// Receive new progress-Information via Channel:
func (m model) ReceiveProgess(new_progress float64) {
	if new_progress > 1 {
		m.percent = 1
	} else {
		m.percent = new_progress
	}
}

func Initialize() {
	prog := progress.NewModel(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))

	if err := tea.NewProgram(model{progress: prog}).Start(); err != nil {
		log.Printf("Bubbletea couldn't create progressbar : %v", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

type model struct {
	percent  float64
	progress progress.Model
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.percent >= 1.0 {
			return m, tea.Quit
		}
		return m, tickCmd()

	default:
		return m, nil
	}
}
func (_ model) Init() tea.Cmd {
	return tickCmd()
}

func (e model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + e.progress.ViewAs(e.percent) + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
