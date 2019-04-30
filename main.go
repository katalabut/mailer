package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type SmtpConf struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	Login    string `required:"true"`
	Password string `required:"true"`
}

func main() {

	os.Setenv("SMTP_HOST", "smtp.yandex.ru")
	os.Setenv("SMTP_PORT", "465")

	var sc SmtpConf
	err := envconfig.Process("smtp", &sc)
	if err != nil {
		log.Fatal(err.Error())
	}

	sendHandler := http.HandlerFunc(send)
	http.Handle("/send", sendHandler)
	http.ListenAndServe(":3000", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
