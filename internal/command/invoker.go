package command

import (
	"errors"
	"fmt"

	"github.com/jlaffaye/ftp"
)

type Invoker struct {
	commands map[string]Command
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: map[string]Command{
			"connect":    NewConnectCommand(),
			"disconnect": NewDisconnectCommand(),
			"noop":       NewNoopCommand(),
		},
	}
}

func (i *Invoker) Execute(args *CommandArgs, client *ftp.ServerConn) (STATUS, error) {
	cmd, ok := i.commands[args.Handle]
	if !ok {
		return ERROR, errors.New(fmt.Sprintf("command '%s' not found. Try help", args.Handle))
	}
	return cmd.Execute(client, args.Args)
}
