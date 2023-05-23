package flag

import (
	"github.com/spf13/pflag"
)

type Flags struct {
	URL    string
	Silent bool
	Help   bool
	Sudo   bool
}

func (f *Flags) ParseFlags() {
	pflag.StringVarP(&f.URL, "url", "u", "", "URL of the dotfile repo (optional, will be prompted if not provided)")
	pflag.BoolVarP(&f.Silent, "silent", "s", false, "Run in silent mode (optional)")
	pflag.BoolVarP(&f.Sudo, "force", "f", false, "Run as sudo (optional)")
	pflag.BoolVarP(&f.Help, "help", "h", false, "Show help message")
	pflag.Parse()
}
