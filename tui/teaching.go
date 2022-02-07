package tui

import (
	"fmt"
	"lpb/multilogger"

	"github.com/charmbracelet/bubbles/spinner"
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
	case teaching_done:
		return m.ViewTeachingDone()
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
		teaching_state = ack_pending
		m.OptionChosen = false
		m.Option = 0
		return m.UpdateTeachingDone(msg)
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

/*
* Save entered string, register keypresses for ending teaching mode
 */
func (m model) UpdateProfileName(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// After pressing the enter button, teaching is changed to " actively running"
		case tea.KeyEnter:
			teaching_state = teaching_running
			// Restart the spinner, since it wasn't running after start ( wasn't started in teaching view,too!)
			return m, spinner.Tick
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

/* Show active teaching with spinner */
func (m model) ViewTeachingRunning() string {
	str := fmt.Sprintf("\n\n   %s Teaching in progress\n Press f to finish or d to quit\n", m.spinner.View())
	return str
}

func (m model) UpdateTeachingRunning(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	/* Was a command entered? */
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		// Teaching is done:
		case tea.KeyRunes:
			// Save buffered dataset
			if msg.String() == "f" {
				dbg_info := fmt.Sprintf("Info:created new profile %s", m.textInput.Value())
				multilogger.AddTuiLog(dbg_info)
				teaching_state = teaching_done
				return m, nil

				// Delete buffered dataset
			} else if msg.String() == "d" {
				multilogger.AddTuiLog("Info:Delete teaching-set")
				teaching_state = teaching_done
				return m, nil

				// Incorrect Input
			} else {
				multilogger.AddTuiLog("Info:Incorrect keypress while Teaching!")
				m.spinner, cmd = m.spinner.Update(msg)
				return m, cmd
			}
		// Delete Set:

		// Wrong input, just update spinner:
		default:
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	//	var cmd tea.Cmd
	//	m.spinner, cmd = m.spinner.Update(msg)
	//	return m, cmd
}

/* Show ! and error message if teaching was unexpectily quit */

/* Wait until user enters 'f' for finished teaching */
func (m model) UpdateTeachingDone(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

/* Show "Teaching Done" on "f" input */
func (m model) ViewTeachingDone() string {
	return "Teaching finished!"

}
