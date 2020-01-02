package dbops

import "log"

func AddUserCredential(loginName, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT  into user (login_name,pwd) VALUES (?,?)")
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
		log.Printf("%v\n", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM user WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("Delete User Error Found: %v", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}