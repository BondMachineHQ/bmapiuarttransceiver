package bmapiuarttransceiver

import (
	"context"
	"fmt"
	"testing"
)

func TestSource(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, dest := UartTransceiver(ctx, "/dev/ttyUSB1")

	i := 0
	for n := range dest {
		fmt.Println(i, n)
		if i == 10 {
			break
		}
		i++
	}
}
