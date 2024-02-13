package session

import (
	"errors"
	"io"
	"log"
	"strings"

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

func (session *Session) List() ([]*ftp.Entry, error) {
	if !session.IsOpen {
		return nil, ErrConnectionClosed
	}

	pwd, err := session.Pwd()
	if err != nil {
		return nil, err
	}

	return session.client.List(pwd)
}

func (session *Session) Cwd(dir string) error {
	if !session.IsOpen {
		return ErrConnectionClosed
	}

	if strings.Compare(dir, "..") == 0 {
		return session.client.ChangeDirToParent()
	}

	return session.client.ChangeDir(dir)
}

func (session *Session) FileSize(file string) (int64, error) {
	return session.client.FileSize(file)
}

func (session *Session) Get(file string) (*ftp.Response, error) {
	if !session.IsOpen {
		return nil, ErrConnectionClosed
	}

	return session.client.Retr(file)
}

func (session *Session) Put(file string, bytes io.Reader) error {
	if !session.IsOpen {
		return ErrConnectionClosed
	}

	return session.client.Stor(file, bytes)
}
