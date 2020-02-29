package main

import (
	"Video-Server/streamserver/limiter"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middlewareHandler struct {
	r *httprouter.Router
	l *limiter.ConnLimiter
}

func NewMiddleWareHanlder(r *httprouter.Router, cc int) http.Handler {
	m := middlewareHandler{
		r: nil,
		l: nil,
	}
	m.r = r
	m.l = limiter.NewConnLimiter(cc)
	return m
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid_id", streamHandler)
	router.POST("/upload/:vid_id", uploadHandler)
	return router
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandler()
	mh := NewMiddleWareHanlder(r, 10)
	http.ListenAndServe(":9000",mh)
}
