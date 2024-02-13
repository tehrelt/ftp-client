package command

import "github.com/jlaffaye/ftp"

type DisconnectCommand struct {
}

func NewDisconnectCommand() *DisconnectCommand {
	return &DisconnectCommand{}
}

func (cmd *DisconnectCommand) Execute(client *ftp.ServerConn, args []string) (STATUS, error) {

	if err := client.Quit(); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
