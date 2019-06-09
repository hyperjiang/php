package php

import (
	"os"
	"testing"
)

func TestDirname(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"/etc/passwd"},
			"/etc",
		},
		{
			"2",
			args{"/etc"},
			"/",
		},
		{
			"3",
			args{"."},
			".",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dirname(tt.args.path); got != tt.want {
				t.Errorf("Dirname() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirnameWithLevels(t *testing.T) {
	type args struct {
		path   string
		levels int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"/usr/local/lib", 2},
			"/usr",
		},
		{
			"2",
			args{"/usr/local/lib", 0},
			"/usr/local/lib",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirnameWithLevels(tt.args.path, tt.args.levels); got != tt.want {
				t.Errorf("DirnameWithLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRealpath(t *testing.T) {

	dir, _ := os.Getwd()

	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"/etc/"},
			"/etc",
		},
		{
			"2",
			args{"/etc/.."},
			"/",
		},
		{
			"3",
			args{"."},
			dir,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Realpath(tt.args.path); got != tt.want {
				t.Errorf("Realpath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTouch(t *testing.T) {
	var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
	if got := Touch(filename); got != nil {
		t.Errorf("Touch() = %v, want %v", got, nil)
	}
	Unlink(filename)
}

func TestMkdir(t *testing.T) {
	var pathname = "/tmp/hyperjiangphpfilesystemtestdir"
	if got := Mkdir(pathname, 0666, false); got != nil {
		t.Errorf("Mkdir() = %v, want %v", got, nil)
	}
	// clean up
	Rmdir(pathname)

	var pathname2 = "/tmp/hyperjiangphpfilesystemtestdir/dir"
	if got := Mkdir(pathname2, 0, true); got != nil {
		t.Errorf("Mkdir() = %v, want %v", got, nil)
	}
	// clean up
	Rmdir(pathname2)
	Rmdir(pathname)
}

func TestIsFile(t *testing.T) {
	var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
	Unlink(filename)
	if IsFile(filename) { // not exists
		t.Fail()
	}
	Touch(filename)
	if !IsFile(filename) {
		t.Fail()
	}
	Unlink(filename)
}

func TestIsDir(t *testing.T) {
	var pathname = "/tmp/hyperjiangphpfilesystemtestdir"
	Rmdir(pathname)
	if IsDir(pathname) { // not exists
		t.Fail()
	}
	Mkdir(pathname, 0, false)
	if !IsDir(pathname) {
		t.Fail()
	}
	Rmdir(pathname)
}

func TestIsExecutable(t *testing.T) {
	var filename = "/tmp/hyperjiangphpfilesystemtest.exe"
	Unlink(filename)
	if IsExecutable(filename) { // not exists
		t.Fail()
	}
	Touch(filename)
	if IsExecutable(filename) { // not executable
		t.Fail()
	}
	Chmod(filename, 0777)
	if !IsExecutable(filename) {
		t.Fail()
	}
	Unlink(filename)
}

func TestIsReadable(t *testing.T) {
	var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
	Unlink(filename)
	if IsReadable(filename) { // not exists
		t.Fail()
	}
	Touch(filename)
	if !IsReadable(filename) {
		t.Fail()
	}
	Chmod(filename, 0222)
	if IsReadable(filename) { // not readable
		t.Fail()
	}
	Unlink(filename)
}

func TestIsWritable(t *testing.T) {
	var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
	Unlink(filename)
	if IsWritable(filename) { // not exists
		t.Fail()
	}
	Touch(filename)
	if !IsWritable(filename) {
		t.Fail()
	}
	Chmod(filename, 0555)
	if IsWritable(filename) { // not writable
		t.Fail()
	}
	Unlink(filename)
}

func TestIsLink(t *testing.T) {
	var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
	Touch(filename)
	defer Unlink(filename)

	var link = "/tmp/hyperjiangphpfilesystemtest.link"
	Unlink(link)
	if IsLink(link) { // not exists
		t.Fail()
	}

	if err := Symlink(filename, link); err != nil {
		t.Error("Create symbolic link failed")
	}
	if !IsLink(link) {
		t.Fail()
	}
	Unlink(link)
}
