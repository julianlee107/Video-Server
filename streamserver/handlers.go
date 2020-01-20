package main

import (
	"Video-Server/streamserver/defs"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := defs.VIDEO_DIR + vid
	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "error of open video: "+err.Error())
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, defs.MAX_UPLOAD_SIZE)
	if err:=r.ParseMultipartForm(defs.MAX_UPLOAD_SIZE); err!=nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too large")
		return
	}

	//此处一定要这么写，否则读不出文件
	file, _, err := r.FormFile("file")
	if err!=nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	data, err:=ioutil.ReadAll(file)
	if err!=nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	fn:=p.ByName("vid-id")
	err = ioutil.WriteFile(defs.VIDEO_DIR + fn, data, 06666)
	if err!=nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully.")
}