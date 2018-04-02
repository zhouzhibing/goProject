package http

import (
	"net/http"
	"fmt"
)

type HttpHandle struct {

}

func (this * HttpHandle) Handle(resp http.ResponseWriter, req *http.Request)  {
	fmt.Println("hello")
	resp.Write([]byte("hello"))
}