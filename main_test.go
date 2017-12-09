package main

import (
	"testing"
)

func TestGoGetArgs(t *testing.T) {
	var ggArgs = goGetArgs{
		{false, "a", ""},
		{false, "b", ""},
	}

	if args := ggArgs.parseArgs(); len(args) > 0 {
		t.Errorf("length of args should be 0, got %d", len(args))
	}
	for i := range ggArgs {
		ggArgs[i].value = true
	}
	if args := ggArgs.parseArgs(); len(args) != len(ggArgs) {
		t.Errorf("length of args should be %d, got %d", len(ggArgs), len(args))
	}
}
