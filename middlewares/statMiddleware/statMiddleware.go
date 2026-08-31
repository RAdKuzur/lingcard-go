package statMiddleware

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

type StatMiddleware struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *StatMiddleware {
	return &StatMiddleware{
		rdb: rdb,
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (m *StatMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()
		wrapped := &responseWriter{w, http.StatusOK}
		key := fmt.Sprintf("http_requests:%s:%s", r.Method, r.URL.Path)
		pipe := m.rdb.Pipeline()
		pipe.Incr(ctx, "http_total_requests")
		pipe.Incr(ctx, key)

		if _, err := pipe.Exec(ctx); err != nil {
			log.Printf("Failed to increment counters: %v", err)
		}
		next.ServeHTTP(wrapped, r)
		latencyMs := time.Since(start).Milliseconds()
		pipe2 := m.rdb.Pipeline()
		pipe2.IncrBy(ctx, "total_request_latency", latencyMs)
		if _, err := pipe2.Exec(ctx); err != nil {
			log.Printf("Failed to record latency: %v", err)
		}
	})
}
