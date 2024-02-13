package command

import (
	"ftp-client/internal/session"
)

type RmCommand struct {
	client *session.Session
}

func NewRmCommand(s *session.Session) *RmCommand {
	return &RmCommand{
		client: s,
	}
}

func (cmd *RmCommand) Execute(args []string) (STATUS, error) {

	fileName := GetFileName(args[0:])

	if err := cmd.client.Remove(fileName); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
