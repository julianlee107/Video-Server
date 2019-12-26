package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Handler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/", CreateUser)
	return router

}

func main() {
	//http.HandleFunc("/",RegisterHandler)
	r := Handler()
	e := http.ListenAndServe(":8000", r)
	if e != nil {
		log.Fatal(e)
	}
}
