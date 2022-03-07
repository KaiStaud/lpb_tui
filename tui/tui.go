/*
* Package handles text user interface to LPBs profile controller.
* The UI is split into a curses-like main screen and a seperate operational screen.
* Information between tui and other program components is shared over channels;
* Stateful goroutines prevent simultaneous access to data storage.
 */

package tui

import (
	"fmt"
	"lpb/multilogger"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
)

const (
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"

	menue_options = 6 // Options in first menue
	max_choices   = 4 // Options in sub-menues, currently only 4 TODO:Use Slice instead!

	profile_editor = 0
	run_options    = 1
	test_results   = 2
	shutdown       = 3
	teaching       = 4
	simulation     = 5
)

// General stuff for styling the view
var (
	p             *tea.Program
	term          = termenv.ColorProfile()
	keyword       = makeFgStyle("211")
	subtle        = makeFgStyle("241")
	progressEmpty = subtle(progressEmptyChar)
	dot           = colorFg(" • ", "236")

	// Gradient colors we'll use for the progress bar
	ramp = makeRamp("#B14FFF", "#00FFA3", progressBarWidth)

	defaultWidth = 20
	defaultHight = 14

	//Expose the logging channel in module
	tuiLogs      chan<- string
	jobqueue     chan<- int
	idle_sim     chan<- bool
	movement_sim chan<- bool
)

// Create channels for sending data across programm
func StartJobQueue(queue chan<- int) {
	jobqueue = queue

}
func GetTui() *tea.Program {
	return p
}

// Initialize and launch textinerface
func Launch(idle chan<- bool, movement chan<- bool) *tea.Program {
	// Connect channels
	idle_sim = idle
	movement_sim = movement

	items := []list.Item{
		item{title: "Home", desc: "Home the robot"},
		item{title: "Shutdown", desc: "Poweroff System"},
		item{title: "Demo", desc: "Perform Demo"},
	}

	ti := textinput.NewModel()
	ti.Placeholder = "Unsaved Profile"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	s := spinner.NewModel()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	spinner.Tick()

	initialModel := model{0, false, 0, false, 10, 0, 0, false, false, list.NewModel(items, list.NewDefaultDelegate(), defaultWidth, defaultHight), nil, "", ti, s, nil}
	initialModel.list.Title = "Saved Profiles"
	p := tea.NewProgram((initialModel))
	p.EnterAltScreen()
	go func() {
		if err := p.Start(); err != nil {
			fmt.Println("could not start program:", err)
			os.Exit(1)
		}
	}()
	return p
}

type tickMsg struct{}
type frameMsg struct{}
type HandshakeMsg struct{ d string } //Signals a finished fsm

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
	return spinner.Tick
}

// The ELM Execution Cycle consists of a Update-Function followed by a View-Function

// Main update function.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
		if k == "esc" {
			m.OptionChosen = false
			m.Option = 0
		}
		if k == "a" {
			m.Chosen = false
		}
	}

	// FinshMessage?
	if _, ok := msg.(HandshakeMsg); ok {
		m.Chosen = false
		time.Sleep(time.Second * 2)
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
		if m.Option == run_options {
			if m.Chosen {
				return chosenView(m)
			} else {
				return choicesView(m)
			}
		} else if m.Option == profile_editor {
			if m.Chosen == false {
				return m.ViewList()
			} else {
				if m.err == nil {
					return m.ViewSucess("Added job to queue")

				} else {
					return m.ViewError(m.err)
				}
			}
		} else if m.Option == teaching {
			return m.ViewTeaching()
		} else if m.Option == simulation {
			if m.Chosen == false {
				return m.ViewSimulation()
			} else {
				return m.ViewSucess("Switched Simulation!")
			}
		} else {
			return terminalOptions(m)
		}

	} else {
		return terminalOptions(m)
	}
}

// Increments index of selected checkbox.
// Index is changed globally, no  need to implement it in xxx_Update!
// Support is limited to 2-layer hierarchy!
func updateHandler(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Keyevents in profile editor are handled seperately:
	if m.OptionChosen {
		switch m.Option {
		case profile_editor:
			m.list, cmd = m.list.Update(msg)

			switch msg := msg.(type) {

			// Parse keyboard inputs:
			case tea.KeyMsg:
				switch msg.String() {

				case "enter": // Get Selected Item:
					i, ok := m.list.SelectedItem().(item)
					if ok {
						m.list_choice = i.Title()
						err := AddJobToQueue(m.list_choice)
						s := fmt.Sprintf("Info:Added item %s to queue, returned %v", m.list_choice, err)
						multilogger.AddTuiLog(s)
						m.Chosen = true
						m.err = err
					}
				default:
					m.Chosen = false

				}
			}
			// If enter-key was pressed, add item to queue
			return m, cmd
		case teaching:
			return m.UpdateTeaching(msg)
		default:
		}
	}

	// All other events are generated on keypresses.
	// Check if which hierarchy they happened and update model:
	switch msg := msg.(type) {

	// Parse keyboard inputs:
	case tea.KeyMsg:

		// User already ticked a  checkbox, we are no longer in top-menue
		if m.OptionChosen {
			if m.Option != 0 { // Option 0 is mapped to Editor view!
				s := msg.String()

				switch msg.String() {
				// The number of choices depends on selected sub-menue:
				//max_choices
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
					if m.Option == simulation {
						m.UpdateSimulation(msg)
					}

				default:
					fmt.Println(s)
				}
			}
			// Still in top menue, increment / decrement option
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
	}

	// Execute Animations after a selection:
	return m, nil
}
