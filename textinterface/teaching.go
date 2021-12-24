package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

/*
* Robots profiles are auto-generated during teaching.
* After Selecting the Teaching-Option from tui, the robot changes its state.
* Subsequently the user is asked for an not-given name for the new profile.
* After manually teaching the robots its new position,
* the user has to finish teaching by entering "f" into the console.
 */

const (
	ack_pending      = iota
	name_pending     = iota
	teaching_running = iota
	teaching_done    = iota
)

var (
	teaching_state = name_pending
)

/* Initialize Teaching TUI*/
func (m model) ViewTeaching() string {
	switch teaching_state {
	case name_pending:
		return m.ViewProfileName()
	case teaching_running:
		return m.ViewTeachingRunning()
	//case teaching_done:

	default:
		return ""
	}
}

func (m model) UpdateTeaching(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch teaching_state {

	case ack_pending:
		teaching_state = name_pending
		return m, nil

	case name_pending:
		return m.UpdateProfileName(msg)

	case teaching_running:

		teaching_state = teaching_running
		return m.UpdateTeachingRunning(msg)

	case teaching_done:
		//teaching_state = ack_pending
		m.OptionChosen = false
		m.Option = 0
		return m, nil
	default:
		teaching_state = ack_pending
		return m, nil
	}

}

/* Change Mode to Teaching, show tick and (changed_segments / total_segments) to show change progress */

/* If Timeout ocurrs, the tick is replaced with a red cross */

/* Get new profile name */
func (m model) ViewProfileName() string {
	return fmt.Sprintf(
		"New program is called:?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func (m model) UpdateProfileName(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			teaching_state = teaching_running
			// For now, just return to the top menue:
			// Later a handshake with the hardware is necessary to return from the teaching view
			//m.OptionChosen = false
			//m.Option = 0
			return m, nil //tea.Quit
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

/* Show active teaching with spinner */
func (m model) ViewTeachingRunning() string {
	str := fmt.Sprintf("\n\n   %s Teaching in progress\n\n", m.spinner.View())
	return str
}

func (m model) UpdateTeachingRunning(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

/* Show ! and error message if teaching was unexpectily quit */

/* Show "Teaching Done" on "f" input */

/* Create new entry in database */
