package php

import (
	"testing"
)

func TestJSONEncode(t *testing.T) {
	type args struct {
		value interface{}
	}
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"1",
			args{1},
			"1",
			false,
		},
		{
			"2",
			args{1.2},
			"1.2",
			false,
		},
		{
			"3",
			args{"abc"},
			"\"abc\"",
			false,
		},
		{
			"4",
			args{
				[]int{1, 2, 3, 4, 5},
			},
			"[1,2,3,4,5]",
			false,
		},
		{
			"5",
			args{
				map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
			},
			"{\"a\":1,\"b\":2,\"c\":3,\"d\":4,\"e\":5}",
			false,
		},
		{
			"6",
			args{
				group,
			},
			"{\"ID\":1,\"Name\":\"Reds\",\"Colors\":[\"Crimson\",\"Red\",\"Ruby\",\"Maroon\"]}",
			false,
		},
		{
			"7",
			args{nil},
			"null",
			false,
		},
		{
			"8",
			args{false},
			"false",
			false,
		},
		{
			"9",
			args{func() {}},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONEncode(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSONEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONDecode(t *testing.T) {
	var m = make(map[string]interface{})
	var a []int
	var s string
	type args struct {
		jsonStr string
		v       interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"1",
			args{"{\"a\":1,\"b\":2,\"c\":3,\"d\":4,\"e\":5}", &m},
			false,
		},
		{
			"2",
			args{"[1,2,3]", &a},
			false,
		},
		{
			"3",
			args{"\"123\"", &s},
			false,
		},
		{
			"4",
			args{"123", &s},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JSONDecode(tt.args.jsonStr, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("JSONDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
