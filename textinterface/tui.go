/*
* Package handles text user interface to LPBs profile controller.
* The UI is split into a curses-like main screen and a seperate operational screen.
* Information between tui and other program components is shared over channels;
* Stateful goroutines prevent simultaneous access to data storage.
 */

package tui

import (
	"fmt"
	"strconv"
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
	// Which View needs to be changed?
	return updateHandler(msg, m)
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
	if m.OptionChosen {
		return choicesView(m)
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
	return m, nil
}

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "j", "down":
			m.Choice += 1
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice -= 1
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			m.Option = 1
			return m, frame()
		}

	case tickMsg:
		m.Ticks += 1
		return m, tick()
	}

	return m, nil
}

// Update loop for the second view after a choice has been made
func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	case frameMsg:
		if !m.Loaded {
			m.Frames += 1
			m.Progress += 0.025
			if m.Progress >= 1 {
				m.Frames = 0
				m.Progress = 0
				m.Loaded = false
				m.Chosen = false
				return m, tick()
			}
			return m, frame()
		}

	case tickMsg:
		if m.Loaded {
			m.Ticks += 1
			return m, tick()
		}
	}

	return m, nil
}

// Sub-views

// The first view, where you're choosing a task
func choicesView(m model) string {
	c := m.Choice

	tpl := "Select Mode\n\n"
	tpl += "%s\n\n"
	tpl += "Up-Time in %s seconds\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("esc: leave menue") + subtle("q: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Teaching", c == 0),
		checkbox("Operation", c == 1),
		checkbox("Shutdown", c == 2),
		checkbox("Test", c == 3),
	)

	return fmt.Sprintf(tpl, choices, colorFg(strconv.Itoa(m.Ticks), "79"))
}

// The second view, after a task has been chosen
func chosenView(m model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("Changing Mode to %s", keyword("Teaching"))
	case 1:
		msg = fmt.Sprintf("Changing Mode to %s", keyword("Teaching"))
	case 2:
		msg = fmt.Sprintf("Changing Mode to %s", keyword("Teaching"))
	default:
		msg = fmt.Sprintf("Changing Mode to %s. Communication Line will be changed to %s...", keyword("Testing"), keyword("Internal Loopback Mode"))
	}

	label := "Requesting OpChange..."

	return msg + "\n\n" + label + "\n" + progressbar(80, m.Progress) + "%"
}
