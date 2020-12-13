package parseargs

import (
	"errors"
	"fmt"
	"strings"
)

type commands []command

type command struct {
	shortOption string
	longOption  string
	value       string
	enabled     bool
}

func (c commands) findCommand(opt string) (command, bool) {
	for _, comm := range c {
		if comm.shortOption == opt || comm.longOption == opt {
			return comm, true
		}
	}
	return command{}, false
}

func stripDashes(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func parseArgs(args string) (commands, error) {
	s := strings.Split(args, " ")
	c := commands{}
	for _, arg := range s {
		if strings.HasPrefix(arg, "--") {
			c = append(c, command{
				longOption: arg,
				enabled:    true,
			})
		} else if strings.HasPrefix(arg, "-") {
			c = append(c, command{
				shortOption: arg,
				enabled:     true,
			})
		} else {
			return nil, errors.New(fmt.Sprintf("Arg %s has no dash prefix", arg))
		}
	}
	return nil, nil
}

// func (c *commands) RegisterArgs(args string) {

// }
