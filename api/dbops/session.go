package dbops

import (
	"Video-Server/api/defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, userName string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("insert into sessions (session_id,TTL,login_name) values (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmtIns.Exec(sid, ttlstr, userName)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSessionById(sid string) (*defs.SimpleSession, error) {
	simpleSession := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL,login_name from sessions where session_id=?")
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	var ttl string
	var username string
	err = stmtOut.QueryRow(sid).Scan(&ttl, username)
	if err != nil && err != sql.ErrNoRows {
		panic(err.Error())
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		simpleSession.TTL = res
		simpleSession.Username = username
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return simpleSession, nil
}

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}
	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("%v\n", err)
			return nil, err
		}
		if ttl, err := strconv.ParseInt(ttlstr, 10, 64); err == nil {
			simpleSession := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, simpleSession)
			log.Printf("session id:%s,ttl:%d\n", id, ttl)
		}
	}
	return m, nil
}

func DeleteSessionById(sid string) error  {
	stmtOut, err:= dbConn.Prepare("delete from sessions where id=?")
	if err !=nil{
		log.Printf("%s\n", err)
		return err
	}
	if _,err := stmtOut.Query(sid);err!=nil{
		log.Printf("%s\n", err)
		return err
	}
	return nil
}