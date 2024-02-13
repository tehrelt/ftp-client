package session

import (
	"errors"
	"log"

	"github.com/jlaffaye/ftp"
)

var (
	ErrConnectionClosed        = errors.New("Connection closed")
	ErrConnectionAlreadyOpened = errors.New("Connection already opened")
)

type Session struct {
	client *ftp.ServerConn

	IsOpen bool
}

func NewSession() *Session {
	return &Session{}
}

func (session *Session) Connect(addr string) error {
	if session.IsOpen {
		return ErrConnectionAlreadyOpened
	}

	var err error
	if session.client, err = ftp.Dial(addr); err != nil {
		return err
	}

	session.IsOpen = true

	return nil
}

func (session *Session) Disconnect() error {
	if !session.IsOpen {
		return ErrConnectionClosed
	}

	if err := session.client.Quit(); err != nil {
		return err
	}

	session.IsOpen = false

	log.Println("Disconnected")

	return nil
}

func (session *Session) Login(username, password string) error {

	if !session.IsOpen {
		return ErrConnectionClosed
	}

	if err := session.client.Login(username, password); err != nil {
		return err
	}

	return nil
}

func (session *Session) Noop() error {

	if !session.IsOpen {
		return ErrConnectionClosed
	}

	return session.client.NoOp()
}

func (session *Session) Pwd() (string, error) {
	if !session.IsOpen {
		return "", ErrConnectionClosed
	}

	return session.client.CurrentDir()
}

func (session *Session) Ls() ([]*ftp.Entry, error) {
	if !session.IsOpen {
		return nil, ErrConnectionClosed
	}

	pwd, err := session.Pwd()
	if err != nil {
		return nil, err
	}

	return session.client.List(pwd)
}
