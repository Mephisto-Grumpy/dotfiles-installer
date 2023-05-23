package main

import (
	"github.com/Mephisto-Grumpy/dotfiles-installer/cmd/cobra"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/cli"
)

func main() {
	c := &cli.CLI{}

	if err := cobra.Execute(); err != nil {
		c.ShowError(err)
	}
}
