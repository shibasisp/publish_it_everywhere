package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

var totalRequests int64

// LogHandler serves handlerfunc
func LogHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		totalRequests = totalRequests + 1
		log.Infoln("Total request received: ", totalRequests)
		log.Infoln(r.Method, ": ", r.RequestURI)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
