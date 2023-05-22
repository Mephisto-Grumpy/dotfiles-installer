package dotfiles

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
)

var askConfirmation = func(message string) bool {
	templates := &promptui.PromptTemplates{
		Prompt:  "\U0001F447" + " " + "{{ . | cyan }}" + " ",
		Valid:   "\U00002705" + " " + "{{ . | green }}" + " ",
		Invalid: "\U0000274C" + " " + "{{ . | red }}" + " ",
		Success: "\U0001F389" + " " + "{{ . | bold }}" + " ",
	}

	promptAsk := promptui.Prompt{
		Label: message,
		Validate: func(input string) error {
			lowerInput := strings.ToLower(input)
			if lowerInput != "y" && lowerInput != "n" {
				return errors.New("input must be 'Y' or 'N'")
			}
			return nil
		},
		Templates: templates,
	}

	confirmation, err := promptAsk.Run()
	if err != nil {
		return false
	}

	if strings.ToLower(confirmation) != "y" {
		return false
	}

	return true
}

func Install(e utils.Executor, fs utils.Filesystem, url string, silent bool, sudo bool) error {
	dir := filepath.Join(os.TempDir(), "dotfiles")

	if _, err := os.Stat(dir); err == nil {
		shouldReplace := askConfirmation("Directory " + dir + " exists. Do you want to replace it? (Y/n): ")

		if !shouldReplace {
			return errors.New("directory already exists and user chose not to replace")
		}

		if err := fs.RemoveAll(dir); err != nil {
			return fmt.Errorf("failed to remove existing directory: %w", err)
		}
	}

	err := e.RunCmd("git", "clone", "--depth=1", url, dir)
	if err != nil {
		return fmt.Errorf("failed to clone repo: %w", err)
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		if rel == ".git" || info.IsDir() {
			return nil
		}

		dst := filepath.Join(os.Getenv("HOME"), rel)

		parentDir := filepath.Dir(dst)
		if _, err := os.Stat(parentDir); os.IsNotExist(err) {
			if err = os.MkdirAll(parentDir, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}

		if _, err := os.Stat(dst); !os.IsNotExist(err) {
			if !askConfirmation(fmt.Sprintf("File %s exists. Do you want to replace it?", dst)) {
				return nil
			}
		}

		if !silent {
			pterm.Info.Printf("Creating symlink: %s\n", dst)
		}

		if sudo {
			cmd := exec.Command("sudo", "ln", "-sf", path, dst)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				return fmt.Errorf("failed to create symlink: %w", err)
			}
		} else {
			cmd := exec.Command("ln", "-sf", path, dst)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				pterm.Warning.Println("Command failed, trying with sudo...")
				cmd = exec.Command("sudo", "ln", "-sf", path, dst)
				err = cmd.Run()
				if err != nil {
					return fmt.Errorf("failed to create symlink: %w", err)
				}
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to create symlinks: %w", err)
	}

	return nil
}
