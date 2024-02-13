package command

import (
	"ftp-client/internal/session"
	"os"
)

type GetCommand struct {
	client *session.Session
}

func NewGetCommand(s *session.Session) *GetCommand {
	return &GetCommand{
		client: s,
	}
}

func (cmd *GetCommand) Execute(args []string) (STATUS, error) {

	fileName := GetFileName(args[0:])

	size, err := cmd.client.FileSize(fileName)
	if err != nil {
		return ERROR, err
	}

	response, err := cmd.client.Get(fileName)
	if err != nil {
		return ERROR, err
	}

	bytes := make([]byte, size)

	_, err = response.Read(bytes)
	if err != nil {
		return ERROR, err
	}

	os.WriteFile(fileName, bytes, 0644)

	return SUCCESS, nil
}

