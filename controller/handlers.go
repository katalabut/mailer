package controller

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type msg struct {
	To      string
	Subject string
	Body    string
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var ms msg
	err := decoder.Decode(&ms)
	if err != nil {
		log.Info(err)
		return
	}

	// m := s.smtp.NewMessage()
	// m.From("testpromsp@yandex.ru")
	// m.Subject(ms.Subject)
	// m.To(ms.To)

	// err = s.smtp.Send(m)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	w.Write([]byte("Message is sended"))
}
