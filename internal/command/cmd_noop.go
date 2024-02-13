package command

import (
	"fmt"
	"ftp-client/internal/session"
)

type NoopCommand struct {
	client *session.Session
}

func NewNoopCommand(s *session.Session) *NoopCommand {
	return &NoopCommand{
		client: s,
	}
}

func (cmd *NoopCommand) Execute(args []string) (STATUS, error) {

	if err := cmd.client.Noop(); err != nil {
		return ERROR, err
	}

	fmt.Printf("Connection alive\n")

	return SUCCESS, nil
}
