package command

import (
	"ftp-client/internal/session"
)

type LoginCommand struct {
	client *session.Session
}

func NewLoginCommand(s *session.Session) *LoginCommand {
	return &LoginCommand{
		client: s,
	}
}

func (cmd *LoginCommand) Execute(args []string) (STATUS, error) {

	if len(args) < 2 {
		return ERROR, ErrArgs
	}

	if err := cmd.client.Login(args[0], args[1]); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
