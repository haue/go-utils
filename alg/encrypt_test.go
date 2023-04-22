package alg

import (
	"fmt"
	"testing"
)

func TestEncryptBytes(t *testing.T) {
	fmt.Printf("-----")
	input := "test"
	output, err := EncryptBytes("1234567890123456", []byte(input))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
