package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	dbConn *sql.DB
	err error
)

func init()  {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	dbConn,err = sql.Open("mysql", path)
	if err !=nil{
		panic(err.Error())
	}
	fmt.Printf("dbConn: %v\n",dbConn)
}
