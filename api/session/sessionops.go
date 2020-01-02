package session

import (
	"Video-Server/api/dbops"
	"Video-Server/api/defs"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func nowInMillSecond() int64{
	return time.Now().UnixNano()/1000000
}

func deleteExpireSession(sid string){
	sessionMap.Delete(sid)
	err := dbops.DeleteSessionById(sid)
	if err !=nil{
		log.Printf("%s\n", err)
	}
}

func LoadSessionFromDB(){
	r,err := dbops.RetrieveAllSession()
	if err != nil{
		log.Printf("%s\n", err)
		return
	}
	r.Range(func(key, value interface{}) bool {
		simpleSession := value.(*defs.SimpleSession)
		sessionMap.Store(key,simpleSession)
		return true
	})
}

func GenerateNewSessionId(username string)string  {

}