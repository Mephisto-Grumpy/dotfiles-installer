package main

import (
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/cli"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/dotfiles"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/utils"
	"github.com/pterm/pterm"
)

func main() {
	c := &cli.CLI{}

	c.ParseFlags()

	if c.Help {
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

	if !c.Silent {
		pterm.Info.Printf("Cloning dotfiles from: %s\n", c.URL)
	}

	err = dotfiles.Install(e, fs, c.URL, c.Silent)
	if err != nil {
		c.ShowError(err)
		return
	}
}
