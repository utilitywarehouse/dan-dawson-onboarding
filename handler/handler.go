package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func serveCurrentTime(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		invalidMethodMsg := fmt.Sprintf("invalid request method used: %v\n", r.Method)
		log.Printf(invalidMethodMsg)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(invalidMethodMsg))
		return
	}
	w.Write([]byte(time.Now().String()))
}

func InitHandler() {
	http.HandleFunc("/time", serveCurrentTime)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("failed to start server: %v\n", err.Error())
	}
}
