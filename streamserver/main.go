package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	//router.GET("videos/:vid_id")
	return router
}
func main() {
	r := RegisterHandler()
	_ = http.ListenAndServe(":9000", r)
}
