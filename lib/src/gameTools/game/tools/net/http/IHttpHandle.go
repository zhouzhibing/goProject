package http

import "net/http"

type IHttpHandle interface {
	Handle(resp http.ResponseWriter, req *http.Request)
}