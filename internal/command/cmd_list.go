package command

import (
	"fmt"
	"ftp-client/internal/session"
	"time"
)

type ListCommand struct {
	client *session.Session
}

func NewListCommand(s *session.Session) *ListCommand {
	return &ListCommand{
		client: s,
	}
}

func (cmd *ListCommand) Execute(args []string) (STATUS, error) {

	entries, err := cmd.client.List()
	if err != nil {
		return ERROR, err
	}

	fmt.Print("type\tmodify time\t\tsize\tfile name\n")
	for _, e := range entries {
		fmt.Printf("%s\t%s\t%d\t%s\n", e.Type.String(), e.Time.Format(time.DateTime), e.Size, e.Name)
	}

	return SUCCESS, nil
}
