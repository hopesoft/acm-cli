package controller

import (
	"acm-cli/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Controller struct {
	Writer 			http.ResponseWriter
	Request 		*http.Request
}

func (c *Controller) Handle(f func()) func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		c.Writer = writer
		c.Request = request
		f()
	}
}
func (c *Controller) Failed(err error, code int) {
	http.Error(c.Writer, err.Error(), code)
	log.Printf("请求错误: %#v", err)
}

func (c *Controller) Response(str string) {
	_, _ = c.Writer.Write([]byte(str))

}

func (c *Controller)Version()  {
	c.Response(handler.AcmVersion)
}