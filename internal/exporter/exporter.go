package exporter

import (
	"net/http"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	crunch "github.com/verizonconnect/42crunch-client-go"
)

const (
	// Namespace is the metrics namespace of the exporter
	Namespace string = "fortytwo_crunch"
)

type Exporter struct {
	Client *crunch.Client
	Logger log.Logger
}

func (e *Exporter) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		registry := prometheus.NewRegistry()

		// Serve
		h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	}
}
