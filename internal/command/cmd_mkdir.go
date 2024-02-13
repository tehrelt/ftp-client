package command

import (
	"ftp-client/internal/session"
)

type MkdirCommand struct {
	client *session.Session
}

func NewMkdirCommand(s *session.Session) *MkdirCommand {
	return &MkdirCommand{
		client: s,
	}
}

func (cmd *MkdirCommand) Execute(args []string) (STATUS, error) {

	dirName := GetFileName(args[0:])

	if err := cmd.client.Mkdir(dirName); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
