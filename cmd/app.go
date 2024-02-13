package main

import (
	"fmt"
	"ftp-client/internal/command"
	"ftp-client/internal/terminal"
	"os"
	"os/signal"
	"syscall"

	"github.com/jlaffaye/ftp"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	statusChan := make(chan command.STATUS)

	cmdChan := make(chan *command.CommandArgs)
	waitChan := make(chan struct{})

	termux := terminal.NewTerminal(cmdChan, waitChan)
	go termux.Run()

	var client *ftp.ServerConn
	invoker := command.NewInvoker()

	go func(client *ftp.ServerConn) {
		_ = <-sigChan
		if client != nil {
			client.Quit()
		}
		os.Exit(-1)
	}(client)

	go func() {
		for {
			select {
			case status := <-statusChan:
				if status == command.EXIT {
					sigChan <- syscall.SIGINT
				}
			}
		}
	}()

	for {
		select {
		case cmd := <-cmdChan:
			status, err := invoker.Execute(cmd, client)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			}

			fmt.Printf("Status recieved %d\n", status)
			statusChan <- status

			waitChan <- struct{}{}
		}
	}
}
