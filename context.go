package sdk

type PerformanceMetrics struct {
	Storage Metric
	Network Metric
}

type PerformCtx struct {
	Metrics PerformanceMetrics
}
