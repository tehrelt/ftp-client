package main

import (
	"fmt"
	"ftp-client/internal/command"
	sess "ftp-client/internal/session"
	"ftp-client/internal/terminal"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	session := sess.NewSession()
	invoker := command.NewInvoker(session)

	go func(client *sess.Session) {
		_ = <-sigChan
		if client.IsOpen {
			client.Disconnect()
		}
		os.Exit(-1)
	}(session)

	go func() {
		for {
			select {
			case status := <-statusChan:
				if status == command.EXIT {
					os.Exit(0)
				}
			}
		}
	}()

	for {
		select {
		case cmd := <-cmdChan:
			status, err := invoker.Execute(cmd)
			if err != nil {
				if err == command.ErrArgs {
					fmt.Printf("Enter 'help' for args-guide\n")
				} else {
					log.Println(err)
				}
			}

			statusChan <- status

			waitChan <- struct{}{}
		}
	}
}
