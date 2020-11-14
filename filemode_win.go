// +build windows

package php

import (
	"os"
	"path/filepath"
	"strings"
)

// IsExecutable tells whether the filename is executable
func IsExecutable(filename string) bool {
	if !IsReadable(filename) {
		return false
	}

	return strings.ToUpper(filepath.Ext(filename)) == ".EXE"
}

// IsReadable tells whether a file exists and is readable
func IsReadable(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return true
}

// IsWritable tells whether the filename is writable
func IsWritable(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}

	return fi.Mode() == 0666
}
