package command

import (
	"ftp-client/internal/session"
	"strings"
)

type CwdCommand struct {
	client *session.Session
}

func NewCwdCommand(s *session.Session) *CwdCommand {
	return &CwdCommand{
		client: s,
	}
}

func (cmd *CwdCommand) Execute(args []string) (STATUS, error) {

	if len(args) < 1 {
		return ERROR, ErrArgs
	}

	if err := cmd.client.Cwd(strings.Join(args[0:], " ")); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
