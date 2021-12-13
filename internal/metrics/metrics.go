package metrics

import (
	"time"

	prometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	errorsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{}, []string{"handler"})

	timerHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{}, []string{"handler"})
)

func init() {
	prometheus.MustRegister(errorsCounter)
	prometheus.MustRegister(timerHistogram)
}

func ErrorsInc(src string) {
	errorsCounter.WithLabelValues(src).Inc()
}

func TimerHistogram(src string) func() {
	startTime := time.Now()

	return func() {
		timerHistogram.WithLabelValues(src).Observe(float64(time.Since(startTime)))
	}
}
