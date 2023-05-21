package utils

import "os"

type Filesystem interface {
	MkdirAll(path string, perm os.FileMode) error
	RemoveAll(path string) error
}

type OSFilesystem struct{}

func (fs OSFilesystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (fs OSFilesystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}
