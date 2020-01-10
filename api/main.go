package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.GET("/user/:username",GetUserInfo)
	return router

}

func main() {
	//http.HandleFunc("/",RegisterHandler)
	//_, err := dbops.AddNewVideo(1, "第一个视频")
	//if err != nil {
	//	panic(err)
	//}
	r:=RegisterHandler()
	http.ListenAndServe(":8000",r)
}
