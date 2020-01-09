package defs

//数据模板
type User struct {
	Id        int
	LoginName string
	Pwd       string
}
type SimpleSession struct {
	Username string
	TTL      int64
}
type UserCredential struct {
	Username string
	Pwd      string
}

type VideoInfo struct {
	Id          string
	AuthorId    int
	Name        string
	DisplayTime string
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
	DisplayTime string

}
