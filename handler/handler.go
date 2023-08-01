package handler

import (
	"net/http"
	"time"
)

func serveCurrentTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().String()

	w.Write([]byte(currentTime))
}

func InitHandler() {
	http.HandleFunc("/", serveCurrentTime)
	http.ListenAndServe(":8080", nil)
}
