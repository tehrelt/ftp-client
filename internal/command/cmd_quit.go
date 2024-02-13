package command

import (
	"fmt"
	"ftp-client/internal/session"
)

type QuitCommand struct {
	client *session.Session
}

func NewQuitCommand(s *session.Session) *QuitCommand {
	return &QuitCommand{
		client: s,
	}
}

func (cmd *QuitCommand) Execute(args []string) (STATUS, error) {

	if cmd.client.IsOpen {
		if err := cmd.client.Disconnect(); err != nil {
			return ERROR, err
		}
	}

	fmt.Println("Bye")

	return EXIT, nil
}
