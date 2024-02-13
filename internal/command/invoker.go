package command

import (
	"errors"
	"fmt"
	"ftp-client/internal/session"
)

type Invoker struct {
	session  *session.Session
	commands map[string]Command
}

func NewInvoker(session *session.Session) *Invoker {
	return &Invoker{
		session: session,
		commands: map[string]Command{
			"connect":    NewConnectCommand(session),
			"disconnect": NewDisconnectCommand(session),
			"noop":       NewNoopCommand(session),
			"quit":       NewQuitCommand(session),
			"help":       NewHelpCommand(),
		},
	}
}

func (i *Invoker) Execute(args *CommandArgs) (STATUS, error) {
	cmd, ok := i.commands[args.Handle]
	if !ok {
		return ERROR, errors.New(fmt.Sprintf("command '%s' not found. Try help", args.Handle))
	}
	return cmd.Execute(args.Args)
}
