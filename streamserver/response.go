package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	_, err := io.WriteString(w, errMsg)
	if err != nil {
		panic(err)
	}
}
