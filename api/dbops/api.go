package dbops

import (
	"Video-Server/api/defs"
	"Video-Server/api/utils"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func AddUserCredential(loginName, pwd string) error {
	stmtOut, _ := dbConn.Prepare("select user.id from user where login_name=?")
	rows := stmtOut.QueryRow(loginName).Scan()
	if rows != sql.ErrNoRows {
		log.Println("用户名重复")
		return errors.New("用户名重复")
	}
	defer stmtOut.Close()
	stmtIns, err := dbConn.Prepare("INSERT into user (login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	crypt := sha256.New()
	crypt.Write([]byte(pwd))
	pwd = fmt.Sprintf("%x", crypt.Sum(nil))
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		panic(err)
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from user where login_name=?")
	if err != nil {
		log.Printf("%v\n", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM user WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("Delete User Error Found: %v", err)
		return err
	}
	crypt := sha256.New()
	crypt.Write([]byte(pwd))
	pwd = fmt.Sprintf("%x", crypt.Sum(nil))
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	defer stmtDel.Close()
	return nil
}

func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("select id,pwd,login_name from user where login_name=?")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	var id int
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd,&loginName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}
	return res, nil
}

func AddNewVideo(authorId int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	rightNow := time.Now()
	ctime := rightNow.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("insert into video_info (id, author_id, name, display_time) values (?,?,?,?)")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	_, err = stmtIns.Exec(vid, authorId, name, ctime)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: authorId, Name: name, DisplayTime: ctime}
	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_time FROM video_info WHERE id=?")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	var authorId int
	var ctime string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&authorId, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: authorId, Name: name, DisplayTime: ctime}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	defer stmtDel.Close()
	return nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select video_info.id,video_info.author_id,video_info.name,video_info.display_time from video_info inner join user on video_info.author_id=user.id where user.login_name=? and video_info.created_time between from_unixtime(?) and from_unixtime(?) order by video_info.created_time DESC ")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	var res []*defs.VideoInfo
	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	for rows.Next() {
		var id, name, ctime string
		var authorId int
		if err := rows.Scan(&id, &authorId, &name, &ctime); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		res = append(res, &defs.VideoInfo{Id: id, AuthorId: authorId, Name: name, DisplayTime: ctime})
	}
	defer stmtOut.Close()
	return res, nil
}

func AddNewComments(vid string, authorId int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	stmtIns, err := dbConn.Prepare("insert into comments (Id, Video_id,author_id,contents ,display_time) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	rightNow := time.Now()
	ctime := rightNow.Format("Jan 02 2006, 15:04:05")
	_, err = stmtIns.Exec(id, vid, authorId, content, ctime)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	var res []*defs.Comment
	stmtOut, err := dbConn.Prepare("select comments.id,user.login_name,comments.contents,comments.display_time from comments inner join user on comments.author_id=user.id where comments.Video_id=? and comments.created_time between from_unixtime(?) and from_unixtime(?)")
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	if rows.Next() {
		var id, author, content, ctime string
		err = rows.Scan(&id, &author, &content, &ctime)
		if err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		res = append(res, &defs.Comment{Id: id, VideoId: vid, Author: author, Content: content, DisplayTime: ctime})

	}
	defer stmtOut.Close()

	return res, nil
}
