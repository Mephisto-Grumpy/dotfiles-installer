package main

import (
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/cli"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/dotfiles"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/utils"
	"github.com/pterm/pterm"
)

func main() {
	c := &cli.CLI{}

	c.Flags.ParseFlags()

	if c.Flags.Help {
		c.ShowHelp()
		return
	}

	err := c.PromptURL()
	if err != nil {
		c.ShowError(err)
		return
	}

	e := &utils.CmdExecutor{}
	fs := &utils.OSFilesystem{}

	if !c.Flags.Silent {
		pterm.Info.Printf("Cloning dotfiles from: %s\n", c.Flags.URL)
	}

	if c.Flags.Sudo {
		pterm.Info.Println("Running as sudo")
	}

	err = dotfiles.Install(e, fs, c.Flags.URL, c.Flags.Silent, c.Flags.Sudo)
	if err != nil {
		c.ShowError(err)
		return
	}
}
