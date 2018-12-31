package php

import (
	"testing"
)

func TestCtypeAlnum(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"0"},
			true,
		},
		{
			"2",
			args{"53"},
			true,
		},
		{
			"3",
			args{"asdf"},
			true,
		},
		{
			"4",
			args{"ADD"},
			true,
		},
		{
			"5",
			args{"A1cbad"},
			true,
		},
		{
			"6",
			args{""},
			false,
		},
		{
			"7",
			args{"-127"},
			false,
		},
		{
			"8",
			args{"53.0"},
			false,
		},
		{
			"9",
			args{"asd df"},
			false,
		},
		{
			"10",
			args{"é"},
			false,
		},
		{
			"11",
			args{"!asdf"},
			false,
		},
		{
			"12",
			args{"\x00asdf"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeAlnum(tt.args.text); got != tt.want {
				t.Errorf("CtypeAlnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeAlpha(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"asdf"},
			true,
		},
		{
			"2",
			args{"ADD"},
			true,
		},
		{
			"3",
			args{"bAcbad"},
			true,
		},
		{
			"4",
			args{""},
			false,
		},
		{
			"5",
			args{"-127"},
			false,
		},
		{
			"6",
			args{"53.0"},
			false,
		},
		{
			"7",
			args{"asd df"},
			false,
		},
		{
			"8",
			args{"é"},
			false,
		},
		{
			"9",
			args{"!asdf"},
			false,
		},
		{
			"10",
			args{"\x00asdf"},
			false,
		},
		{
			"11",
			args{"13addfadsf2"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeAlpha(tt.args.text); got != tt.want {
				t.Errorf("CtypeAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeCntrl(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"\x00"},
			true,
		},
		{
			"2",
			args{"\n"},
			true,
		},
		{
			"3",
			args{"\r"},
			true,
		},
		{
			"4",
			args{""},
			false,
		},
		{
			"5",
			args{"-127"},
			false,
		},
		{
			"6",
			args{"53.0"},
			false,
		},
		{
			"7",
			args{"asd df"},
			false,
		},
		{
			"8",
			args{"é"},
			false,
		},
		{
			"9",
			args{"!asdf"},
			false,
		},
		{
			"10",
			args{"1234"},
			false,
		},
		{
			"11",
			args{"\x00asdf"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeCntrl(tt.args.text); got != tt.want {
				t.Errorf("CtypeCntrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeDigit(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"0"},
			true,
		},
		{
			"2",
			args{"53"},
			true,
		},
		{
			"3",
			args{"01234"},
			true,
		},
		{
			"4",
			args{""},
			false,
		},
		{
			"5",
			args{"-127"},
			false,
		},
		{
			"6",
			args{"53.0"},
			false,
		},
		{
			"7",
			args{"asd df"},
			false,
		},
		{
			"8",
			args{"é"},
			false,
		},
		{
			"9",
			args{"!asdf"},
			false,
		},
		{
			"10",
			args{"1234B"},
			false,
		},
		{
			"11",
			args{"\x00asdf"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeDigit(tt.args.text); got != tt.want {
				t.Errorf("CtypeDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeGraph(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"-127"},
			true,
		},
		{
			"2",
			args{"A1cbad"},
			true,
		},
		{
			"3",
			args{"LKA#@%.54"},
			true,
		},
		{
			"4",
			args{""},
			false,
		},
		{
			"5",
			args{"asd df"},
			false,
		},
		{
			"6",
			args{"é"},
			false,
		},
		{
			"7",
			args{"\r"},
			false,
		},
		{
			"8",
			args{"\n"},
			false,
		},
		{
			"9",
			args{"\t"},
			false,
		},
		{
			"10",
			args{"\x00asdf"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeGraph(tt.args.text); got != tt.want {
				t.Errorf("CtypeGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeLower(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"asdf"},
			true,
		},
		{
			"2",
			args{""},
			false,
		},
		{
			"3",
			args{"-127"},
			false,
		},
		{
			"4",
			args{"53.0"},
			false,
		},
		{
			"5",
			args{"asd df"},
			false,
		},
		{
			"6",
			args{"DD"},
			false,
		},
		{
			"7",
			args{"123"},
			false,
		},
		{
			"8",
			args{"é"},
			false,
		},
		{
			"9",
			args{"!asdf"},
			false,
		},
		{
			"10",
			args{"1234B"},
			false,
		},
		{
			"11",
			args{"\x00asdf"},
			false,
		},
		{
			"12",
			args{"\n"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeLower(tt.args.text); got != tt.want {
				t.Errorf("CtypeLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypePrint(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"-127!!"},
			true,
		},
		{
			"2",
			args{"Asd df2"},
			true,
		},
		{
			"3",
			args{"LKA#@%.54"},
			true,
		},
		{
			"4",
			args{""},
			false,
		},
		{
			"5",
			args{"\r\n\t"},
			false,
		},
		{
			"6",
			args{"é"},
			false,
		},
		{
			"7",
			args{"\x00asdf"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypePrint(tt.args.text); got != tt.want {
				t.Errorf("CtypePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypePunct(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"@@!#^$"},
			true,
		},
		{
			"2",
			args{"*&$()"},
			true,
		},
		{
			"3",
			args{""},
			false,
		},
		{
			"4",
			args{"ABasdk!@!$#"},
			false,
		},
		{
			"5",
			args{"\r\n\t"},
			false,
		},
		{
			"6",
			args{"é"},
			false,
		},
		{
			"7",
			args{"\x00asdf"},
			false,
		},
		{
			"8",
			args{"!@ # $"},
			false,
		},
		{
			"9",
			args{"123"},
			false,
		},
		{
			"10",
			args{"abc"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypePunct(tt.args.text); got != tt.want {
				t.Errorf("CtypePunct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeSpace(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{" "},
			true,
		},
		{
			"2",
			args{"\r\n\t"},
			true,
		},
		{
			"3",
			args{""},
			false,
		},
		{
			"4",
			args{"\x01"},
			false,
		},
		{
			"6",
			args{"é"},
			false,
		},
		{
			"7",
			args{"\x00asdf"},
			false,
		},
		{
			"8",
			args{"!@ # $"},
			false,
		},
		{
			"9",
			args{"123"},
			false,
		},
		{
			"10",
			args{"abc"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeSpace(tt.args.text); got != tt.want {
				t.Errorf("CtypeSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeUpper(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"ASDF"},
			true,
		},
		{
			"2",
			args{""},
			false,
		},
		{
			"3",
			args{"-127"},
			false,
		},
		{
			"4",
			args{"53.0"},
			false,
		},
		{
			"5",
			args{"ASD DF"},
			false,
		},
		{
			"6",
			args{"abc"},
			false,
		},
		{
			"7",
			args{"123"},
			false,
		},
		{
			"8",
			args{"é"},
			false,
		},
		{
			"9",
			args{"!ASDF"},
			false,
		},
		{
			"10",
			args{"1234B"},
			false,
		},
		{
			"11",
			args{"\x00ASDF"},
			false,
		},
		{
			"12",
			args{"\n"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeUpper(tt.args.text); got != tt.want {
				t.Errorf("CtypeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCtypeXdigit(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{"01234"},
			true,
		},
		{
			"2",
			args{"A4fD"},
			true,
		},
		{
			"3",
			args{""},
			false,
		},
		{
			"4",
			args{"-127"},
			false,
		},
		{
			"5",
			args{"53.0"},
			false,
		},
		{
			"6",
			args{"ASD DF"},
			false,
		},
		{
			"7",
			args{"hhh"},
			false,
		},
		{
			"8",
			args{"é"},
			false,
		},
		{
			"9",
			args{"!ASDF"},
			false,
		},
		{
			"10",
			args{"Z"},
			false,
		},
		{
			"11",
			args{"\x00ASDF"},
			false,
		},
		{
			"12",
			args{"\n"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtypeXdigit(tt.args.text); got != tt.want {
				t.Errorf("CtypeXdigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
