package php

import (
	"os"
	"path/filepath"
	"syscall"
)

// Basename returns trailing name component of path
func Basename(path string) string {
	return filepath.Base(path)
}

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

// Symlink creates a symbolic link
func Symlink(target, link string) error {
	return os.Symlink(target, link)
}

// Chmod changes file mode
func Chmod(filename string, mode uint32) error {
	return os.Chmod(filename, os.FileMode(mode))
}

// Chown changes file owner and group
func Chown(filename string, uid, gid int) error {
	return os.Chown(filename, uid, gid)
}

// IsFile tells whether the filename is a regular file
func IsFile(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

// IsDir tells whether the filename is a directory
func IsDir(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

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

// IsLink tells whether the filename is a symbolic link
func IsLink(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}
