package php

import (
	"os"
)

// Getcwd gets the current working directory, it will return empty string when error occurs
func Getcwd() string {
	dir, _ := os.Getwd()
	return dir
}

// Chdir changes current directory to dir
func Chdir(dir string) error {
	return os.Chdir(dir)
}

// Scandir lists files and directories inside the specified path
func Scandir(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	names, err := f.Readdirnames(-1)
	f.Close()

	return names, err
}
