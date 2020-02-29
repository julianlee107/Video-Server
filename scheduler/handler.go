package scheduler

import (
	"Video-Server/scheduler/dbops"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func vidSetDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video id should not be empty")
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal Server error")
		return
	}
	sendResponse(w, 200, "delete video successfully")
}

func vidGetDelRecHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	count, err := strconv.Atoi(p.ByName("count"))
	var record []string
	if err != nil {
		sendResponse(w,500,"the count doesn't look right")
		return
	}
	if count == 0 {
		sendResponse(w, 400, "video id should not be empty")
		return
	}
	record,err = dbops.GetVideoDeletionRecord(count)
	if err != nil {
		sendResponse(w,500,"bad record")
		return
	}
	if resp,err := json.Marshal(record);err!=nil{
		sendResponse(w,200,string(resp))
	}else {
		sendResponse(w,500,"bad record")
	}

}
