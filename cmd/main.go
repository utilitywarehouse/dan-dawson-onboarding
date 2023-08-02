package main

import (
	"github.com/utilitywarehouse/dan-dawson-onboarding/handler"
	"log"
)

func main() {
	err := handler.InitHandler()

	if err != nil {
		log.Printf("server failed to start: %s", err.Error())
	}
}
