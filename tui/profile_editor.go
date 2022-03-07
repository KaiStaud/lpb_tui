package tui

/*
* The profile editor allowes the user to change database entries during runtime.
* Stored profiles are displayed with a paginator.
* Items can be selected with the arrow keys and enter key.
* After selecting, the user can access profile operations on the appearing pop-up-menue.
 */

import (
	"lpb/storage"

	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title,
	desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func (m model) ViewList() string {
	return docStyle.Render(m.list.View())
}

// Send Profile into Waiting-Queue
func AddJobToQueue(name string) error {
	// Get index by name
	//err, tcp := storage.GetCoordinatesByName(name)
	err, id := storage.GetIDByName(name)

	if err == nil {
		jobqueue <- id
	}
	// Push info into queue
	return err
}
