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

// Find if command is present in slice. Returns the value assigned
// to the command (if any) as well as a bool for quick checking if
// the command was found or not.
func (c commands) FindCommand(opt string) (value string, ok bool) {
	for _, comm := range c {
		if comm.option == opt {
			return comm.value, true
		}
	}
	return "", false
}

// Given a slice of type string, parse the slice for command line style
// arguments. You can directly pass os.Args and a struct slice will be
// be returned of type command. This struct has two properties, option
// and value, as well as having the FindCommand interface for quickly
// finding commands.
func ParseArgs(args []string) (c commands, err error) {
	var comms commands
	for _, arg := range args[1:] {
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
