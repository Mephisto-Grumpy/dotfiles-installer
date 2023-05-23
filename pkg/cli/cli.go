package cli

import (
	"errors"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/flag"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
)

type CLI struct {
	Flags flag.Flags
}

func (c *CLI) ShowHelp() {
	header := pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithFullWidth()
	header.Println("📦 " + AppName + " (version: " + AppVersion + ")")
	helpBuilder := new(strings.Builder)
	w := tabwriter.NewWriter(helpBuilder, 0, 0, 3, ' ', 0)

	fmt.Fprintf(w, "\n%s:\t%s [options]\n\n", pterm.Bold.Sprintf("Usage"), pterm.Cyan("📦 "+AppName+" (version: "+AppVersion+")"))
	fmt.Fprintln(w, pterm.Magenta("Options:"))
	fmt.Fprintf(w, "  %s\t\tURL of the dotfile repo (optional, will be prompted if not provided)\n", pterm.Green("-url <url>"))
	fmt.Fprintf(w, "  %s\t\tRun in silent mode (optional)\n", pterm.Green("-s"))
	fmt.Fprintf(w, "  %s\t\tRun as sudo (optional)\n", pterm.Green("-force"))
	fmt.Fprintf(w, "  %s\t\tShow help message\n", pterm.Green("-h"))
	fmt.Fprintf(w, "\n%s:\tThis program clones a GitHub repo and creates symbolic links for all files and directories it contains.\n", pterm.Yellow("Description"))

	w.Flush()

	fmt.Println(helpBuilder.String())
}

func (c *CLI) PromptURL() error {
	if c.Flags.URL != "" {
		return nil
	}

	promptURL := promptui.Prompt{
		Label:     "URL of the dotfile repo",
		Validate:  validateURL,
		Templates: customTemplates(),
	}

	url, err := promptURL.Run()
	if err != nil {
		return err
	}

	c.Flags.URL = url

	promptSilent := promptui.Prompt{
		Label:     "Do you want to run in silent mode? (Y/n)",
		Validate:  validateSilent,
		Templates: customTemplates(),
	}

	silent, err := promptSilent.Run()
	if err != nil {
		return err
	}

	c.Flags.Silent = strings.ToLower(silent) == "y"

	promptSudo := promptui.Prompt{
		Label:     "Do you want to run with sudo? (Y/n)",
		Validate:  validateSudo,
		Templates: customTemplates(),
	}

	sudo, err := promptSudo.Run()
	if err != nil {
		return err
	}

	c.Flags.Sudo = strings.ToLower(sudo) == "y"

	return nil
}

func validateURL(input string) error {
	if input == "" {
		return errors.New("URL must not be empty")
	}
	return nil
}

func validateSilent(input string) error {
	lowerInput := strings.ToLower(input)
	if lowerInput != "y" && lowerInput != "n" {
		return errors.New("input must be 'Y' or 'N'")
	}
	return nil
}

func validateSudo(input string) error {
	lowerInput := strings.ToLower(input)
	if lowerInput != "y" && lowerInput != "n" {
		return errors.New("input must be 'Y' or 'N'")
	}
	return nil
}

func (c *CLI) ShowError(err error) {
	pterm.Error.Println("An error occurred: " + err.Error())
}

func customTemplates() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Prompt:  "\U0001F447" + " " + "{{ . | cyan }}" + " ",
		Valid:   "\U00002705" + " " + "{{ . | green }}" + " ",
		Invalid: "\U0000274C" + " " + "{{ . | red }}" + " ",
		Success: "\U0001F389" + " " + "{{ . | bold }}" + " ",
	}
}
