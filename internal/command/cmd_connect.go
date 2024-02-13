package command

import (
	"fmt"
	"ftp-client/internal/session"
)

type ConnectCommand struct {
	client *session.Session
}

func NewConnectCommand(s *session.Session) *ConnectCommand {
	return &ConnectCommand{
		client: s,
	}
}

func (c *ConnectCommand) Execute(args []string) (STATUS, error) {

	if len(args) < 2 {
		fmt.Printf("usage: connect ip port\n")
		return ERROR, ErrArgs
	}
	if err := c.client.Connect(fmt.Sprintf("%s:%s", args[0], args[1])); err != nil {
		return ERROR, err
	}

	fmt.Printf("Connected to %s:%s\n", args[0], args[1])

	return SUCCESS, nil
}
