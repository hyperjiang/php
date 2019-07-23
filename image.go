package php

import (
	"bytes"
	"image"
	_ "image/gif"  // gif format
	_ "image/jpeg" // jpeg format
	_ "image/png"  // png format
	"mime"
	"os"
)

// ImageInfo stores the info of an image
type ImageInfo struct {
	Width  int
	Height int
	Format string
	Mime   string
}

// GetImageSize gets the size of an image
func GetImageSize(filename string) (ImageInfo, error) {

	var info ImageInfo

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return info, err
	}

	cfg, format, err := image.DecodeConfig(file)
	if err != nil {
		return info, err
	}

	info.Width = cfg.Width
	info.Height = cfg.Height
	info.Format = format
	info.Mime = mime.TypeByExtension("." + format)

	return info, nil
}

// GetImageSizeFromString gets the size of an image from a string
func GetImageSizeFromString(data []byte) (ImageInfo, error) {
	var info ImageInfo

	cfg, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return info, err
	}

	info.Width = cfg.Width
	info.Height = cfg.Height
	info.Format = format
	info.Mime = mime.TypeByExtension("." + format)

	return info, nil
}
