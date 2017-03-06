package tmux

import (
	"os"
	"testing"
)

func TextParseEnv(t *testing.T) {
	os.Setenv("TMUX", "/tmp/tmux/socket:55")
	info, err := ParseEnv()
	if err != nil {
		t.Errorf("Received unexpected error: %v", err)
	}

	if info.socket != "/tmp/tmux/socket" {
		t.Errorf("Unexpected socket, expected /tmp/tmux/socket, got %s", info.socket)
	}

	os.Setenv("TMUX", "")
	info, err = ParseEnv()
	if err == nil {
		t.Errorf("Did not receive error, env var: %s", os.Getenv("TMUX"))
	}
}
