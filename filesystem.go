package php

import (
	"os"
	"path/filepath"
	"syscall"
)

// Dirname returns a parent directory's path
func Dirname(path string) string {
	return filepath.Dir(path)
}

// DirnameWithLevels returns a parent directory's path which is
// levels up from the current directory.
func DirnameWithLevels(path string, levels int) string {
	for i := 0; i < levels; i++ {
		path = Dirname(path)
	}
	return path
}

// Realpath returns canonicalized absolute pathname
func Realpath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}

// Touch creates the file if it does not exist
func Touch(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if f != nil {
		f.Close()
	}
	return err
}

// Unlink deletes a file
func Unlink(filename string) error {
	return syscall.Unlink(filename)
}

// Mkdir attempts to create the directory specified by pathname.
func Mkdir(pathname string, mode uint32, recursive bool) error {
	if mode == 0 {
		mode = 0777
	}
	var err error
	if recursive {
		err = os.MkdirAll(pathname, os.FileMode(mode))
	} else {
		err = os.Mkdir(pathname, os.FileMode(mode))
	}
	return err
}

// Rmdir removes empty directory
func Rmdir(dirname string) error {
	return syscall.Rmdir(dirname)
}
