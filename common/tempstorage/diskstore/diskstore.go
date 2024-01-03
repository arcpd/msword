package diskstore

import (
	_gb "github.com/arcpd/msword/common/tempstorage"
	_gg "io/ioutil"
	_b "os"
	_a "strings"
)

// TempFile creates a new temp directory by calling ioutil TempDir
func (_ba diskStorage) TempDir(pattern string) (string, error) { return _gg.TempDir("", pattern) }

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage() { _e := diskStorage{}; _gb.SetAsStorage(&_e) }

// RemoveAll removes all files in the directory
func (_d diskStorage) RemoveAll(dir string) error {
	if _a.HasPrefix(dir, _b.TempDir()) {
		return _b.RemoveAll(dir)
	}
	return nil
}

// Add is not applicable in the diskstore implementation
func (_ee diskStorage) Add(path string) error { return nil }

type diskStorage struct{}

// Open opens file from disk according to a path
func (_bd diskStorage) Open(path string) (_gb.File, error) {
	return _b.OpenFile(path, _b.O_RDWR, 0644)
}

// TempFile creates a new temp file by calling ioutil TempFile
func (_ec diskStorage) TempFile(dir, pattern string) (_gb.File, error) {
	return _gg.TempFile(dir, pattern)
}
