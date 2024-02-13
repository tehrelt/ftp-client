package command

import (
	"fmt"
	"ftp-client/internal/session"
)

type PwdCommand struct {
	client *session.Session
}

func NewPwdCommand(s *session.Session) *PwdCommand {
	return &PwdCommand{
		client: s,
	}
}

func (cmd *PwdCommand) Execute(args []string) (STATUS, error) {

	dir, err := cmd.client.Pwd()
	if err != nil {
		return ERROR, err
	}

	fmt.Printf("current dir: %s\n", dir)

	return SUCCESS, nil
}
