package command

import (
	"ftp-client/internal/session"
)

type DisconnectCommand struct {
	client *session.Session
}

func NewDisconnectCommand(s *session.Session) *DisconnectCommand {
	return &DisconnectCommand{
		client: s,
	}
}

func (cmd *DisconnectCommand) Execute(args []string) (STATUS, error) {

	if err := cmd.client.Disconnect(); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
