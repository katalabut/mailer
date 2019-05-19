package mail

import (
	"fmt"
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

var sm *client

type client struct {
	addr string

	auth smtp.Auth
}

//NewClient created auth and check smtp dial
func NewClient(host string, port int, user string, pass string) {
	sm = &client{}
	sm.addr = fmt.Sprintf("%s:%d", host, port)
	sm.auth = smtp.PlainAuth("", user, pass, host)

	c, err := smtp.Dial(sm.addr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// if ok, _ := c.Extension("AUTH"); ok {
	// 	if err := c.Auth(sm.auth); err != nil {
	// 		log.Fatal("Auth: ", err)
	// 	}
	// }
}

//NewMessage - created new messagae for smtp clint
func NewMessage() *Message {
	return &Message{}
}

//Send - sended message
func Send(m *Message) error {
	msg := m.getMsg()
	return smtp.SendMail(sm.addr, sm.auth, m.from, m.to, msg)
}
