package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/utilitywarehouse/go-operational/op"
)

var RM = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "total_time_requests",
	Help: "Total number of time requests",
})

var FailedReqs = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "total_failed_requests",
	Help: "Total number of failed requests",
})

var metricsHandler = op.NewHandler(
	op.NewStatus("dan-dawson-onboarding", "application that does stuff").
		AddOwner("team x", "#team-x").
		SetRevision("7470d3dc24ce7876a9fc53ca7934401273a4017a").
		AddChecker("db check", func(cr *op.CheckResponse) { cr.Healthy("dummy db connection check succeeded") }).
		AddMetrics(RM, FailedReqs).
		ReadyUseHealthCheck(),
)

func InitMetrics() {
	http.ListenAndServe(":8081", metricsHandler)
}
