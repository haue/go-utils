package alg

import (
	"fmt"
	"testing"
)

func TestIdGen(t *testing.T) {
	for i := 0; i < 100; i++ {
		id, err := Id()
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("%d\n", id)
	}
}
