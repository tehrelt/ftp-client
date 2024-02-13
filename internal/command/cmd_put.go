package command

import (
	"bytes"
	"ftp-client/internal/session"
	"os"
)

type PutCommand struct {
	client *session.Session
}

func NewPutCommand(s *session.Session) *PutCommand {
	return &PutCommand{
		client: s,
	}
}

func (cmd *PutCommand) Execute(args []string) (STATUS, error) {

	fileName := GetFileName(args[0:])

	file, err := os.ReadFile(fileName)
	if err != nil {
		return ERROR, err
	}

	if err := cmd.client.Put(fileName, bytes.NewReader(file)); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}
