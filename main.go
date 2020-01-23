package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/katalabut/mailer/controller"
	"github.com/katalabut/mailer/mail"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)
}

func main() {
	mail.NewClient("host", 587, "mail@local", "pass")
	controller.InitAndServ()
}
