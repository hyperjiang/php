package php

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGetimagesize(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    ImageInfo
		wantErr bool
	}{
		{
			"happy flow",
			args{"./fish.jpg"},
			ImageInfo{
				Width:  559,
				Height: 533,
				Format: "jpeg",
				Mime:   "image/jpeg",
			},
			false,
		},
		{
			"no such file",
			args{"./not-exist-file"},
			ImageInfo{},
			true,
		},
		{
			"unknown format",
			args{"./LICENSE"},
			ImageInfo{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetImageSize(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImageSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetImageSizeFromString(t *testing.T) {

	data1, _ := ioutil.ReadFile("./fish.jpg")
	data2, _ := ioutil.ReadFile("./LICENSE")

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    ImageInfo
		wantErr bool
	}{
		{
			"happy flow",
			args{data1},
			ImageInfo{
				Width:  559,
				Height: 533,
				Format: "jpeg",
				Mime:   "image/jpeg",
			},
			false,
		},
		{
			"unknown format",
			args{data2},
			ImageInfo{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetImageSizeFromString(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImageSizeFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetImageSizeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
