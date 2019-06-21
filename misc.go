package php

import (
	"errors"
	"os"
	"runtime"
	"strings"
)

// Getenv gets the value of an environment variable
func Getenv(varname string) string {
	return os.Getenv(varname)
}

// Putenv sets the value of an environment variable.
// The setting should be a key-value pair, like "FOO=BAR"
func Putenv(setting string) error {
	s := strings.Split(setting, "=")
	if len(s) != 2 {
		return errors.New("Invalid setting: " + setting)
	}
	return os.Setenv(s[0], s[1])
}

// MemoryGetUsage returns the amount of allocated memory in bytes.
// Set realUsage to TRUE to get total memory allocated from system, including unused pages.
// If realUsage is FALSE then only the used memory is reported.
func MemoryGetUsage(realUsage bool) uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	if realUsage {
		return stat.Sys
	}
	return stat.Alloc
}
