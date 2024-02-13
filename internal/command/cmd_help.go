package command

import "fmt"

type HelpCommand struct {
}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (cmd *HelpCommand) Execute(args []string) (STATUS, error) {

	fmt.Println("connect <ip> <port>\t\t- Подключение к FTP-серверу")
	fmt.Println("disconnect\t\t\t- Отключение")
	fmt.Println("noop\t\t\t\t- проверить подключение")
	fmt.Println("login <username> <password>\t- Логин")
	fmt.Println("pwd\t\t\t\t- Вывести текущий каталог")
	fmt.Println("ls\t\t\t\t- Вывести содержимое текущего каталога")
	fmt.Println("cd <dir_name>\t\t\t- Переход в другой каталог")
	fmt.Println("get <file_name>\t\t\t- Скачать файл")
	fmt.Println("put <file_name>\t\t\t- Загрузить файл")

	return SUCCESS, nil
}
