package parseargs

import (
	"testing"
)

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

	if _, ok := c.FindCommand("a"); !ok {
		t.Fatalf("Expected command \"a\" to be found")
	}

	if _, ok := c.FindCommand("bravo"); !ok {
		t.Fatalf("Expected command \"bravo\" to be found")
	}

	if _, ok := c.FindCommand("charlie"); !ok {
		t.Fatalf("Expected command \"charlie\" to be found")
	}
}

func TestParseArgs(t *testing.T) {
	if cmds, err := ParseArgs([]string{"/process/main", "-a"}); err == nil {
		if len(cmds) != 1 {
			t.Fatalf("Expected slice of len 1 but got %d", len(cmds))
		}

		if cmds[0].option != "a" {
			t.Fatalf("Expected command with option \"a\"  but got \"%s\"", cmds[0].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"process/main", "-a", "-b"}); err == nil {
		if len(cmds) != 2 {
			t.Fatalf("Expected slice of len 2 but got %d", len(cmds))
		}

		if cmds[0].option != "a" {
			t.Fatalf("Expected command with option \"a\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[1].option != "b" {
			t.Fatalf("Expected command with option \"b\"  but got \"%s\"", cmds[1].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"process/main", "-ab"}); err == nil {
		if len(cmds) != 2 {
			t.Fatalf("Expected slice of len 2 but got %d", len(cmds))
		}

		if cmds[0].option != "a" {
			t.Fatalf("Expected command with option \"a\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[1].option != "b" {
			t.Fatalf("Expected command with option \"b\"  but got \"%s\"", cmds[1].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"process/main", "-ab", "-c"}); err == nil {
		if len(cmds) != 3 {
			t.Fatalf("Expected slice of len 3 but got %d", len(cmds))
		}

		if cmds[0].option != "a" {
			t.Fatalf("Expected command with option \"a\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[1].option != "b" {
			t.Fatalf("Expected command with option \"b\"  but got \"%s\"", cmds[1].option)
		}

		if cmds[2].option != "c" {
			t.Fatalf("Expected command with option \"c\"  but got \"%s\"", cmds[2].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"process/main", "--alpha"}); err == nil {
		if len(cmds) != 1 {
			t.Fatalf("Expected slice of len 1 but got %d", len(cmds))
		}

		if cmds[0].option != "alpha" {
			t.Fatalf("Expected command with option \"alpha\"  but got \"%s\"", cmds[0].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"process/main", "--alpha", "--bravo"}); err == nil {
		if len(cmds) != 2 {
			t.Fatalf("Expected slice of len 2 but got %d", len(cmds))
		}

		if cmds[0].option != "alpha" {
			t.Fatalf("Expected command with option \"alpha\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[1].option != "bravo" {
			t.Fatalf("Expected command with option \"bravo\"  but got \"%s\"", cmds[1].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"process/main", "-ab", "--charlie"}); err == nil {
		if len(cmds) != 3 {
			t.Fatalf("Expected slice of len 3 but got %d", len(cmds))
		}

		if cmds[0].option != "a" {
			t.Fatalf("Expected command with option \"a\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[1].option != "b" {
			t.Fatalf("Expected command with option \"b\"  but got \"%s\"", cmds[1].option)
		}

		if cmds[2].option != "charlie" {
			t.Fatalf("Expected command with option \"charlie\"  but got \"%s\"", cmds[2].option)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"/process/main", "--alpha=go"}); err == nil {
		if len(cmds) != 1 {
			t.Fatalf("Expected slice of len 1 but got %d", len(cmds))
		}

		if cmds[0].option != "alpha" {
			t.Fatalf("Expected command with option \"alpha\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[0].value != "go" {
			t.Fatalf("Expected command with value \"go\"  but got \"%s\"", cmds[0].value)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}

	if cmds, err := ParseArgs([]string{"/process/main", "--alpha=go", "--beta=go"}); err == nil {
		if len(cmds) != 2 {
			t.Fatalf("Expected slice of len 2 but got %d", len(cmds))
		}

		if cmds[0].option != "alpha" {
			t.Fatalf("Expected command with option \"alpha\"  but got \"%s\"", cmds[0].option)
		}

		if cmds[0].value != "go" {
			t.Fatalf("Expected command with value \"go\"  but got \"%s\"", cmds[0].value)
		}

		if cmds[1].option != "beta" {
			t.Fatalf("Expected command with option \"beta\"  but got \"%s\"", cmds[1].option)
		}

		if cmds[1].value != "go" {
			t.Fatalf("Expected command with value \"go\"  but got \"%s\"", cmds[1].value)
		}
	} else {
		t.Fatalf("Unexpected error %v", err)
	}
}
