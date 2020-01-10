package defs

//request
type UserCredential struct {
	Username string
	Pwd      string
}
type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

type NewVideo struct {
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
}

//response
type SignUp struct {
	Success bool
	SessionId string
}

type SignIn struct {
	Success bool
	SessionId string
}
type UserInfo struct {
	Id int
}
type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

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
