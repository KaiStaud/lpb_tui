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
	Option       int
	OptionChosen bool
	Choice       int
	Chosen       bool
	Ticks        int
	Frames       int
	Progress     float64
	Loaded       bool
	Quitting     bool

	list        list.Model
	list_items  []item
	list_choice string

	textInput textinput.Model
	spinner   spinner.Model
	err       error
}
