package cli

import (
	"os"
	"testing"

	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/flag"
)

func TestParseFlags(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "--url", "testurl", "--silent", "--force"}

	flags := &flag.Flags{}
	flags.ParseFlags()

	if flags.URL != "testurl" {
		t.Errorf("Expected URL 'testurl', got '%s'", flags.URL)
	}
	if flags.Silent != true {
		t.Errorf("Expected Silent to be true, got '%v'", flags.Silent)
	}
	if flags.Sudo != true {
		t.Errorf("Expected Sudo to be true, got '%v'", flags.Sudo)
	}
}
