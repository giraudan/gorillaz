package gorillaz

func Init() {
	parseConfiguration()
	InitLogs()

	if tracing {
		InitTracing()
	}
}
