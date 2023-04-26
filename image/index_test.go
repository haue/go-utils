package image

import (
	"fmt"
	"os"
	"testing"
)

func TestCompressImageBytes(t *testing.T) {
	filePath := "1680782914302839.png"
	_, err := os.Stat(filePath)
	if err != nil {
		return
	}
	input, _ := os.ReadFile(filePath)
	if len(input) > (1 << 20) {
		output, err := CompressImageBytes(input, 50)
		if err != nil {
			fmt.Println(err)
		}
		os.WriteFile("compressed.jpg", output, 0644)
	}
	fmt.Println("Image compression completed.")
}
