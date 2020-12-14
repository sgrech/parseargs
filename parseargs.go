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

func (c commands) FindCommand(opt string) (value string, ok bool) {
	for _, comm := range c {
		if comm.option == opt {
			return comm.value, true
		}
	}
	return "", false
}

func ParseArgs(args string) (c commands, err error) {
	f := strings.Fields(args)
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
