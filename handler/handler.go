package handler

import (
	"fmt"
	"github.com/utilitywarehouse/dan-dawson-onboarding/metrics"
	"log"
	"net/http"
	"time"
)

func serveCurrentTime(w http.ResponseWriter, r *http.Request) {
	metrics.RM.Inc()
	fmt.Println("got here")
	if r.Method != http.MethodGet {
		invalidMethodMsg := fmt.Sprintf("invalid request method used: %v\n", r.Method)
		log.Print(invalidMethodMsg)
		metrics.FailedReqs.Inc()
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(invalidMethodMsg))

		if err != nil {
			log.Fatalf("error writing value to client: %v\n", err)
		}
		return
	}
	_, err := w.Write([]byte(time.Now().String()))

	if err != nil {
		metrics.FailedReqs.Inc()
		log.Fatalf("error writing value to client: %v\n", err)
	}
}

func InitHandler() {
	http.HandleFunc("/time", serveCurrentTime)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("failed to start server: %v\n", err.Error())
	}
}
