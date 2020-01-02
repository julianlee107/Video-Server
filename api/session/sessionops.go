package session

import (
	"Video-Server/api/dbops"
	"Video-Server/api/defs"
	"Video-Server/api/utils"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMillSecond() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpireSession(sid string) {
	sessionMap.Delete(sid)
	err := dbops.DeleteSessionById(sid)
	if err != nil {
		log.Printf("%s\n", err)
	}
}

func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSession()
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	r.Range(func(key, value interface{}) bool {
		simpleSession := value.(*defs.SimpleSession)
		sessionMap.Store(key, simpleSession)
		return true
	})
}

func GenerateNewSessionId(username string) string {
	id, err := utils.NewUUID()
	if err != nil {
		log.Printf("%v\n", err)
		return ""
	}
	currentTime := nowInMillSecond()
	ttl := currentTime + 30*60*1000
	simpleSession := &defs.SimpleSession{Username: username, TTL: ttl}
	sessionMap.Store(id, simpleSession)
	err = dbops.InsertSession(id, ttl, username)
	if err != nil {
		log.Printf("%v\n", err)
		return ""
	}
	return id
}

func IsSessionIdExpired(sid string) (string, bool) {
	simpleSession, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMillSecond()
		if simpleSession.(*defs.SimpleSession).TTL < ct {
			deleteExpireSession(sid)
			return "", true
		}
		return simpleSession.(*defs.SimpleSession).Username, false
	}
	return "", true
}
