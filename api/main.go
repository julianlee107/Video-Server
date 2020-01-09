package main

import (
	"Video-Server/api/dbops"
)

//func Handler() *httprouter.Router {
//	router := httprouter.New()
//	router.POST("/", CreateUser)
//	return router
//
//}

func main() {
	//http.HandleFunc("/",RegisterHandler)
	_, err := dbops.AddNewVideo(1, "第一个视频")
	if err != nil {
		panic(err)
	}
}
