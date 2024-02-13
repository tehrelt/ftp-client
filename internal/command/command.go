package command

import (
	"errors"
	"strings"
)

type STATUS int

const (
	SUCCESS STATUS = iota
	EXIT
	ERROR
)

var (
	ErrArgs = errors.New("invalid args")
)

type Command interface {
	Execute(args []string) (STATUS, error)
}

type CommandArgs struct {
	Handle string
	Args   []string
}

func NewArgs(prompt string) *CommandArgs {
	args := strings.Split(prompt, " ")
	return &CommandArgs{
		Handle: args[0],
		Args:   args[1:],
	}
}

func GetFileName(r []string) string {
	return strings.Join(r, " ")
}
