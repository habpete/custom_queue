package metrics

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	errorsHistogram = prometheus.NewCounterVec(prometheus.CounterOpts{}, []string{"handler"})
)

prometheus.MustRegister(errorsHistogram)

func ErrorsInc(src string) {
	errorsHistogram.WithLabelValues(src).Inc()
}
