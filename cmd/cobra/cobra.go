package cobra

import (
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/cli"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/dotfiles"
	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// root cmd represents the base command when called without any subcommands
var root = &cobra.Command{
	Short: "This program clones a GitHub repo and creates symbolic links for all files and directories it contains.",
	Use:   "dotfiles-installer",
	Run:   run,
	SuggestFor: []string{
		"dotfiles",
		"installer",
	},
	Version: cli.AppVersion,
}

var (
	url    string
	silent bool
	sudo   bool
)

func Execute() error {
	return root.Execute()
}

func init() {
	root.PersistentFlags().StringVarP(&url, "url", "u", "", "URL of the dotfile repo (optional, will be prompted if not provided)")
	root.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Run in silent mode (optional)")
	root.PersistentFlags().BoolVarP(&sudo, "force", "f", false, "Run as sudo (optional)")
}

func run(cmd *cobra.Command, args []string) {
	c := &cli.CLI{}

	c.Flags.URL = url
	c.Flags.Silent = silent
	c.Flags.Sudo = sudo

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
