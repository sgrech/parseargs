package parseargs

import (
	"testing"
)

func TestStripDashes(t *testing.T) {
	sDash := stripDashes("-a")
	mDash := stripDashes("--a")
	if sDash != "a" {
		t.Fatalf("Expected \"a\" but got %s", sDash)
	}
	if mDash != "a" {
		t.Fatalf("Expected \"a\" but got %s", mDash)
	}
}

func TestFindCommand(t *testing.T) {
	c := commands{
		command{
			shortOption: "a",
			longOption:  "alpha",
		},
		command{
			shortOption: "b",
			longOption:  "bravo",
		},
		command{
			shortOption: "c",
			longOption:  "charlie",
		},
	}

	if _, ok := c.findCommand("a"); !ok {
		t.Fatalf("Expected arg \"a\" to be found")
	}

	if _, ok := c.findCommand("bravo"); !ok {
		t.Fatalf("Expected arg \"bravo\" to be found")
	}
}
