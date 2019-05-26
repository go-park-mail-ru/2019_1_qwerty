package handlers

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var FooCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "foo_total",
	Help: "Number of foo successfully processed.",
})

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

// ErrorMux - Мультиплексор ошибок
func ErrorMux(w *http.ResponseWriter, r *http.Request, statusCode int) {
	Hits.WithLabelValues(strconv.Itoa(statusCode), r.URL.String()).Inc()
	FooCount.Add(1)
	(*w).WriteHeader(statusCode)
}
