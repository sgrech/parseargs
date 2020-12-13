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
			option: "a",
		},
		command{
			option: "bravo",
		},
		command{
			option: "charlie",
			value:  "cvalue",
		},
	}

	if _, ok := c.findCommand("a"); !ok {
		t.Fatalf("Expected command \"a\" to be found")
	}

	if _, ok := c.findCommand("bravo"); !ok {
		t.Fatalf("Expected command \"bravo\" to be found")
	}

	if _, ok := c.findCommand("charlie"); !ok {
		t.Fatalf("Expected command \"charlie\" to be found")
	}
}

func TestParseLongOption(t *testing.T) {
	if opt, _ := parseLongOption("alpha"); opt != "alpha" {
		t.Fatalf("Expected %s but got %s", "alpha", opt)
	}

	if opt, val := parseLongOption("beta=betaval"); opt != "beta" && val != "betaval" {
		t.Fatalf("Expected %s option and %s value but got %s option && %s value", "beta", "betaval", opt, val)
	}
}

func TestParseShortOption(t *testing.T) {
	opts := parseShortOption("a")
	if len(opts) != 1 {
		t.Fatalf("Expected slice of len 1 but got %d", len(opts))
	}

	if opts[0] != "a" {
		t.Fatalf("Expected first element to be \"a\" but got \"%s\"", opts[0])
	}

	opts = parseShortOption("abc")

	if len(opts) != 3 {
		t.Fatalf("Expected slice of len 3 but got %d", len(opts))
	}

	if opts[0] != "a" && opts[1] != "b" && opts[2] != "c" {
		t.Fatalf("Expected first three elements to be \"a\" \"b\" \"c\" but got \"%s\" \"%s\" \"%s\"", opts[0], opts[1], opts[2])
	}
}

func TestParseArgs(t *testing.T) {
	if cmds, err := ParseArgs("-a"); err == nil {
		if len(cmds) != 1 {
			t.Fatalf("Expected slice of len 1 but got %d", len(cmds))
		}

		if cmds[0].option != "a" {
			t.Fatalf("Expected command with option \"a\"  but got \"%s\"", cmds[0].option)
		}

	} else {
		t.Fatalf("Unexpected error %v", err)
	}
}
