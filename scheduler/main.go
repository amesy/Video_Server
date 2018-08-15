package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"video_server/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router:=httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	go taskrunner.Start()
	r:=RegisterHandlers()
	http.ListenAndServe(":9001", r)
}