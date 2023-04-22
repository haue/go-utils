package image

import (
	"bytes"
	"image"
	"image/jpeg"
)

func CompressImageBytes(imgBytes []byte, quality int) ([]byte, error) {
	// Decode image from byte array
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}

	// Create buffer to store compressed image
	var buffer bytes.Buffer

	// Set JPEG encoding options
	opts := &jpeg.Options{
		Quality: quality,
	}

	// Encode resized image to buffer
	err = jpeg.Encode(&buffer, img, opts)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
