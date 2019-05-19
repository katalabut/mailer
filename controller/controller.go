package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// InitAndServ server controller
func InitAndServ() {
	mw := chainMiddleware(withTracing)
	http.Handle("/send", mw(sendHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
