package command

import "github.com/jlaffaye/ftp"

type NoopCommand struct {
}

func NewNoopCommand() *NoopCommand {
	return &NoopCommand{}
}

func (cmd *NoopCommand) Execute(client *ftp.ServerConn, args []string) (STATUS, error) {

	if err := client.NoOp(); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
