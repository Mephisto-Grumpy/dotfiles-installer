package cli

import (
	"errors"
	"strings"

	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/flag"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
)

type Prompter interface {
	PromptUser(label string, validate promptui.ValidateFunc) (string, error)
}

type CLI struct {
	Flags flag.Flags
	Prompter
}

var customTemplates = &promptui.PromptTemplates{
	Prompt:  "\U0001F447" + " " + "{{ . | cyan }}" + " ",
	Valid:   "\U00002705" + " " + "{{ . | green }}" + " ",
	Invalid: "\U0000274C" + " " + "{{ . | red }}" + " ",
	Success: "\U0001F389" + " " + "{{ . | bold }}" + " ",
}

func (c *CLI) PromptUser(label string, validate promptui.ValidateFunc) (string, error) {
	prompt := promptui.Prompt{
		Label:     label,
		Validate:  validate,
		Templates: customTemplates,
	}

	return prompt.Run()
}

func validateYesNoInput(input string) error {
	lowerInput := strings.ToLower(input)
	if lowerInput != "y" && lowerInput != "n" {
		return errors.New("input must be 'Y' or 'N'")
	}
	return nil
}

func (c *CLI) PromptURL() error {
	var err error
	if c.Flags.URL == "" {
		c.Flags.URL, err = c.PromptUser("URL of the dotfile repo", validateURL)
		if err != nil {
			return err
		}
	}

	var response string
	response, err = c.PromptUser("Do you want to run in silent mode? (Y/n)", validateYesNoInput)
	if err != nil {
		return err
	}
	c.Flags.Silent = strings.ToLower(response) == "y"

	response, err = c.PromptUser("Do you want to run with sudo? (Y/n)", validateYesNoInput)
	if err != nil {
		return err
	}
	c.Flags.Sudo = strings.ToLower(response) == "y"

	return nil
}

func validateURL(input string) error {
	if input == "" {
		return errors.New("URL must not be empty")
	}
	return nil
}

func (c *CLI) ShowError(err error) {
	pterm.Error.Println("An error occurred: " + err.Error())
}
