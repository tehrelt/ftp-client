package terminal

import (
	"bufio"
	"fmt"
	"ftp-client/internal/command"
	"os"
	"strings"
)

type Terminal struct {
	cmdChannel  chan<- *command.CommandArgs
	waitChannel <-chan struct{}
}

func NewTerminal(cmd chan<- *command.CommandArgs, wait <-chan struct{}) *Terminal {
	return &Terminal{
		cmdChannel:  cmd,
		waitChannel: wait,
	}
}

func (t *Terminal) Prompt() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("unreal-ftp > ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

func (t *Terminal) Run() error {
	for {
		prompt, err := t.Prompt()
		if err != nil {
			return err
		}

		t.cmdChannel <- command.NewArgs(prompt)

		fmt.Print("wait channel is waiting\n")
		<-t.waitChannel
		fmt.Print("wait channel recieved\n")
	}
}
