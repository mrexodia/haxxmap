package proxy

import (
	"errors"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/client"
)

type user struct {
	be *Backend
	c *client.Client
	username string
}

func (u *user) Username() string {
	return u.username
}

func (u *user) listMailboxes(subscribed bool, name string) ([]backend.Mailbox, error) {
	mailboxes := make(chan *imap.MailboxInfo)
	done := make(chan error, 1)
	go func () {
		if subscribed {
			done <- u.c.Lsub("", name, mailboxes)
		} else {
			done <- u.c.List("", name, mailboxes)
		}
	}()

	var list []backend.Mailbox
	for m := range mailboxes {
		list = append(list, &mailbox{u: u, name: m.Name, info: m})
	}

	return list, <-done
}

func (u *user) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	return u.listMailboxes(subscribed, "*")
}

func (u *user) GetMailbox(name string) (backend.Mailbox, error) {
	mailboxes, err := u.listMailboxes(false, name)
	if err != nil {
		return nil, err
	}
	if len(mailboxes) == 0 {
		return nil, errors.New("No such mailbox")
	}

	m := mailboxes[0]
	if err := m.(*mailbox).ensureSelected(); err != nil {
		return nil, err
	}

	return m, err
}

func (u *user) CreateMailbox(name string) error {
	return u.c.Create(name)
}

func (u *user) DeleteMailbox(name string) error {
	return u.c.Delete(name)
}

func (u *user) RenameMailbox(existingName, newName string) error {
		return u.c.Rename(existingName, newName)
}

func (u *user) Logout() error {
	return u.c.Logout()
}
