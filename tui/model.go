package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
)

/*
* Export the shared model from tui.go to this file.
* Therefore its easier to find models members
 */
type model struct {
	Option       int // First menu level
	OptionChosen bool

	Choice int // Second menu level
	Chosen bool

	Ticks  int // Time in seconds
	Frames int

	Progress float64 // Movement-Progress in %
	Loaded   bool    // true when progress ==1
	Quitting bool

	list        list.Model
	list_items  []item
	list_choice string

	textInput textinput.Model
	spinner   spinner.Model
	err       error
}
