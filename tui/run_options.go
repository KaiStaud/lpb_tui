package tui

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

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
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("esc: leave menue") + dot + subtle("q: quit")

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
		msg = fmt.Sprintf("Changing Mode to %s", keyword("Operation"))
	case 2:
		msg = fmt.Sprintf("Changing Mode to %s", keyword("Shutdown"))
	default:
		msg = fmt.Sprintf("Changing Mode to %s. Communication Line will be changed to %s...", keyword("Testing"), keyword("Internal Loopback Mode"))
	}

	label := "Requesting OpChange..."

	return msg + "\n\n" + label + "\n" + progressbar(80, m.Progress) + "%"
}
