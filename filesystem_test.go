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
