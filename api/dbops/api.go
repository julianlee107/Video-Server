package dbops

import "log"

func AddUserCredential(loginName, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT  into user (login_name,pWd) VALUES (?,?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from user where login_name=?")
	if err != nil {
		log.Print("%s\n", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}
