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

func TestParseLongOption(t *testing.T) {
	if opt, _ := parseLongOption("alpha"); opt != "alpha" {
		t.Fatalf("Expected %s but got %s", "alpha", opt)
	}

	if opt, val := parseLongOption("bravo=bravoval"); opt != "bravo" && val != "bravoval" {
		t.Fatalf("Expected %s option and %s value but got %s option && %s value", "bravo", "bravoval", opt, val)
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
