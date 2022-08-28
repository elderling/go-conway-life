package main

import (
	"testing"
)

func TestPrint(t *testing.T) {
	p := PlayField{{true, true, false},
		{false, false, false},
		{true, true}}

	s := p.String()

	if s != "00 \n   \n00\n" {
		t.Error("String() broken")
	}
}
