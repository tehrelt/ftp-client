package command

import "fmt"

type NoopCommand struct {
}

func NewHelpCommand() *NoopCommand {
	return &NoopCommand{}
}

func (cmd *NoopCommand) Execute(args []string) (STATUS, error) {

	fmt.Println("connect <ip> <port> - Подключение к FTP-серверу")
	fmt.Println("disconnect - Отключение")
	fmt.Println("noop - проверить подключение")
	fmt.Println("pwd - TODO")
	fmt.Println("ls - TODO")
	fmt.Println("cd <dir_name> - TODO")
	fmt.Println("get <file_name> - TODO")
	fmt.Println("put <file_name> - TODO")

	return SUCCESS, nil
}
