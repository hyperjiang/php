package php

import (
	"io"
	"io/ioutil"
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

// Mkdir attempts to create the directory specified by pathname
func Mkdir(pathname string, mode os.FileMode, recursive bool) error {
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

// Link create a hard link
func Link(target, link string) error {
	return os.Link(target, link)
}

// Chmod changes file mode
func Chmod(filename string, mode os.FileMode) error {
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

// IsLink tells whether the filename is a symbolic link
func IsLink(filename string) bool {
	fi, err := os.Lstat(filename)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}

// Copy copies the src file to dst, any existing file will be overwritten and will not copy file attributes
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// FileExists checks whether a file or directory exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}

// FileSize gets file size
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// Rename renames a file or directory
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// FilePutContents writes data to a file
func FilePutContents(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0666)
}

// FileGetContents reads entire file into a string
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}
