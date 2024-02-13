package command

import (
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
)

type ConnectCommand struct {
}

func NewConnectCommand() *ConnectCommand {
	return &ConnectCommand{}
}

func (c *ConnectCommand) Execute(client *ftp.ServerConn, args []string) (STATUS, error) {
	var err error

	if len(args) < 2 {
		fmt.Printf("usage: connect ip port\n")
		return ERROR, ErrArgs
	}

	if client, err = ftp.Dial(fmt.Sprintf("%s:%s", args[0], args[1]), ftp.DialWithTimeout(5*time.Second)); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
