package command

import (
	"ftp-client/internal/session"
	"strings"
)

type RmdirCommand struct {
	client *session.Session
}

func NewRmdirCommand(s *session.Session) *RmdirCommand {
	return &RmdirCommand{
		client: s,
	}
}

func (cmd *RmdirCommand) Execute(args []string) (STATUS, error) {

	var dirName string
	r := false

	if strings.Compare("-r", args[0]) == 0 {
		dirName = GetFileName(args[1:])

		r = true
	} else {
		dirName = GetFileName(args[0:])
	}

	if err := cmd.client.RmDir(dirName, r); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
