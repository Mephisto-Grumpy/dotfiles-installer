package dotfiles

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Mephisto-Grumpy/dotfiles-installer/pkg/utils"
	"github.com/pterm/pterm"
)

func Install(e utils.Executor, fs utils.Filesystem, url string, silent bool) error {
	dir := filepath.Join(os.TempDir(), "dotfiles")

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

		if rel == "." || rel == ".git" {
			return nil
		}

		dst := filepath.Join(os.Getenv("HOME"), rel)

		if !silent {
			pterm.Info.Printf("Creating symlink: %s\n", dst)
		}

		return e.RunCmd("ln", "-s", path, dst)
	})

	if err != nil {
		return fmt.Errorf("failed to create symlinks: %w", err)
	}

	return nil
}
