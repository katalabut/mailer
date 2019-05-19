package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spotbut/mailer/controller"
	"github.com/spotbut/mailer/mail"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	mail.NewClient("smtp.yandex.ru", 587, "testpromsp@yandex.ru", "evdrimazriaedhqg")
	controller.InitAndServ()
}
