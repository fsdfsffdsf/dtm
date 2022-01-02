package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// https://yunlzheng.gitbook.io/prometheus-book/parti-prometheus-ji-chu/promql/what-is-prometheus-metrics-and-labels
// Metric: <metric name>{<label name>=<label value>, ...}
var (
	globalTransactionTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "tm_global_transaction_total",
		Help: "All transactions processed by tm",
	},
		[]string{"role", "model", "status"})

	globalTransactionCost = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "tm_global_transaction_cost",
		Help: "All transactions cost time by tm",
	},
		[]string{"role", "model", "status"})
)