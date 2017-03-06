package tmux

import "os/exec"

// Function window creates a new window with the title s.
func NewWindow(s string) error {
	cmds := []string{"new-window"}
	if s != "" {
		cmds = append(cmds, "-n", s)
	}

	c := exec.Command("tmux", cmds...)
	return c.Run()
}

// Function split splits the current pane and applies the passed layout
func SplitWindow(layout string) error {
	cmds := []string{"split-window"}
	c := exec.Command("tmux", cmds...)
	err := c.Run()
	if err != nil {
		return err
	}

	return SelectLayout(layout)
}

// Function SelectLayout applies the passed layout to the current window
func SelectLayout(layout string) error {
	cmds := []string{"select-layout", layout}
	c := exec.Command("tmux", cmds...)
	return c.Run()
}
