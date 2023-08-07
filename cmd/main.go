package main

import (
	"github.com/utilitywarehouse/dan-dawson-onboarding/handler"
	"github.com/utilitywarehouse/dan-dawson-onboarding/metrics"
)

func main() {
	go metrics.InitMetrics()
	handler.InitHandler()
}
