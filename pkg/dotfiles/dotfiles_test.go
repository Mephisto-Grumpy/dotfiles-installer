package dotfiles

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockExecutor struct {
	err error
}

func (m mockExecutor) RunCmd(command string, args ...string) error {
	return m.err
}

type mockFilesystem struct {
	err error
}

func (m mockFilesystem) MkdirAll(path string, perm os.FileMode) error {
	return m.err
}

func (m mockFilesystem) RemoveAll(path string) error {
	return m.err
}

func (m mockFilesystem) Stat(name string) (os.FileInfo, error) {
	return nil, m.err
}

func (m mockFilesystem) IsExist(err error) bool {
	return os.IsExist(err)
}

func TestInstall(t *testing.T) {
	t.Run("should return error when RunCmd fails", func(t *testing.T) {
		e := mockExecutor{err: errors.New("error")}
		fs := mockFilesystem{}

		err := Install(e, fs, "url", true, false)

		assert.Error(t, err)
	})

	t.Run("should return error when /tmp/dotfiles already exists and user chooses not to replace", func(t *testing.T) {
		dir := "/tmp/dotfiles"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(dir)

		e := mockExecutor{}
		fs := mockFilesystem{err: os.ErrExist}

		// Force askConfirmation to return false
		originalAskConfirmation := askConfirmation
		askConfirmation = func(message string) bool {
			return false
		}
		defer func() { askConfirmation = originalAskConfirmation }()

		err := Install(e, fs, "url", true, false)

		assert.Error(t, err)
	})

	t.Run("should return nil when /tmp/dotfiles already exists and user chooses to replace", func(t *testing.T) {
		dir := "/tmp/dotfiles"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(dir)

		e := mockExecutor{}
		fs := mockFilesystem{err: os.ErrExist}

		// Force askConfirmation to return true
		originalAskConfirmation := askConfirmation
		askConfirmation = func(message string) bool {
			return true
		}
		defer func() { askConfirmation = originalAskConfirmation }()

		err := Install(e, fs, "url", true, false)

		assert.Error(t, err)
	})

	t.Run("should return nil when RunCmd succeeds and /tmp/dotfiles doesn't exist", func(t *testing.T) {
		dir := "/tmp/dotfiles"
		os.RemoveAll(dir) // ensure /tmp/dotfiles doesn't exist

		e := mockExecutor{}
		fs := mockFilesystem{}

		err := Install(e, fs, "url", true, false)

		assert.Error(t, err)
	})
}
