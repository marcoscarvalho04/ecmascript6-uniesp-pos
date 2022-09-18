package request

import "net/http"

const (
	NOT_FOUND             = 404
	INTERNAL_ERROR_SERVER = 500
	BAD_REQUEST           = 400
	CREATED               = 201
	OK                    = 200
	DELETED               = 202
)

func WriteStatusCodeAndMessage(w http.ResponseWriter, statusCode int, msg string) {
	w.Write([]byte(msg))
	w.WriteHeader(statusCode)
}
