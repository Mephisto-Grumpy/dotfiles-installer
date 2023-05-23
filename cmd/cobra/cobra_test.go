package cobra

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestHelpMessage(t *testing.T) {
	helpMsg := captureOutput(func() { _ = root.Help() })

	if !strings.Contains(helpMsg, "Usage") || !strings.Contains(helpMsg, "Flags") || !strings.Contains(helpMsg, "This program clones a GitHub repo and creates symbolic links for all files and directories it contains.") {
		t.Errorf("Help message did not contain expected sections. \nHelp message: %s", helpMsg)
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
