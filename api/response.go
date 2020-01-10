package main

import (
	"Video-Server/api/defs"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, err := json.Marshal(&errResp.Error)
	if err != nil {
		log.Printf("%v", err)
	}
	_, err = io.WriteString(w, string(resStr))
	if err != nil {
		log.Printf("%v", err)
	}
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	_, err := io.WriteString(w, resp)
	if err != nil {
		log.Printf("%v", err)
	}
}
