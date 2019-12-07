package controller

import (
	"acm-cli/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)
type Route struct {
	Path        string
	Method      string
	ActionFunc func()
}

func routers(ctrl *Controller) []Route {
	return []Route{
		{"/config", "GET", ctrl.Get},
		{"/config", "POST", ctrl.Set},
		{"/config", "Del", ctrl.Del},
		{"/version", "GET", ctrl.Version},
	}
}

func ListenServer() {
	router := httprouter.New()
	ctrl := &Controller{}
	for _, route := range routers(ctrl) {
		router.Handle(route.Method, route.Path, ctrl.Handle(route.ActionFunc))
	}
	err := http.ListenAndServe(":" + handler.AcmEnv["port"], router)
	if err != nil {
		log.Fatalln(err)
	}
}