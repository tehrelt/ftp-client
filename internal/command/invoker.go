package command

import (
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
			"help":       NewHelpCommand(),
			"connect":    NewConnectCommand(session),
			"disconnect": NewDisconnectCommand(session),
			"noop":       NewNoopCommand(session),
			"quit":       NewQuitCommand(session),
			"pwd":        NewPwdCommand(session),
			"login":      NewLoginCommand(session),
			"ls":         NewListCommand(session),
			"cd":         NewCwdCommand(session),
			"get":        NewGetCommand(session),
			"put":        NewPutCommand(session),
			"mkdir":      NewMkdirCommand(session),
			"rm":         NewRmCommand(session),
			"rmdir":      NewRmdirCommand(session),
		},
	}
}

func (i *Invoker) Execute(args *CommandArgs) (STATUS, error) {
	cmd, ok := i.commands[args.Handle]
	if !ok {
		return ERROR, fmt.Errorf("command '%s' not found. Try help", args.Handle)
	}
	return cmd.Execute(args.Args)
}
