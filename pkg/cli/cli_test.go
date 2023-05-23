package cli

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/flag"
)

func TestParseFlags(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "--url", "testurl", "-s"}

	flags := &flag.Flags{}
	flags.ParseFlags()

	if flags.URL != "testurl" {
		t.Errorf("Expected URL 'testurl', got '%s'", flags.URL)
	}
	if flags.Silent != true {
		t.Errorf("Expected Silent to be true, got '%v'", flags.Silent)
	}
}

func TestShowHelp(t *testing.T) {
	cli := &CLI{}
	helpMsg := captureOutput(cli.ShowHelp)

	if !strings.Contains(helpMsg, "Usage") || !strings.Contains(helpMsg, "Options") || !strings.Contains(helpMsg, "Description") {
		t.Errorf("Help message did not contain expected sections")
	}
}

// captureOutput captures all output sent to standard output from the provided function.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf strings.Builder
	if _, err := io.Copy(&buf, r); err != nil {
		panic(err)
	}
	return buf.String()
}
