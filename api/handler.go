package main

import (
	"Video-Server/api/dbops"
	"Video-Server/api/defs"
	"Video-Server/api/session"
	"Video-Server/api/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

//新增用户
func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	res, _ := ioutil.ReadAll(r.Body)
	userBody := &defs.User{}
	if err:=json.Unmarshal(res, userBody);err!=nil{
		log.Println(err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err:=dbops.AddUserCredential(userBody.LoginName,userBody.Pwd);err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	//生成session id
	id := session.GenerateNewSessionId(userBody.LoginName)
	signUp := &defs.SignUp{Success:true,SessionId:id}
	if resp,err:= json.Marshal(signUp);err!=nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else {
		sendNormalResponse(w,string(resp),201)
	}
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	res, _ := ioutil.ReadAll(r.Body)
	userBody := &defs.UserCredential{}
	if err:=json.Unmarshal(res, userBody);err!=nil{
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	userName := p.ByName("username")
	if userName!=userBody.Username{
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}
	pwd,err :=dbops.GetUserCredential(userName)
	if err != nil||len(pwd)==0||pwd !=userBody.Pwd {
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}
	id := session.GenerateNewSessionId(userName)
	signIn := &defs.SignIn{Success:true,SessionId:id}
	if resp,err:=json.Marshal(signIn);err!=nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
	}else {
		sendNormalResponse(w,string(resp),200)
	}
}

func GetUserInfo(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	if !validateUser(w,r){
		log.Println("Unauthorized user")
		return
	}
	userName := p.ByName("username")
	user,err := dbops.GetUser(userName)
	if err != nil {
		log.Println("error in finding user")
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	//指针传递和值传递皆可。
	userInfo := &defs.UserInfo{Id:user.Id}
	//userInfo := defs.UserInfo{Id:user.Id}
	if resp,err:= json.Marshal(userInfo);err!=nil{
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}



func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !validateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &defs.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	vi, err := dbops.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	log.Printf("Author id : %d, name: %s \n", nvbody.AuthorId, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}

func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !validateUser(w, r) {
		return
	}

	uname := p.ByName("username")
	vs, err := dbops.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ListAllvideos: %s", err)
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	vsi := &defs.VideosInfo{Videos: vs}
	if resp, err := json.Marshal(vsi); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}
