package main

import (
	"github.com/utilitywarehouse/dan-dawson-onboarding/handler"
	"github.com/utilitywarehouse/dan-dawson-onboarding/metrics"
)

//
//var RM = prometheus.NewCounter(prometheus.CounterOpts{
//	Name: "requests_metric",
//	Help: "Total number of requests",
//})

func main() {
	//h := op.NewHandler(
	//	op.NewStatus("dan-dawson-onboarding", "application that does stuff").
	//		AddOwner("team x", "#team-x").
	//		SetRevision("7470d3dc24ce7876a9fc53ca7934401273a4017a").
	//		AddChecker("db check", func(cr *op.CheckResponse) { cr.Healthy("dummy db connection check succeeded") }).
	//		AddMetrics(RM).
	//		ReadyUseHealthCheck(),
	//)
	//
	//go http.ListenAndServe(":8081", h)

	go metrics.InitMetrics()
	handler.InitHandler()
}
