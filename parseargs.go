package parseargs

import (
	"errors"
	"fmt"
	"strings"
)

type commands []command

type command struct {
	option string
	value  string
}

func (c commands) findCommand(opt string) (value string, ok bool) {
	for _, comm := range c {
		if comm.option == opt {
			return comm.value, true
		}
	}
	return "", false
}

func stripDashes(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func parseLongOption(s string) (option string, value string) {
	if strings.Contains(s, "=") {
		kvp := strings.Split(s, "=")
		return kvp[0], kvp[1]
	}
	return s, ""
}

func parseShortOption(s string) (options []string) {
	if len(s) > 1 {
		return strings.Split(s, "")
	}
	return []string{s}
}

func ParseArgs(args string) (c commands, err error) {
	f := strings.Fields(args)
	fmt.Println(f)
	var comms commands
	for _, arg := range f {
		if strings.HasPrefix(arg, "--") {
			opt, val := parseLongOption(stripDashes(arg))
			comms = append(comms, command{
				option: opt,
				value:  val,
			})
		} else if strings.HasPrefix(arg, "-") {
			opts := parseShortOption(stripDashes(arg))
			fmt.Println(arg, strings.HasPrefix(arg, "-"), opts)
			for _, opt := range opts {
				comms = append(comms, command{
					option: opt,
				})
			}
		} else {
			return nil, errors.New(fmt.Sprintf("Arg %s has no dash prefix", arg))
		}
	}
	return comms, nil
}
