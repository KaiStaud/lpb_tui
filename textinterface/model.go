package tui

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
}
