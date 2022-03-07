package tui

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	sim_options = 4
	sim_move    = 0
	sim_idle    = 1
	sim_err     = 2
	sim_booting = 3
)

// Change simulated frames after selecting
func (m model) UpdateSimulation(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Unselect 'choosen' flag
	m.Option = 0
	m.OptionChosen = false
	// Return to top view
	return m, nil
}

// Show selected simulation mode
func (m model) ViewSimulation() string {
	c := m.Choice

	tpl := "select simulation\n\n"
	tpl += "%s\n\n"
	tpl += "Up-Time in %s seconds\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("esc: leave menue") + dot + subtle("q: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Movement", c == 0),
		checkbox("Idle", c == 1),
		checkbox("Error", c == 2),
		checkbox("Boot-Up", c == 3),
	)

	return fmt.Sprintf(tpl, choices, colorFg(strconv.Itoa(m.Ticks), "79"))
}
