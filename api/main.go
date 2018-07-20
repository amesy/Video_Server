package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New() 
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8800", r)
}

//handler -> validation(校验){1.request, 2.user} -> business logic(逻辑处理) -> response
//1. data model
//2. error handling.
//
//session
