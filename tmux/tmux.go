package tmux

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tmux struct {
	socket string
	pid    int
	envvar string
}

// Function ParseEnv returns a struct containing information for
// a specific tmux session.
func ParseEnv() (*Tmux, error) {
	env := os.Getenv("TMUX")

	if env == "" {
		return nil, errors.New("Environment variable $TMUX is empty")
	}

	split := strings.Split(env, ":")

	pid, err := strconv.Atoi(split[1])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("PID in environment variable is not parseable: %v", err))
	}

	info := &Tmux{
		socket: split[0],
		pid:    pid,
		envvar: env,
	}

	return info, nil
}
