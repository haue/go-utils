package alg

import (
	"fmt"
	"testing"
)

func TestIdGen(t *testing.T) {
	id, err := IdGen()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%d\n", id)
}
