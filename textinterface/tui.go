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

	"github.com/charmbracelet/bubbles/list"
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

	profile_editor = 1
	run_options    = 2
	test_results   = 3
	shutdown       = 4
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

	defaultWidth = 20
	defaultHight = 14
)

func Launch() {

	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "Bitter melon", desc: "It cools you down"},
		item{title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{title: "Eight hours of sleep", desc: "I had this once"},
		item{title: "Cats", desc: "Usually"},
		item{title: "Plantasia, the album", desc: "My plants love it too"},
		item{title: "Pour over coffee", desc: "It takes forever to make though"},
		item{title: "VR", desc: "Virtual reality...what is there to say?"},
		item{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		item{title: "Linux", desc: "Pretty much the best OS"},
		item{title: "Business school", desc: "Just kidding"},
		item{title: "Pottery", desc: "Wet clay is a great feeling"},
		item{title: "Shampoo", desc: "Nothing like clean hair"},
		item{title: "Table tennis", desc: "It’s surprisingly exhausting"},
		item{title: "Milk crates", desc: "Great for packing in your extra stuff"},
		item{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		item{title: "Stickers", desc: "The thicker the vinyl the better"},
		item{title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		item{title: "Warm light", desc: "Like around 2700 Kelvin"},
		item{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		item{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		item{title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	initialModel := model{0, false, 0, false, 10, 0, 0, false, false, list.NewModel(items, list.NewDefaultDelegate(), defaultWidth, defaultHight), nil, ""}
	initialModel.list.Title = "My Fave Things"
	p := tea.NewProgram(initialModel)
	p.EnterAltScreen()

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
		if k == "q" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
		if k == "esc" {
			m.OptionChosen = false
			m.Option = 0
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

	// Are there any changes in any level?
	if m.OptionChosen {
		if m.Option == 1 {
			if m.Chosen {
				return chosenView(m)
			} else {
				return choicesView(m)
			}
		} else if m.Option == 0 {
			return m.ViewList()
		} else {
			return terminalOptions(m)
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
			if m.Option != 0 {
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
			} else {
				return m.UpdateList(msg)
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
	}

	// Execute Animations after a selection:
	return m, nil
}
