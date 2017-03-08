package tmux

// Function window creates a new window with the title s.
func NewWindow(s string) error {
	cmds := []string{"new-window"}
	if s != "" {
		cmds = append(cmds, "-n", s)
	}

	return cmd(args)
}

// Function split splits the current pane and applies the passed layout
func SplitWindow(layout string) error {
	err := cmd([]string{"split-window"})
	if err != nil {
		return err
	}

	return SelectLayout(layout)
}

// Function SelectLayout applies the passed layout to the current window
func SelectLayout(layout string) error {
	return cmd([]string{"select-layout", layout})
}
