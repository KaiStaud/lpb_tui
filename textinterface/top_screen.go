package tui

/*
* The top screen wraps sub-functions into seperate selection menues.
* selection is made with curses-style checkboxes, which can be ticked and unticked.
 */

import (
	"fmt"
	"strconv"
)

// Select a subview from the top menue:
func terminalOptions(m model) string {
	c := m.Option

	tpl := "Select Option\n\n"
	tpl += "%s\n\n"
	tpl += "Up-Time in %s seconds\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Profile Editor", c == 0),
		checkbox("Run Options", c == 1),
		checkbox("Test Results", c == 2),
		checkbox("Shutdown", c == 3),
	)

	return fmt.Sprintf(tpl, choices, colorFg(strconv.Itoa(m.Ticks), "79"))
}
