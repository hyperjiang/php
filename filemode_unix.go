// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package php

import "syscall"

// IsExecutable tells whether the filename is executable
func IsExecutable(filename string) bool {
	return syscall.Access(filename, 0x1) == nil
}

// IsReadable tells whether a file exists and is readable
func IsReadable(filename string) bool {
	return syscall.Access(filename, 0x4) == nil
}

// IsWritable tells whether the filename is writable
func IsWritable(filename string) bool {
	return syscall.Access(filename, 0x2) == nil
}
