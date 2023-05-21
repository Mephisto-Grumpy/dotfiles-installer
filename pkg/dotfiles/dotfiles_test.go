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

func TestInstall(t *testing.T) {
	t.Run("should return error when RunCmd fails", func(t *testing.T) {
		e := mockExecutor{err: errors.New("error")}
		fs := mockFilesystem{}

		err := Install(e, fs, "url", true)

		assert.Error(t, err)
	})

	t.Run("should return nil when RunCmd succeeds", func(t *testing.T) {
		dir := "/tmp/dotfiles"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(dir)

		e := mockExecutor{}
		fs := mockFilesystem{}

		err := Install(e, fs, "url", true)

		assert.Nil(t, err)
	})
}
