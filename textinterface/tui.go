/*
* Package handles text user interface to LPBs profile controller.
* The UI is split into a curses-like main screen and a seperate operational screen.
* Information between tui and other program components is shared over channels;
* Stateful goroutines prevent simultaneous access to data storage.
 */

package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"

	menue_options = 4
	max_choices   = 4
)

// General stuff for styling the view
var (
	term          = termenv.ColorProfile()
	keyword       = makeFgStyle("211")
	subtle        = makeFgStyle("241")
	progressEmpty = subtle(progressEmptyChar)
	dot           = colorFg(" • ", "236")

	// Gradient colors we'll use for the progress bar
	ramp = makeRamp("#B14FFF", "#00FFA3", progressBarWidth)
)

func Launch() {
	initialModel := model{0, false, 0, false, 10, 0, 0, false, false}
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

type tickMsg struct{}
type frameMsg struct{}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return tick()
}

// The ELM Execution Cycle consists of a Update-Function followed by a View-Function

// Main update function.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}
	// Handle Selections  and Animations differently:

	// Is A selection made?
	if !m.Chosen {
		return updateHandler(msg, m)
	}

	// There are new animation request:
	return updateChosen(msg, m)
}

/* View is preformed after Update to display the changes
Due to complexity View itself is spit up into seperate subviews, which are executed when selected
by their hierarchialy superior topview.
*/
func (m model) View() string {
	s := viewHandler(m)
	return indent.String("\n"+s+"\n\n", 2)
}

// Sub-View functions
func viewHandler(m model) string {

	// Are there any changes in any level?
	if m.OptionChosen {
		if m.Chosen {
			return chosenView(m)
		} else {
			return choicesView(m)
		}
	} else {
		return terminalOptions(m)
	}
}

// Sub-update functions
func updateHandler(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		// User already ticked a  checkbox, we are no longer in top-menue
		if m.OptionChosen {
			switch msg.String() {
			case "j", "down":
				if m.Choice < max_choices {
					m.Choice += 1
				}
			case "k", "up":
				if m.Choice > 0 {
					m.Choice -= 1
				}
			case "enter":
				m.Chosen = true
			}
			// Still in top menue:
		} else {
			switch msg.String() {
			case "j", "down":
				// Catch out-of-range:
				if m.Option < menue_options {
					m.Option += 1
				}
			case "k", "up":
				// Catch out-of-range:
				if m.Option > 0 {
					m.Option -= 1
				}
			case "enter":
				m.OptionChosen = true
			}
		}
		return m, frame()

	case tickMsg:
		m.Ticks += 1
		return m, tick()
	}

	// Execute Animations after a selection:
	return m, nil
}
