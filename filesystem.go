package php

import "path/filepath"

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
