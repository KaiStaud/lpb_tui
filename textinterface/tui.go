/*
* Package handles text user interface to LPBs profile controller.
* The UI is split into a curses-like main screen and a seperate operational screen.
* Information between tui and other program components is shared over channels;
* Stateful goroutines prevent simultaneous access to data storage.
 */

package tui

import (
	"fmt"
	"lpb/progressbar"

	tea "github.com/charmbracelet/bubbletea"
)

//----------------------- Constansts ----------------------- //
const subviews = 4

//----------------------- Variables ----------------------- //

// The model stores the current state of the tui.
type model struct {
	choices  []string // selectable options as in initialMode()
	cursor   int      // item the cursor is pointing at
	Frames   int
	selected map[int]struct{} // which items are selected
}

func initialModel() model {
	return model{
		// Operational modes:
		choices: []string{"Configuration", "Teaching", "Shutdown", "Running", "Test"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Which key was pressed
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {

				// Unselect  all previously selected items:
				for i := 0; i <= subviews; i++ {
					delete(m.selected, i)
				}
				// Select the item, where the cursor is pointing at
				m.selected[m.cursor] = struct{}{}
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Select operational mode:\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}
	// Send the UI for rendering
	return s
}

func Launch() {
	//p := tea.NewProgram(initialModel())
	progressbar.Initialize()
	/*
		if err := p.Start(); err != nil {
			fmt.Printf("terminal ui failed with error: %v", err)
			os.Exit(1)
		}
	*/
}
