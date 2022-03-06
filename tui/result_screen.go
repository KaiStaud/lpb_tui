package tui

import "fmt"

/*
Succes / Errorscreens for actions
*/

func (m model) ViewSucess(s string) string {
	return fmt.Sprintf("✔️  %s", s)
}

func (m model) ViewError(err error) string {
	tpl := subtle("esc : ack error") + dot + subtle("q: leave")
	return fmt.Sprintf("❌ Failed to add job, Error = %v \n", err) + tpl
}
