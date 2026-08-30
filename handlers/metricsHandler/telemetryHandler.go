package metricsHandler

import (
	"github.com/redis/go-redis/v9"
	"net/http"
)

type TelemetryHandler struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *TelemetryHandler {
	return &TelemetryHandler{
		rdb: rdb,
	}
}

func (h *TelemetryHandler) Metrics(w http.ResponseWriter, r *http.Request) {
	httpTotalRequest := ""
	totalErrors := ""
	averageLatency := ""
	response := "http_total_requests " + httpTotalRequest + "\n" +
		"total_errors " + totalErrors + "\n" +
		"total_request_latency " + averageLatency
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	w.Write([]byte(response))
}
