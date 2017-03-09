package tmux

// NewWindow creates a new window with the title s.
func NewWindow(s string) error {
	args := []string{"new-window"}
	if s != "" {
		args = append(args, "-n", s)
	}

	return cmd(args)
}

// SplitWindow splits the current pane and applies the passed layout
func SplitWindow(layout string) error {
	err := cmd([]string{"split-window"})
	if err != nil {
		return err
	}

	return SelectLayout(layout)
}

// SelectLayout applies the passed layout to the current window
func SelectLayout(layout string) error {
	return cmd([]string{"select-layout", layout})
}
